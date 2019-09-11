package main

import (
	"errors"
	"fmt"
	"image"
	"image/png"
	"os"
	"time"

	"github.com/kbinani/screenshot"
)

// save *image.RGBA to filePath with PNG format.
func save(img *image.RGBA, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return png.Encode(file, img)
}

func run() error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic")
			fmt.Println(err)
		}
	}()

	// Capture each displays.
	n := screenshot.NumActiveDisplays()
	if n <= 0 {
		return errors.New("no screen")
	}

	var all image.Rectangle = image.Rect(0, 0, 0, 0)

	for i := 0; i < n; i++ {
		bounds := screenshot.GetDisplayBounds(i)
		all = bounds.Union(all)

		_, err := screenshot.CaptureRect(bounds)
		if err != nil {
			fmt.Printf("err: %v", err)
		}
		fileName := fmt.Sprintf("%d_%dx%d.png", i, bounds.Dx(), bounds.Dy())
		//save(img, fileName)

		fmt.Printf("#%d : %v \"%s\"\n", i, bounds, fileName)
	}

	// Capture all desktop region into an image.
	fmt.Printf("%v\n", all)
	img, err := screenshot.Capture(all.Min.X, all.Min.Y, all.Dx(), all.Dy())
	if err != nil {
		return errors.New("capture error")
	}
	return save(img, "all.png")
}

func main() {
	for {
		err := run()
		if err != nil {
			time.Sleep(time.Second * 5)
		} else {
			time.Sleep(time.Second * 60 * 1)
		}
	}
}
