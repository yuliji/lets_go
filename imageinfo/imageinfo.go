package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"os"

	"github.com/fatih/color"
)

func main() {
	fileName := os.Args[1]
	printImageDimension(fileName)
}
func printImageDimension(imagePath string) {
	file, err := os.Open(imagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	defer file.Close()

	image, format, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v\n", imagePath, err)
	}

	color.Green("format: %s", format)

	color.Green("width: %d", image.Width)
	color.Green("height: %d", image.Height)
}
