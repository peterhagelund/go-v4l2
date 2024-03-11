package v4l2

import "testing"

func TestNewCamera(t *testing.T) {
	camera, err := NewCamera(&CameraConfig{
		Path:      "/dev/video0",
		BufType:   BufTypeVideoCapture,
		PixFormat: PixFmtMJPEG,
		Width:     1920,
		Height:    1280,
		Memory:    MemoryMmap,
		BufCount:  4,
	})
	if err != nil {
		t.Fatal("unable to create new camera")
	}
	camera.Close()
}
