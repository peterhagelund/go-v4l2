package v4l2

import (
	"testing"
	"unsafe"

	"golang.org/x/sys/unix"
)

func TestQueryCapabilit(t *testing.T) {
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
				t.Fatal("zero-width frame size")
			}
			if frameSizeDiscrete.Height == 0 {
				t.Fatal("zero-height frame size")
			}
		} else {
			frameSizeStepwise := (*FrameSizeStepwise)(unsafe.Pointer(&frameSizeEnum.M))
			if frameSizeStepwise.MinHeight == 0 {
				t.Fatal("zero-min-height")
			}
			if frameSizeStepwise.MaxHeight == 0 {
				t.Fatal("zero-max-height")
			}
			if frameSizeStepwise.MinWidth == 0 {
				t.Fatal("zero-min-width")
			}
			if frameSizeStepwise.MaxWidth == 0 {
				t.Fatal("zero-max-width")
			}
		}
	}
}
