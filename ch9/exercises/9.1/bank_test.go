// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank_test

import (
	"fmt"
	"testing"

	bank "gopl.io/ch9/exercises/9.1"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// Alice
	go func() {
		bank.Deposit(200)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		bank.Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := bank.Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}

func TestWithdraw(t *testing.T) {
	done := make(chan struct{})

	// Deposit
	go func() {
		bank.Deposit(500)
		fmt.Println("=", bank.Balance())
		done <- struct{}{}
	}()
	// wait for initial deposit
	<-done

	// Alice
	go func() {
		if got, want := bank.Withdraw(200), true; got != want {
			t.Errorf("Alice withdraw result = %t, want %t", got, want)
		}
		done <- struct{}{}
	}()

	// Bob
	go func() {
		if got, want := bank.Withdraw(200), true; got != want {
			t.Errorf("Bob withdraw result = %t, want %t", got, want)
		}
		done <- struct{}{}
	}()

	// Wait for both transaction
	<-done
	<-done

	// David
	go func() {
		if got, want := bank.Withdraw(200), false; got != want {
			t.Errorf("David withdraw result = %t, want %t", got, want)
		}
		done <- struct{}{}
	}()

	// Wait for last transaction
	<-done

	if got, want := bank.Balance(), 100; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
