package v4l2

import (
	"testing"

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
}
