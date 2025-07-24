package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"strings"
	"github.com/dayiamin/imgtool/internal"
)

var (
	inputPath   string
	outputPath  string
	format      string
	suffix      string
	quality     int
	recursive   bool
)

func isImageFile(path string) bool {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".jpg", ".jpeg", ".png":
		return true
	default:
		return false
	}
}

func init() {
	compressCmd.Flags().StringVarP(&inputPath, "input", "i", "", "Path to input file or directory")
	compressCmd.Flags().StringVarP(&outputPath, "output", "o", "./output", "Path to output folder")
	compressCmd.Flags().StringVarP(&format, "format", "f", "original", "Output format: original or webp")
	compressCmd.Flags().StringVarP(&suffix, "suffix", "s", "_compressed", "Suffix to add to compressed file names")
	compressCmd.Flags().IntVarP(&quality, "quality", "q", 80, "Compression quality (for lossy formats)")
	compressCmd.Flags().BoolVarP(&recursive, "recursive", "r", false, "Scan directories recursively")
	compressCmd.MarkFlagRequired("input")
	rootCmd.AddCommand(compressCmd)
}

var compressCmd = &cobra.Command{
	Use:   "compress",
	Short: "Compress images from a file or folder",
	Run: func(cmd *cobra.Command, args []string) {
	fmt.Println("🔧 Compressing images...")

	info, err := os.Stat(inputPath)
	if err != nil {
		fmt.Println("❌ Invalid input path:", err)
		return
	}

	if info.IsDir() {
		err := filepath.Walk(inputPath, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() && path != inputPath && !recursive {
				return filepath.SkipDir
			}
			if !info.IsDir() {
				if isImageFile(path) {
					fmt.Println("🖼 Compressing:", path)
					err := internal.ProcessImage(path, outputPath, suffix, format, quality)
					if err != nil {
						fmt.Println("⚠️ Failed:", err)
					}
				}
			}
			return nil
		})
		if err != nil {
			fmt.Println("❌ Error:", err)
		}
	} else {
		err := internal.ProcessImage(inputPath, outputPath, suffix, format, quality)
		if err != nil {
			fmt.Println("❌ Error:", err)
		}
	}
},

}
