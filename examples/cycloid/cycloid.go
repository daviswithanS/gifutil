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

var width, height = 600, 400
var steps = 32

// get the xy coordinates for a given angle according to the cycloid functions
func getCoordinates(theta float64) (x, y float64) {
	sin, cos := math.Sincos(theta)
	x = theta - sin
	y = 1 - cos
	return
}

type frameGenerator struct {
	dc *gg.Context
}

func newFrameGenerator() (fg *frameGenerator) {
	fg = &frameGenerator{dc: gg.NewContext(width, height)}

	fg.dc.InvertY()
	fg.dc.Scale(100, 100)
	fg.dc.Translate(3, 1)

	return fg
}

func (fg *frameGenerator) GetFrame(step int) *image.Image {
	const dTheta = math.Pi / 16
	theta := float64(step) * dTheta

	fg.dc.SetHexColor("FCF")
	fg.dc.Clear()

	fg.dc.Translate(-dTheta, 0) // keep the circle in the center

	padding := 2 * math.Pi
	// the path to the current point and the circle are redrawn every time
	// so that they can update their position
	for t := -padding; t <= theta+padding; t += dTheta {
		x, y := getCoordinates(t)
		fg.dc.LineTo(x, y)
	}
	fg.dc.SetHexColor("000")
	fg.dc.SetLineWidth(4)
	fg.dc.Stroke()

	fg.dc.DrawCircle(theta, 1, 1)
	fg.dc.SetLineWidth(2)
	fg.dc.Stroke()

	x, y := getCoordinates(theta)
	fg.dc.DrawPoint(x, y, 6)
	fg.dc.SetHexColor("66C")
	fg.dc.Fill()

	frame := fg.dc.Image()
	return &frame
}

func main() {
	// create a new gif
	out := gifutil.NewGIF(palette.WebSafe, width, height)

	// fill it with frames
	fg := newFrameGenerator()
	gifutil.Populate(out, steps, fg)

	// output gif to file
	writeErr := gifutil.WriteToFile(out, "cycloid.gif")
	if writeErr != nil {
		log.Fatal(writeErr)
	}
}
