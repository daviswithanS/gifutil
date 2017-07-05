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

func getFrameGenerator() (getFrame func(int) *image.Image) {
	// getFrame() closes over the drawing context so that it remains
	// persistent through calls of getFrame()
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

		fmt.Printf("%v/%v\n", step, steps-1)

		if step == steps-1 {
			dc.SavePNG("flower.png")
		}

		frame := dc.Image()
		return &frame
	}
	return
}

func main() {
	out := gifutil.NewGIF(palette.WebSafe, width, height)

	getFrame := getFrameGenerator()
	gifutil.Populate(out, steps, getFrame)

	out.Delay[len(out.Delay)-1] = 100

	writeErr := gifutil.WriteToFile(out, "flower.gif")
	if writeErr != nil {
		log.Fatal(writeErr)
	}
}
