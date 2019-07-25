package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	widthPtr := flag.Int("width", 0, "width")
	heightPtr := flag.Int("height", 0, "height")
	flag.Parse()
	fromImg := flag.Args()[0]
	file, err := os.Open(fromImg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	defer file.Close()

	image, _, err := image.Decode(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", fromImg, err)
	}
	m := resize.Resize(uint(*widthPtr), uint(*heightPtr), image, resize.Lanczos3)

	out, err := os.Create("out.png")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	png.Encode(out, m)
}
