# V4L2

V4L2 definitions for Go.

## Copyright and Licensing

Copyright (c) 2020-2023 Peter Hagelund

This software is licensed under the [MIT License](https://en.wikipedia.org/wiki/MIT_License)

See `LICENSE.txt`

## Installing

```bash
go get -u github.com/peterhagelund/go-v4l2
```

### Using modules

In `go.mod`:

```
require "github.com/peterhagelund/go-v4l2/pkg/v4l2" v0.5.0
```

## Using

```go
package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"os"
	"unsafe"

	"github.com/peterhagelund/go-v4l2/v4l2"
	"golang.org/x/sys/unix"
)

func main() {
	fd, err := unix.Open("/dev/video0", unix.O_RDWR, 0)
	if err != nil {
		panic(err)
	}
	defer unix.Close(fd)
	capability, err := v4l2.QueryCapabilities(fd)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Driver....: %s\n", v4l2.BytesToString(capability.Driver[:]))
	fmt.Printf("Card......: %s\n", v4l2.BytesToString(capability.Card[:]))
	fmt.Printf("BusInfo...: %s\n", v4l2.BytesToString(capability.BusInfo[:]))
	fmtDescs, err := v4l2.EnumFormats(fd, v4l2.BufTypeVideoCapture)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Formats:\n")
	for _, fmtDesc := range fmtDescs {
		fmt.Printf("\tFormat...: %s\n", v4l2.BytesToString(fmtDesc.Description[:]))
	}
	fmt.Printf("Frame sizes for JPEG:\n")
	frameSizeEnums, err := v4l2.EnumFrameSizes(fd, v4l2.PixFmtJPEG)
	if err != nil {
		panic(err)
	}
	for _, frameSizeEnum := range frameSizeEnums {
		if frameSizeEnum.Type == v4l2.FrmSizeTypeDiscrete {
			discrete := (*v4l2.FrameSizeDiscrete)(unsafe.Pointer(&frameSizeEnum.M))
			fmt.Printf("\tDiscrete:\n")
			fmt.Printf("\t\tWidth....: %d\n", discrete.Width)
			fmt.Printf("\t\tHeight...: %d\n", discrete.Height)
		} else {
			if frameSizeEnum.Type == v4l2.FrmSizeTypeStepwise {
				fmt.Printf("\tStepwise:\n")
			} else {
				fmt.Printf("\tContinuous:\n")
			}
			stepwise := (*v4l2.FrameSizeStepwise)(unsafe.Pointer(&frameSizeEnum.M))
			fmt.Printf("\t\tMinWidth.....: %d\n", stepwise.MinWidth)
			fmt.Printf("\t\tMaxWidth.....: %d\n", stepwise.MaxWidth)
			fmt.Printf("\t\tStepWidth....: %d\n", stepwise.StepWidth)
			fmt.Printf("\t\tMinHeight....: %d\n", stepwise.MinHeight)
			fmt.Printf("\t\tMaxHeight....: %d\n", stepwise.MaxHeight)
			fmt.Printf("\t\tStepHeight...: %d\n", stepwise.StepHeight)
		}
	}
	width, height, err := v4l2.SetFormat(fd, v4l2.BufTypeVideoCapture, v4l2.PixFmtJPEG, 1024, 768)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Frame size:\n")
	fmt.Printf("\tWidth....: %d\n", width)
	fmt.Printf("\tHeight...: %d\n", height)
	count, err := v4l2.RequestDriverBuffers(fd, 4, v4l2.BufTypeVideoCapture, v4l2.MemoryMmap)
	if err != nil {
		panic(err)
	}
	defer v4l2.RequestDriverBuffers(fd, 0, v4l2.BufTypeVideoCapture, v4l2.MemoryMmap)
	fmt.Printf("Driver buffer count...: %d\n", count)
	buffers, err := v4l2.MmapBuffers(fd, count, v4l2.BufTypeVideoCapture)
	if err != nil {
		panic(err)
	}
	defer v4l2.MunmapBuffers(buffers)
	if err := v4l2.StreamOn(fd, v4l2.BufTypeVideoCapture); err != nil {
		panic(err)
	}
	defer v4l2.StreamOff(fd, v4l2.BufTypeVideoCapture)
	frame, err := v4l2.GrabFrame(fd, v4l2.BufTypeVideoCapture, v4l2.MemoryMmap, buffers)
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile("test1.jpeg", frame, 0644); err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(frame)
	img, name, err := image.Decode(buffer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Image:\n")
	fmt.Printf("\tName.....: %s\n", name)
	fmt.Printf("\tWidth....: %d\n", img.Bounds().Dx())
	fmt.Printf("\tHeight...: %d\n", img.Bounds().Dy())
	file, err := os.Create("test2.jpeg")
	if err != nil {
		panic(err)
	}
	err = jpeg.Encode(file, img, nil)
	if err != nil {
		panic(err)
	}
}
```

### On an Ubuntu PC with a cheap USB Web camera, the output could be:

```
Driver....: uvcvideo
Card......: Streaming Webcam: Streaming Web
BusInfo...: usb-0000:00:14.0-8
Formats:
        Format...: Motion-JPEG
        Format...: YUYV 4:2:2
Frame sizes for JPEG:
Frame size:
        Width....: 1024
        Height...: 768
Driver buffer count...: 4
Image:
        Name.....: jpeg
        Width....: 1024
        Height...: 768
```bash
$ ls -l *.jpeg
.rw-r--r-- 129k peter  1 Mar 18:51 test1.jpeg
.rw-rw-r--  63k peter  1 Mar 18:51 test2.jpeg
```
