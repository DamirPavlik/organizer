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

	if err := os.Mkdir("images", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	if err := os.Mkdir("videos", os.ModePerm); err != nil {
		log.Fatal(err)
	}

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
		if strings.HasSuffix(v.Name(), ".jpg") {
			images = append(images, v.Name())
		}
		if strings.HasSuffix(v.Name(), ".png") {
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
		fmt.Println("img: ", img)
	}
}
