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
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

const (


)



func main() {

	// default values
	var width, height float64 = 600, 320;            // canvas size in pixels
	var cells float64        = 100                 // number of grid cells
	var xyrange float64      = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale := float64(width) / 2 / xyrange // pixels per x or y unit
	zscale := height * 0.4        // pixels per z unit
	angle := math.Pi / 6         // angle of x, y axes (=30°)


	// Exercise 3.4 : Write SVG data to client browser
	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {

			// read parameters from URL
			if err := r.ParseForm(); err != nil {
				log.Print(err)
			}

			var err error;
			if r.Form.Has("width") {
				width, err = strconv.ParseFloat(r.Form.Get("width"), 64);
				log.Print("width = ", width);
				if err != nil {
					log.Print(err)
				}
			}

			if r.Form.Has("height") {
				width, err = strconv.ParseFloat(r.Form.Get("height"), 64);
				log.Print("height = ", height);
				if err != nil {
					log.Print(err)
				}
			}

			w.Header().Set("Content-Type", "image/svg+xml");

			// update parameters
			xyscale = float64(width) / 2 / xyrange // pixels per x or y unit
			zscale = float64(height) * 0.4        // pixels per z unit
			angle = math.Pi / 6         // angle of x, y axes (=30°)

			writeSVG(w, width, height, cells, xyrange, xyscale, zscale, angle);
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe(":8000", nil))
		return
	}
	//!+main
	writeSVG(os.Stdout, width, height, cells, xyrange, xyscale, zscale, angle);

}

// writeSVG compute surfaces and write SVG to out
func writeSVG(out io.Writer, width float64, height float64, cells float64, xyrange float64, xyscale float64, zscale float64, angle float64) {

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < int(cells); i++ {
		for j := 0; j < int(cells); j++ {
			ax, ay, invalid := corner(i+1, j, width, height, cells, xyrange, xyscale, zscale, angle);
			if invalid != nil {
			 	err := fmt.Errorf("Skipping invalid polygon at at position i=%d,j=%d. Reason : %w\n",i,j, invalid);
				log.Println(err.Error())
				continue;
			}
			bx, by, invalid := corner(i, j, width, height, cells, xyrange, xyscale, zscale, angle);
			if invalid != nil {
			 	err := fmt.Errorf("Skipping invalid polygon at at position i=%d,j=%d. Reason : %w\n",i,j, invalid);
				log.Println(err.Error())
				continue;
			}
			cx, cy, invalid := corner(i, j+1, width, height, cells, xyrange, xyscale, zscale, angle);
			if invalid != nil {
			 	err := fmt.Errorf("Skipping invalid polygon at at position i=%d,j=%d. Reason : %w\n",i,j, invalid);
				log.Println(err.Error())
				continue;
			}
			dx, dy, invalid := corner(i+1, j+1, width, height, cells, xyrange, xyscale, zscale, angle);
			if invalid != nil {
			 	err := fmt.Errorf("Skipping invalid polygon at at position i=%d,j=%d. Reason : %w\n",i,j, invalid);
				log.Println(err.Error())
				continue;
			}
			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

// corner returns an error if the polygon corner cannot be computed or invalid
func corner(i, j int, width float64, height float64, cells float64, xyrange float64, xyscale float64, zscale float64, angle float64) (float64, float64, error) {

	// default angle
	var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/float64(cells) - 0.5)
	y := xyrange * (float64(j)/float64(cells) - 0.5)

	// Compute surface height z.
	z := f(x, y)

	// if infinite skip processing polygon and returns
	if math.IsInf(z, 0) {
		return 0,0, errors.New("f returns infinite value that cannot be used to create a polygon corner");
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(width)/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy, nil;
}


func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

//!-
