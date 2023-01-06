// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 261.
//!+

// Package bank provides a concurrency-safe bank with one account.

package bank

import "fmt"

var deposits = make(chan int) // send amount to deposit
var balances = make(chan int) // receive balance

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }

// Exercise 9.1 (P262)
// Add a function Withdraw(amount int) bool to the gopl.io/ch9/bank1 program.
// The result should indicate whether the transaction succeeded or failed due
// to insufficient funds. The message send to the monitor goroutine must contain
// both the amount to withdraw and a new channel over which the monitor
// goroutine can send the boolean result back to Withdraw.

// request contains both the amount to withdraw and a new channel over which the
// monitor goroutine can send the boolean result back
type request struct {
	amount   int
	response chan bool // the client wants to know if the result of the operation
}

var withdraws = make(chan request) // channel to send amount to withdraw

func Withdraw(amount int) bool {
	var request = request{
		amount:   amount,
		response: make(chan bool), // receive the result of the withdraw operation
	}
	withdraws <- request // send to amount to withdraw to the teller go routine

	return <-request.response // return the result of the withdraw operation
}

func teller() {
	var balance int // balance is confined to teller goroutine
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case request := <-withdraws:
			if balance >= request.amount {
				balance -= request.amount
				request.response <- true
			} else { // insufficent balance
				fmt.Printf("Insufficient balance (%d) to withdraw (%d)", balance, request.amount)
				request.response <- false
			}

		case balances <- balance:
		}
	}
}

func init() {
	go teller() // start the monitor goroutine
}

//!-
