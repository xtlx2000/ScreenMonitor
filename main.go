package main

import (
	"fmt"
	"time"

	"github.com/go-vgo/robotgo"
)

func save() string {
	bitmap := robotgo.CaptureScreen()
	defer robotgo.FreeBitmap(bitmap)
	return robotgo.SaveBitmap(bitmap, "screen.png")
}

func main() {
	for {
		err := save()
		if err != "" {
			fmt.Printf("capture err: %v\n", err)
			time.Sleep(time.Second * 5)
		} else {
			fmt.Printf("capture ok\n")
			time.Sleep(time.Second * 60 * 1)
		}
	}
}
