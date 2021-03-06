// Package gifutil contains utility functions for working with standard library "image/gif" GIF structs.
package gifutil

import (
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"os"
)

// FrameGetter is an interface used by the Populate() function. All that is
// required to implement it is a function that, when given a frame number,
// returns the corresponding image associated with that frame.
type FrameGetter interface {
	GetFrame(int) *image.Image
}

// NewGIF creates a new, empty GIF struct that loops indefinitely.
func NewGIF(pal color.Palette, width, height int) *gif.GIF {
	return NewFiniteGIF(pal, width, height, 0)
}

// NewFiniteGIF creates a new, empty GIF struct that loops a fixed number of times.
func NewFiniteGIF(pal color.Palette, width, height, loopCount int) *gif.GIF {
	config := image.Config{ColorModel: pal, Width: width, Height: height}
	return &gif.GIF{LoopCount: loopCount, Config: config}
}

// AttachImage appends an Image to a GIF struct's set of frames.
func AttachImage(g *gif.GIF, img *image.Image) {
	AttachImageDelayed(g, img, 0)
}

// AttachImageDelayed appends an Image to a GIF struct's set of frames,
// as well as a delay measured in hundredths of a second
func AttachImageDelayed(g *gif.GIF, img *image.Image, delay int) {
	bounds := (*img).Bounds()
	// the gif struct uses PalettedImages, not Images
	palettedFrame := image.NewPaletted(bounds, g.Config.ColorModel.(color.Palette))
	draw.Draw(palettedFrame, palettedFrame.Rect, *img, bounds.Min, draw.Over)

	g.Image = append(g.Image, palettedFrame)
	g.Delay = append(g.Delay, delay)
}

// Populate provides a framework for populating an empty gif with frames
// according to a custom function defined by the user.
func Populate(g *gif.GIF, frames int, fg FrameGetter) {
	for step := 0; step < frames; step++ {
		frame := fg.GetFrame(step)
		AttachImage(g, frame)
	}
}

// WriteToFile encodes a GIF struct into a file of the given filename.
func WriteToFile(g *gif.GIF, filename string) error {
	file, fileErr := os.Create(filename)
	if fileErr != nil {
		return fileErr
	}

	encodeErr := gif.EncodeAll(file, g)
	if encodeErr != nil {
		return encodeErr
	}
	return nil
}
