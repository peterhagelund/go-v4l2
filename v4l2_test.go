package v4l2

import (
	"bytes"
	"image"
	"testing"
	"unsafe"

	_ "image/jpeg"

	"golang.org/x/sys/unix"
)

func TestQueryCapabilities(t *testing.T) {
	fd, err := unix.Open("/dev/video0", unix.O_RDWR, 0)
	if err != nil {
		t.Fatal("unable to open device")
	}
	defer unix.Close(fd)
	capability, err := QueryCapabilities(fd)
	if err != nil {
		t.Fatal("unable to query capabilities")
	}
	if capability == nil {
		t.Fatal("nil capability returned")
	}
	driver := BytesToString(capability.Driver[:])
	if len(driver) == 0 {
		t.Fatal("no driver")
	}
	card := BytesToString(capability.Card[:])
	if len(card) == 0 {
		t.Fatal("no card")
	}
	busInfo := BytesToString(capability.BusInfo[:])
	if len(busInfo) == 0 {
		t.Fatal("no bus info")
	}
	if capability.Capabilities&CapVideoCapture == 0x00000000 {
		t.Fatal("video capture not a capability")
	}
}

func TestEnumFormats(t *testing.T) {
	fd, err := unix.Open("/dev/video0", unix.O_RDWR, 0)
	if err != nil {
		t.Fatal("unable to open device")
	}
	defer unix.Close(fd)
	fmtDescs, err := EnumFormats(fd, BufTypeVideoCapture)
	if err != nil {
		t.Fatal("uanble to enumerate formats")
	}
	if fmtDescs == nil {
		t.Fatal("nil format descriptors returned")
	}
	if len(fmtDescs) == 0 {
		t.Fatal("no format descriptors returned")
	}
	found := false
	for _, fmtDesc := range fmtDescs {
		description := BytesToString(fmtDesc.Description[:])
		if description == "JFIF JPEG" {
			found = true
			break
		}
	}
	if !found {
		t.Fatal("well-known format not present")
	}
}

func TestEnumFrameSizes(t *testing.T) {
	fd, err := unix.Open("/dev/video0", unix.O_RDWR, 0)
	if err != nil {
		t.Fatal("unable to open device")
	}
	defer unix.Close(fd)
	frameSizeEnums, err := EnumFrameSizes(fd, PixFmtJPEG)
	if err != nil {
		t.Fatal("unable to enumerate frame sizes")
	}
	if frameSizeEnums == nil {
		t.Fatal("nil frame size enums returned")
	}
	if len(frameSizeEnums) == 0 {
		t.Fatal("no frame size enums returned")
	}
	for _, frameSizeEnum := range frameSizeEnums {
		if frameSizeEnum.Type == FrmSizeTypeDiscrete {
			frameSizeDiscrete := (*FrameSizeDiscrete)(unsafe.Pointer(&frameSizeEnum.M))
			if frameSizeDiscrete.Width == 0 {
				t.Fatal("zero width frame size")
			}
			if frameSizeDiscrete.Height == 0 {
				t.Fatal("zero height frame size")
			}
		} else {
			frameSizeStepwise := (*FrameSizeStepwise)(unsafe.Pointer(&frameSizeEnum.M))
			if frameSizeStepwise.MinHeight == 0 {
				t.Fatal("zero min height")
			}
			if frameSizeStepwise.MaxHeight == 0 {
				t.Fatal("zero max height")
			}
			if frameSizeStepwise.MinWidth == 0 {
				t.Fatal("zero min width")
			}
			if frameSizeStepwise.MaxWidth == 0 {
				t.Fatal("zero max width")
			}
		}
	}
}

func TestSetFormat(t *testing.T) {
	fd, err := unix.Open("/dev/video0", unix.O_RDWR, 0)
	if err != nil {
		t.Fatal("unable to open device")
	}
	defer unix.Close(fd)
	width, height, err := SetFormat(fd, BufTypeVideoCapture, PixFmtJPEG, 1024, 768)
	if err != nil {
		t.Fatal("unable to set format")
	}
	if width == 0 {
		t.Fatal("zero width returned")
	}
	if height == 0 {
		t.Fatal("zero height returned")
	}
}

func TestRequestDriverBuffers(t *testing.T) {
	fd, err := unix.Open("/dev/video0", unix.O_RDWR, 0)
	if err != nil {
		t.Fatal("unable to open device")
	}
	defer unix.Close(fd)
	count, err := RequestDriverBuffers(fd, 4, BufTypeVideoCapture, MemoryMmap)
	if err != nil {
		t.Fatal("unable to request driver buffers")
	}
	if count == 0 {
		t.Fatal("no driver buffers available")
	}
	_, err = RequestDriverBuffers(fd, 0, BufTypeVideoCapture, MemoryMmap)
	if err != nil {
		t.Fatal("unable to adjust requested driver buffers down to zero")
	}
}

func TestGrabFrame(t *testing.T) {
	fd, err := unix.Open("/dev/video0", unix.O_RDWR, 0)
	if err != nil {
		t.Fatal("unable to open device")
	}
	defer unix.Close(fd)
	if _, _, err := SetFormat(fd, BufTypeVideoCapture, PixFmtJPEG, 1024, 768); err != nil {
		t.Fatal("unable to set format")
	}
	count, err := RequestDriverBuffers(fd, 4, BufTypeVideoCapture, MemoryMmap)
	if err != nil {
		t.Fatal("unable to request driver buffers")
	}
	defer RequestDriverBuffers(fd, 0, BufTypeVideoCapture, MemoryMmap)
	buffers, err := MmapBuffers(fd, count, BufTypeVideoCapture)
	if err != nil {
		t.Fatal("unable to mmap buffers")
	}
	defer MunmapBuffers(buffers)
	if err := StreamOn(fd, BufTypeVideoCapture); err != nil {
		t.Fatal("unable to turn on streaming")
	}
	defer StreamOff(fd, BufTypeVideoCapture)
	frame, err := GrabFrame(fd, BufTypeVideoCapture, MemoryMmap, buffers)
	if err != nil {
		t.Fatal("unable to grab frame")
	}
	if frame == nil {
		t.Fatal("nil frame returned")
	}
	if len(frame) == 0 {
		t.Fatal("empty frame returned")
	}
	buffer := bytes.NewBuffer(frame)
	image, name, err := image.Decode(buffer)
	if err != nil {
		t.Fatal("unable to decode frame")
	}
	if name != "jpeg" {
		t.Fatal("returned frame is not a JPEG image")
	}
	bounds := image.Bounds()
	if bounds.Dx() != 1024 || bounds.Dy() != 768 {
		t.Fatal("image has incorrect size")
	}
}

func TestBytesToString(t *testing.T) {
	b1 := [...]byte{}
	s1 := BytesToString(b1[:])
	if len(s1) != 0 {
		t.Fatal("empty slice did not yield zero length string")
	}
	b2 := [...]byte{0x41, 0x42, 0x43, 0x31, 0x32, 0x33}
	s2 := BytesToString(b2[:])
	if s2 != "ABC123" {
		t.Fatal("incorrect string returned")
	}
}
