package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	mainPath := "../test-organizer"

	createDir("images")
	createDir("videos")

	files, err := readFiles(mainPath)
	if err != nil {
		log.Fatalf("err reading files: %v", err)
	}

	images, videos := categorizeFiles(files)

	for _, vid := range videos {
		sourcePath := filepath.Join(mainPath, vid)
		destPath := filepath.Join("videos", vid)
		if err := os.Rename(sourcePath, destPath); err != nil {
			fmt.Printf("failed to move video %s: %v\n", vid, err)
		} else {
			fmt.Printf("moved vid: %s\n", vid)
		}
	}

	for _, img := range images {
		sourcePath := filepath.Join(mainPath, img)
		destPath := filepath.Join("images", img)
		if err := os.Rename(sourcePath, destPath); err != nil {
			fmt.Printf("failed to move img %s: %v\n", img, err)
		} else {
			fmt.Printf("moved img: %s\n", img)
		}
	}
}

func createDir(dirName string) {
	if err := os.Mkdir(dirName, os.ModePerm); err != nil {
		log.Fatal(err)
	}
}

func readFiles(path string) ([]os.FileInfo, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open dir: %w", err)
	}
	defer f.Close()

	files, err := f.Readdir(0)
	if err != nil {
		return nil, fmt.Errorf("failed to read dir: %w", err)
	}

	return files, nil
}

func categorizeFiles(files []os.FileInfo) (images, videos []string) {
	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".jpg") || strings.HasSuffix(file.Name(), ".png") {
			images = append(images, file.Name())
		}
		if strings.HasSuffix(file.Name(), ".mp4") {
			videos = append(videos, file.Name())
		}
	}
	return images, videos
}
