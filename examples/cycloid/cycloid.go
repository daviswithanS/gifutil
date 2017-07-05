// cycloid.go creates a gif demonstrating a cycloid - the shape that
// is formed when you trace one point on a rolling circle. it uses the
// gg library for all of the drawing, found at github.com/fogleman/gg.
package main

import (
	"github.com/daviswithanS/gifutil"
	"github.com/fogleman/gg"
	"image"
	"image/color/palette"
	"log"
	"math"
)

var width, height int = 600, 400
var steps int = 32

// get the xy coordinates for a given angle according to the cycloid functions
func getCoordinates(theta float64) (x, y float64) {
	sin, cos := math.Sincos(theta)
	x = theta - sin
	y = 1 - cos
	return
}

// getFrameGenerator() creates the function that will be passed to
// gifutil.Populate() to create the frames of the gif. it is written
// this way so that, by closure, any variables created in this generator
// function will remain persistent through calls of getFrame()
func getFrameGenerator() (getFrame func(int) *image.Image) {
	dc := gg.NewContext(width, height)
	dc.InvertY()
	dc.Scale(100, 100)
	dc.Translate(3, 1)

	dTheta := math.Pi / 16

	getFrame = func(step int) *image.Image {
		theta := float64(step) * dTheta

		dc.SetHexColor("FCF")
		dc.Clear()

		dc.Translate(-dTheta, 0) // keep the circle in the center

		padding := 2 * math.Pi
		// the path to the current point and the circle are redrawn every time
		// so that they can update their position
		for t := -padding; t <= theta+padding; t += dTheta {
			x, y := getCoordinates(t)
			dc.LineTo(x, y)
		}
		dc.SetHexColor("000")
		dc.SetLineWidth(4)
		dc.Stroke()

		dc.DrawCircle(theta, 1, 1)
		dc.SetLineWidth(2)
		dc.Stroke()

		x, y := getCoordinates(theta)
		dc.DrawPoint(x, y, 6)
		dc.SetHexColor("66C")
		dc.Fill()

		frame := dc.Image()
		return &frame
	}
	return
}

func main() {
	// create a new gif
	out := gifutil.NewGIF(palette.WebSafe, width, height)

	// fill it with frames
	getFrame := getFrameGenerator()
	gifutil.Populate(out, steps, getFrame)

	// output gif to file
	writeErr := gifutil.WriteToFile(out, "cycloid.gif")
	if writeErr != nil {
		log.Fatal(writeErr)
	}
}
