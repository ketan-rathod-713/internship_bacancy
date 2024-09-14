package main

import (
	"fmt"
	"image"
	"log"
	"os"

	_ "image/png"
)

var img chan image.Image = make(chan image.Image)

func main() {
	img, err := loadImage("../imageprocessingpipeline/image3.jpeg")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(img)
}

func loadImage(filename string) (image.Image, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	img, formate, err := image.Decode(file)
	if err != nil {
		log.Println(formate)
		return nil, err
	}

	log.Println("Formate of image is", formate)

	return img, nil
}

func resizeImage() {

}

func saveImage() {

}
