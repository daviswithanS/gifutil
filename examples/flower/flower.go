// flower.go creates a gif of a flower - that is, a shape formed by
// the polar equation r = k * cos(p * theta), where k is the radius of
// the final product and p is the number of petals. it uses the gg
// library for all of the drawing, found at github.com/fogleman/gg.
package main

import (
	"fmt"
	"github.com/daviswithanS/gifutil"
	"github.com/fogleman/gg"
	"image"
	"image/color/palette"
	"log"
	"math"
)

var width, height int = 800, 800
var steps int = 33

func getCoordinates(theta float64) (x, y float64) {
	r := 3 * math.Cos(64*theta)
	x = r * math.Cos(theta)
	y = r * math.Sin(theta)
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
	dc.Translate(4, 4)

	dc.SetHexColor("333")
	dc.Clear()

	theta := float64(0)
	dTheta := math.Pi / 128

	dc.MoveTo(getCoordinates(theta))

	getFrame = func(step int) *image.Image {
		if step == 0 {
			x, y := getCoordinates(theta)
			dc.MoveTo(x, y)
		} else {
			// update theta and extend the line 8 times per step in order to have
			// a higher resolution
			for i := 0; i < 8; i++ {
				theta += dTheta
				x, y := getCoordinates(theta)
				dc.LineTo(x, y)
			}
		}

		dc.SetLineWidth(2)
		dc.SetHexColor("FFF")
		dc.StrokePreserve()

		fmt.Printf("%v/%v\n", step, steps-1) // a "progress bar" of sorts

		if step == steps-1 {
			dc.SavePNG("flower.png")
		}

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

	out.Delay[len(out.Delay)-1] = 100

	// output gif to file
	writeErr := gifutil.WriteToFile(out, "flower.gif")
	if writeErr != nil {
		log.Fatal(writeErr)
	}
}
