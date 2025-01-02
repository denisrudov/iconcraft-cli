package main

import (
	"bytes"
	"fmt"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"
	"strings"
)

func main() {
	search := NewSearch()
	icons := search.Perform("user")

	if len(icons) > 0 {
		icon := icons[0]
		svgString := strings.ReplaceAll(icon.Svg, "currentColor", "#FFFFFF")
		toImage, err := renderSVGToImage([]byte(svgString), 40, 40)
		if err != nil {
			log.Fatal(err)
		}
		renderInConsole(toImage)
	}

}

func renderSVGToImage(svgSource []byte, width, height int) (image.Image, error) {
	// Create a reader for the SVG source
	reader := bytes.NewReader(svgSource)

	// Read the SVG icon
	icon, err := oksvg.ReadIconStream(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to parse SVG: %v", err)
	}

	// Set the rendering dimensions
	icon.SetTarget(0, 0, float64(width), float64(height))

	// Create an RGBA image to render the SVG onto
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	// Create a rasterizer
	scanner := rasterx.NewScannerGV(width, height, img, img.Bounds())
	raster := rasterx.NewDasher(width, height, scanner)
	raster.SetColor(color.Black)

	// Render the SVG onto the image
	icon.Draw(raster, 1.0)

	return img, nil
}

func saveImageAsPNG(img image.Image, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		return fmt.Errorf("failed to encode PNG: %v", err)
	}

	return nil
}

func renderInConsole(img image.Image) {
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y += 2 { // Process two rows at a time
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			// Get the top pixel color
			topColor := img.At(x, y)
			r1, g1, b1, _ := topColor.RGBA()

			// Get the bottom pixel color (or use black if out of bounds)
			var r2, g2, b2 uint32
			if y+1 < bounds.Max.Y {
				bottomColor := img.At(x, y+1)
				r2, g2, b2, _ = bottomColor.RGBA()
			}

			// Convert colors to ANSI escape codes
			topANSI := getANSIColor(r1, g1, b1)
			bottomANSI := getANSIColor(r2, g2, b2)

			// Print the combined character using "▀" (upper half block)
			fmt.Printf("\x1b[%sm\x1b[%sm▀", bottomANSI, topANSI)
		}
		fmt.Println("\x1b[0m") // Reset ANSI colors at the end of the line
	}
}

// Helper function to convert RGB to ANSI color codes
func getANSIColor(r, g, b uint32) string {
	brightness := (r + g + b) / 3
	if brightness > 0xFFFF/2 {
		return "37" // White
	}
	return "30" // Black
}
