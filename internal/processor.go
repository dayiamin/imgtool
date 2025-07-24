package internal

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/chai2010/webp"
	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
	"time"
)


func ProcessImage(inputPath, outputDir, suffix, format string, quality int) error {
	startTime := time.Now()
	file, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	img, imgFormat, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("cannot decode image: %w", err)
	}

	file.Seek(0, 0)
	exifData, err := exif.Decode(file)
	if err == nil {
		orientation, _ := exifData.Get(exif.Orientation)
		if orientation != nil {
			orientationInt, _ := orientation.Int(0)
			img = correctOrientation(img, orientationInt)
		}
	}

	// creating output file
	base := filepath.Base(inputPath)
	ext := filepath.Ext(base)
	nameOnly := strings.TrimSuffix(base, ext)

	// output format
	var outExt string
	var encoded []byte

	switch strings.ToLower(format) {
	case "webp":
		outExt = ".webp"
		buf := new(bytes.Buffer)
		err = webp.Encode(buf, img, &webp.Options{
			// Lossless: true,
			Quality: float32(quality),
		})
		encoded = buf.Bytes()
	case "original":
		switch imgFormat {
		case "jpeg", "jpg":
			outExt = ".jpg"
			buf := new(bytes.Buffer)
			err = jpeg.Encode(buf, img, &jpeg.Options{Quality: quality})
			encoded = buf.Bytes()
		case "png":
			outExt = ".png"
			buf := new(bytes.Buffer)
			enc := png.Encoder{CompressionLevel: png.BestCompression}
			err = enc.Encode(buf, img)
			encoded = buf.Bytes()
		default:
			return fmt.Errorf("unsupported original format: %s", imgFormat)
		}
	default:
		return fmt.Errorf("unknown output format: %s", format)
	}

	if err != nil {
		return fmt.Errorf("encoding error: %w", err)
	}

	

	// output path
	os.MkdirAll(outputDir, 0755)
	outputFile := filepath.Join(outputDir, nameOnly+suffix+outExt)
	err = os.WriteFile(outputFile, encoded, 0644)
	// Show compression stats
	originalStat, _ := os.Stat(inputPath)
	newStat, _ := os.Stat(outputFile)
	originalSize := float64(originalStat.Size())
	newSize := float64(newStat.Size())
	reduction := (originalSize - newSize) / originalSize * 100
	
	fmt.Printf("✅ %s →  (%.2f%% smaller)\t", base, reduction)
	fmt.Printf("Processing Time: %d ms\n", time.Since(startTime).Milliseconds())
	return err
}


// fixing image orientation
func correctOrientation(img image.Image, orientation int) image.Image {
	switch orientation {
	case 2:
		return imaging.FlipH(img)
	case 3:
		return imaging.Rotate180(img)
	case 4:
		return imaging.Rotate180(imaging.FlipH(img))
	case 5:
		return imaging.Rotate270(imaging.FlipH(img))
	case 6:
		return imaging.Rotate270(img)
	case 7:
		return imaging.Rotate90(imaging.FlipH(img))
	case 8:
		return imaging.Rotate90(img)
	default:
		return img
	}
}