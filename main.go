package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var images []string
	var videos []string
	mainPath := "../test-organizer"

	createDir("images")
	createDir("videos")

	f, err := os.Open(mainPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, v := range files {
		if strings.HasSuffix(v.Name(), ".jpg") || strings.HasSuffix(v.Name(), ".png") {
			images = append(images, v.Name())
		}
		if strings.HasSuffix(v.Name(), ".mp4") {
			videos = append(videos, v.Name())
		}
	}

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
