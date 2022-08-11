package downsize

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"math/rand"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
	"golang.org/x/image/webp"
)

func ImageForInfiniteZoom(img string, scale float64, transition bool, transitionScale float64, convertSource bool) error {
	originalFile, err := os.Open(img)
	if err != nil {
		return fmt.Errorf("failed to open image: %v", err)
	}
	defer originalFile.Close()

	var originalImage image.Image
	originalIsPNG := false
	ext := filepath.Ext(originalFile.Name())
	if ext == ".png" {
		originalIsPNG = true
		originalImage, err = png.Decode(originalFile)
		if err != nil {
			return fmt.Errorf("failed to decode PNG image while reading it: %v", err)
		}
	} else if ext == ".jpg" || ext == ".jpeg" {
		originalImage, err = jpeg.Decode(originalFile)
		if err != nil {
			return fmt.Errorf("failed to decode JPEG image while reading it: %v", err)
		}
	} else if ext == ".webp" {
		originalImage, err = webp.Decode(originalFile)
		if err != nil {
			return fmt.Errorf("failed to decode WEBP image while reading it: %v", err)
		}
	} else {
		return fmt.Errorf("unsupported image file type")
	}

	// Convert the original image in place to a PNG for easy rendering later
	if !originalIsPNG && convertSource {
		converted, err := os.Create(strings.TrimSuffix(img, filepath.Ext(img)) + ".png")
		if err != nil {
			return fmt.Errorf("failed to save original converted image to PNG: %v", err)
		}
		defer converted.Close()

		if err = png.Encode(converted, originalImage); err != nil {
			return fmt.Errorf("failed to write converted PNG image to file: %v", err)
		}

		// Clean up the original non-PNG file for good housekeeping
		if err = os.Remove(originalFile.Name()); err != nil {
			return fmt.Errorf("failed to clean up original non-PNG file: %v", err)
		}
	}

	width := uint(float64(originalImage.Bounds().Max.X) * scale)
	resized := resize.Resize(width, 0, originalImage, resize.Lanczos3)

	b := resized.Bounds().Add(image.Pt(
		originalImage.Bounds().Max.X/2-(resized.Bounds().Max.X/2),
		originalImage.Bounds().Max.Y/2-(resized.Bounds().Max.Y/2)))
	canvas := image.NewRGBA(image.Rect(0, 0, originalImage.Bounds().Max.X, originalImage.Bounds().Max.Y))
	draw.Draw(
		canvas,
		b,
		resized,
		image.Point{0, 0},
		draw.Over)

	// Create a fuzzy transition at the edges of image i to enable
	//    future DALL-E 2 inpainting to create a more gradual seam between
	//    frames i and i+1. We do this by randomly zeroing out pixels less
	//    frequently the farther we move from the edge inward.
	//    CREDIT: Concept / algorithm inspired by pi314159265358978's p5.js
	//    downsizing app @ https://editor.p5js.org/Pi_p5/sketches/qAqoieAhx
	if transition {
		transitionSize := int(transitionScale * float64(resized.Bounds().Max.X))
		for i := b.Min.X; i <= b.Max.X; i++ {
			for j := b.Min.Y; j <= b.Min.Y+transitionSize; j++ {
				if rand.Float64() > float64(j-b.Min.Y)/float64(transitionSize) {
					canvas.Set(i, j, color.RGBA{0, 0, 0, 0})
				}
			}
		}
		for i := b.Min.X; i <= b.Min.X+transitionSize; i++ {
			for j := b.Min.Y; j <= b.Max.Y; j++ {
				if rand.Float64() > float64(i-b.Min.X)/float64(transitionSize) {
					canvas.Set(i, j, color.RGBA{0, 0, 0, 0})
				}
			}
		}
		for i := b.Min.X; i <= b.Max.X; i++ {
			for j := b.Max.Y; j >= b.Max.Y-transitionSize; j-- {
				if rand.Float64() > float64(b.Max.Y-j)/float64(transitionSize) {
					canvas.Set(i, j, color.RGBA{0, 0, 0, 0})
				}
			}
		}
		for i := b.Max.X; i >= b.Max.X-transitionSize; i-- {
			for j := b.Min.Y; j <= b.Max.Y; j++ {
				if rand.Float64() > float64(b.Max.X-i)/float64(transitionSize) {
					canvas.Set(i, j, color.RGBA{0, 0, 0, 0})
				}
			}
		}
	}

	downsizedOut, err := os.Create(strings.TrimSuffix(img, filepath.Ext(img)) + "_downsized.png")
	if err != nil {
		return fmt.Errorf("failed to save downsized image: %v", err)
	}
	defer downsizedOut.Close()

	if err = png.Encode(downsizedOut, canvas); err != nil {
		return fmt.Errorf("failed to write downsized image to file: %v", err)
	}

	return nil
}
