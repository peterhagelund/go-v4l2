// Copyright (c) 2020-2024 Peter Hagelund
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"
	"time"

	"github.com/peterhagelund/go-v4l2/v4l2"
)

func main() {
	camera, err := v4l2.NewCamera(&v4l2.CameraConfig{
		Path:      "/dev/video0",
		BufType:   v4l2.BufTypeVideoCapture,
		PixFormat: v4l2.PixFmtMJPEG,
		Width:     1920,
		Height:    1280,
		Memory:    v4l2.MemoryMmap,
		BufCount:  4,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer camera.Close()
	fmt.Printf("Driver....: %s\n", camera.Driver())
	fmt.Printf("Card......: %s\n", camera.Card())
	fmt.Printf("BusInfo...: %s\n", camera.BusInfo())
	if err := camera.StreamOn(); err != nil {
		log.Fatal(err)
	}
	defer camera.StreamOff()
	for i := 1; i <= 5; i++ {
		frame, err := camera.GrabFrame()
		if err != nil {
			log.Fatal(err)
		}
		if err := os.WriteFile(fmt.Sprintf("grab_%d.jpeg", i), frame, 0644); err != nil {
			log.Fatal(err)
		}
		buffer := bytes.NewBuffer(frame)
		img, _, err := image.Decode(buffer)
		if err != nil {
			log.Fatal(err)
		}
		file, err := os.Create(fmt.Sprintf("decode_%d.jpeg", i))
		if err != nil {
			log.Fatal(err)
		}
		err = jpeg.Encode(file, img, nil)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second)
	}
}
