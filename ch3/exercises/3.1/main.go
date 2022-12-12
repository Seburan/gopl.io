// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 58.
//!+

// Exercise 3.1 : Modify the program to skip invalid polygons (thos with
// non-finite float64 value )

// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"errors"
	"fmt"
	"log"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, invalid := corner(i+1, j);
			if invalid != nil {
			 	err := fmt.Errorf("Skipping invalid polygon at at position i=%d,j=%d. Reason : %w\n",i,j, invalid);
				log.Println(err.Error())
				continue;
			}
			bx, by, invalid := corner(i, j)
			if invalid != nil {
			 	err := fmt.Errorf("Skipping invalid polygon at at position i=%d,j=%d. Reason : %w\n",i,j, invalid);
				log.Println(err.Error())
				continue;
			}
			cx, cy, invalid := corner(i, j+1)
			if invalid != nil {
			 	err := fmt.Errorf("Skipping invalid polygon at at position i=%d,j=%d. Reason : %w\n",i,j, invalid);
				log.Println(err.Error())
				continue;
			}
			dx, dy, invalid := corner(i+1, j+1)
			if invalid != nil {
			 	err := fmt.Errorf("Skipping invalid polygon at at position i=%d,j=%d. Reason : %w\n",i,j, invalid);
				log.Println(err.Error())
				continue;
			}
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Println("</svg>")
}

// corner returns an error if the polygon corner cannot be computed or invalid
func corner(i, j int) (float64, float64, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// if infinite skip processing polygon and returns
	if math.IsInf(z, 0) {
		return 0,0, errors.New("f returns infinite value that cannot be used to create a polygon corner");
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, nil;
}


func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
