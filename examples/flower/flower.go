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

var width, height = 800, 800
var steps = 33

// convert the current theta to a set of x y coordinates
func getCoordinates(theta float64) (x, y float64) {
	r := 3 * math.Cos(64*theta)
	x = r * math.Cos(theta)
	y = r * math.Sin(theta)
	return
}

// frameGenerator implements GetFrame, creating the next iteration of the
// flower each step
type frameGenerator struct {
	dc *gg.Context
	theta float64
}

func newFrameGenerator() *frameGenerator {
	fg := &frameGenerator{dc: gg.NewContext(width, height)}
	fg.dc.InvertY()
	fg.dc.Scale(100, 100)
	fg.dc.Translate(4, 4)

	fg.dc.SetHexColor("333")
	fg.dc.Clear()

	theta := float64(0)

	fg.dc.MoveTo(getCoordinates(theta))
	return fg
}

func (fg *frameGenerator) GetFrame(step int) *image.Image {
	const dTheta = math.Pi / 128

	if step == 0 {
		x, y := getCoordinates(fg.theta)
		fg.dc.MoveTo(x, y)
	} else {
		// update theta and extend the line 8 times per step in order to have
		// a higher resolution
		for i := 0; i < 8; i++ {
			fg.theta += dTheta
			x, y := getCoordinates(fg.theta)
			fg.dc.LineTo(x, y)
		}
	}

	fg.dc.SetLineWidth(2)
	fg.dc.SetHexColor("FFF")
	fg.dc.StrokePreserve()

	fmt.Printf("%v/%v\n", step, steps-1) // a "progress bar" of sorts

	if step == steps-1 {
		fg.dc.SavePNG("flower.png")
	}

	frame := fg.dc.Image()
	return &frame
}

func main() {
	// create a new gif
	out := gifutil.NewGIF(palette.WebSafe, width, height)

	// fill it with frames
	fg := newFrameGenerator()
	gifutil.Populate(out, steps, fg)

	out.Delay[len(out.Delay)-1] = 100

	// output gif to file
	writeErr := gifutil.WriteToFile(out, "flower.gif")
	if writeErr != nil {
		log.Fatal(writeErr)
	}
}
