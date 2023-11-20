package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
	"strconv"
)

func main() {
	for i := 1; i <= 560; i++ {
		LDfile, _ := os.Open("/Users/zane/Desktop/B301MM/hight/LD" + strconv.Itoa(i) + ".png")
		ULDfile, _ := os.Open("/Users/zane/Desktop/B301MM/low/ULD" + strconv.Itoa(i) + ".png")

		defer LDfile.Close()
		defer ULDfile.Close()

		img1, _ := png.Decode(LDfile)
		img2, _ := png.Decode(ULDfile)
		newImg := image.NewGray(image.Rect(0, 0, 512, 256))
		draw.Draw(newImg, image.Rect(0, 0, 256, 256), img1, image.Point{}, draw.Src)
		draw.Draw(newImg, image.Rect(256, 0, 512, 256), img2, image.Point{}, draw.Src)
		os.Mkdir("/Users/zane/Desktop/B301MM/new", 0777)
		outputFile, _ := os.Create("/Users/zane/Desktop/B301MM/new/new" + strconv.Itoa(i) + ".png")
		defer outputFile.Close()
		png.Encode(outputFile, newImg)
		fmt.Println("已经处理完第" + strconv.Itoa(i) + "张图片")
	}

}
