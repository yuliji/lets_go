package main

import (
	"github.com/fogleman/gg"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"
)

func main() {
	im, err := gg.LoadPNG("./gologo.png")
	if err != nil {
		panic(err)
	}

	size := im.Bounds().Size()

	dc := gg.NewContext(size.X, size.Y)
	dc.DrawImage(im, 0, 0)

	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic(err)
	}

	face := truetype.NewFace(font, &truetype.Options{Size: 48})

	dc.SetFontFace(face)
	dc.SetRGB(255, 0, 0)
	dc.DrawStringAnchored("666", float64(size.X), 0, 1, 1)
	dc.SavePNG("out.png")
}
