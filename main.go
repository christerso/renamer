package main

import (
	"flag"
	"fmt"
	"hash/crc32"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func getImageDimensions(filePath string) (int, int, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	imageData, _, err := image.DecodeConfig(file)
	if err != nil {
		return 0, 0, err
	}

	return imageData.Width, imageData.Height, nil
}

func generateCRC32Checksum(filePath string) (uint32, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	hash := crc32.NewIEEE()
	if _, err := io.Copy(hash, file); err != nil {
		return 0, err
	}

	return hash.Sum32(), nil
}

func getAspectRatio(width int, height int) string {
	gcd := func(a, b int) int {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	divisor := gcd(width, height)
	aspectRatio := fmt.Sprintf("%d:%d", width/divisor, height/divisor)
	return strings.ReplaceAll(aspectRatio, ":", "_")
}

func renameImages(dirPath string) error {
	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(info.Name())
		if ext == "" {
			return nil
		}

		width, height, err := getImageDimensions(path)
		if err != nil {
			return nil
		}

		crc32Checksum, err := generateCRC32Checksum(path)
		if err != nil {
			return nil
		}

		dateStr := time.Now().Format("20060102")
		aspectRatio := getAspectRatio(width, height)
		currentDirName := filepath.Base(filepath.Dir(path))
		newName := fmt.Sprintf("%s-%s-%dx%d-%s-%08x%s", currentDirName, dateStr, width, height, aspectRatio, crc32Checksum, ext)
		newPath := filepath.Join(filepath.Dir(path), newName)

		if err := os.Rename(path, newPath); err != nil {
			return err
		}

		return nil
	})
}

func main() {
	dirPtr := flag.String("dir", ".", "directory to process")
	flag.Parse()

	dirPath := *dirPtr

	if err := renameImages(dirPath); err != nil {
		fmt.Printf("error renaming images: %v\n", err)
	}
}
