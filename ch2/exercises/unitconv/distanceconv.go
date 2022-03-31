// Package tempconv performs Feet and Meter conversions.
package unitconv

import "fmt"

// Exercise 2.2 :
type Feet float64
type Meter float64


func (f Feet) String() string    { return fmt.Sprintf("%g ft", f) }
func (m Meter) String() string { return fmt.Sprintf("%g m", m) }


func FToM(f Feet) Meter { return Meter(f/3.2808) }

//!-
