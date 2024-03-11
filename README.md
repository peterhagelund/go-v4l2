# V4L2

V4L2 definitions for Go.

## Copyright and Licensing

Copyright (c) 2020-2024 Peter Hagelund

This software is licensed under the [MIT License](https://en.wikipedia.org/wiki/MIT_License)

See `LICENSE.txt`

## Installing

```bash
go get github.com/peterhagelund/go-v4l2
```

## Using

```go
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
```

### On an Ubuntu PC with a cheap USB Web camera, the output could be:

```
Driver....: uvcvideo
Card......: Streaming Webcam: Streaming Web
BusInfo...: usb-0000:00:14.0-8
```
```bash
$ ls -l *.jpeg
.rw-rw-r-- 130k peter 10 Mar 20:21 decode_1.jpeg
.rw-rw-r-- 129k peter 10 Mar 20:21 decode_2.jpeg
.rw-rw-r-- 129k peter 10 Mar 20:21 decode_3.jpeg
.rw-rw-r-- 129k peter 10 Mar 20:22 decode_4.jpeg
.rw-rw-r-- 129k peter 10 Mar 20:22 decode_5.jpeg
.rw-r--r-- 269k peter 10 Mar 20:21 grab_1.jpeg
.rw-r--r-- 268k peter 10 Mar 20:21 grab_2.jpeg
.rw-r--r-- 268k peter 10 Mar 20:21 grab_3.jpeg
.rw-r--r-- 268k peter 10 Mar 20:22 grab_4.jpeg
.rw-r--r-- 267k peter 10 Mar 20:22 grab_5.jpeg
```
