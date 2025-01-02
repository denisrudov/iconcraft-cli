package main

import (
	"bytes"
	"fmt"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
	"golang.design/x/clipboard"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"image"
	"image/color"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func camelCaseFromDash(input string) string {
	words := strings.Split(input, "-")
	c := cases.Title(language.Und)
	for i := 0; i < len(words); i++ {
		words[i] = c.String(words[i])
	}
	return strings.Join(words, "")
}

func renderSVGToImage(svgSource []byte, width, height int) (image.Image, error) {
	reader := bytes.NewReader(svgSource)

	icon, err := oksvg.ReadIconStream(reader)
	if err != nil {
		return nil, fmt.Errorf("failed to parse SVG: %v", err)
	}

	icon.SetTarget(0, 0, float64(width), float64(height))

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	scanner := rasterx.NewScannerGV(width, height, img, img.Bounds())
	raster := rasterx.NewDasher(width, height, scanner)
	raster.SetColor(color.Black)

	icon.Draw(raster, 1.0)

	return img, nil
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

func ClearConsole() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func copyToClipboard(text string) error {
	// Initialize the clipboard
	if err := clipboard.Init(); err != nil {
		return fmt.Errorf("failed to initialize clipboard: %w", err)
	}

	// Copy the text to the clipboard
	clipboard.Write(clipboard.FmtText, []byte(text))
	return nil
}
