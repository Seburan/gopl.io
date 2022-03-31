// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Exercise 2.2 is a general purpose unit-conversion program that reads numbers
// from its command-line argument or from the standard input if there are no
// arguement and converts each number into units like temperature un celisus
// and farenheit, lengths in feet and meters, weight in pounds and kilograms
package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"gopl.io/ch2/exercises/unitconv"
)

var from = flag.String("from", " ", "Conversion From Unit : Farenheit, Feet, Pound etc.")
var to = flag.String("to", "", "Conversion to Unit : Celsius, Meter, Kg")

func main() {

	flag.Parse()
	var fromUnit string = strings.ToLower(*from)
	var toUnit string = strings.ToLower(*to)

	fmt.Printf("Conversion from %s to %s\n", fromUnit, toUnit)

	for _, arg := range flag.Args() {
		f, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		switch fromUnit {
		case "kelvin":
			fromValue := unitconv.Kelvin(f)
			switch toUnit {
			case "celsius":
				fmt.Printf("%s = %s\n", fromValue, unitconv.KToC(fromValue))
			default:
				println("Sorry this conversion is not available yet")
			}
		case "farenheit":
			fromValue := unitconv.Fahrenheit(f)
			switch toUnit {
			case "celsius":
				fmt.Printf("%s = %s\n", fromValue, unitconv.FToC(fromValue))
			default:
				println("Sorry this conversion is not available yet")
			}
		case "feet":
			fromValue := unitconv.Feet(f)
			switch toUnit {
			case "meter":
				fmt.Printf("%s = %s\n", fromValue, unitconv.FToM(fromValue))
			default:
				println("Sorry this conversion is not available yet")
			}
		default:
			println("Sorry this conversion is not availableyet ")
		}

	}
}

//!-
