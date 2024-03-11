package v4l2

import (
	"io"
	"unsafe"

	"golang.org/x/sys/unix"
)

type Camera interface {
	io.Closer
	Path() string
	Driver() string
	Card() string
	BusInfo() string
	QueryCapabilities() (*Capability, error)
	EnumFormats(bufType BufType) ([]*FmtDesc, error)
	EnumFormatDescriptions(bufType BufType) ([]string, error)
	HasFormat(bufType BufType, pixFormat PixFmt) (bool, error)
	HasFormatDescription(bufType BufType, description string) (bool, error)
	EnumFrameSizes(pixFormat PixFmt) ([]*FrameSizeEnum, error)
	QueryControls() ([]*QueryCtrl, error)
	GetControl(id CtrlID) (*Control, error)
	SetControl(control *Control) error
	QueryMenus(id CtrlID) ([]*QueryMenu, error)
	StreamOn() error
	StreamOff() error
	GrabFrame() ([]byte, error)
}

type CameraConfig struct {
	Path      string
	BufType   BufType
	PixFormat PixFmt
	Width     uint32
	Height    uint32
	Memory    Memory
	BufCount  uint32
}

type camera struct {
	path      string
	fd        int
	driver    string
	card      string
	busInfo   string
	bufType   BufType
	pixFormat PixFmt
	memory    Memory
	width     uint32
	height    uint32
	buffers   [][]byte
}

func (c *camera) Close() error {
	if err := unix.Close(c.fd); err != nil {
		return err
	}
	c.fd = -1
	return nil
}

func (c *camera) Path() string {
	return c.path
}

func (c *camera) Driver() string {
	return c.driver
}

func (c *camera) Card() string {
	return c.card
}

func (c *camera) BusInfo() string {
	return c.busInfo
}

func (c *camera) QueryCapabilities() (*Capability, error) {
	return QueryCapabilities(c.fd)
}

func (c *camera) EnumFormats(bufType BufType) ([]*FmtDesc, error) {
	return EnumFormats(c.fd, bufType)
}

func (c *camera) EnumFormatDescriptions(bufType BufType) ([]string, error) {
	fmtDescs, err := c.EnumFormats(bufType)
	if err != nil {
		return nil, err
	}
	descriptions := make([]string, len(fmtDescs))
	for i, fmtDesc := range fmtDescs {
		descriptions[i] = BytesToString(fmtDesc.Description[:])
	}
	return descriptions, nil
}

func (c *camera) HasFormat(bufType BufType, pixFormat PixFmt) (bool, error) {
	fmtDescs, err := c.EnumFormats(bufType)
	if err != nil {
		return false, err
	}
	for _, fmtDesc := range fmtDescs {
		if fmtDesc.PixFormat == pixFormat {
			return true, nil
		}
	}
	return false, nil
}

func (c *camera) HasFormatDescription(bufType BufType, description string) (bool, error) {
	fmtDescs, err := c.EnumFormats(bufType)
	if err != nil {
		return false, err
	}
	for _, fmtDesc := range fmtDescs {
		if BytesToString(fmtDesc.Description[:]) == description {
			return true, nil
		}
	}
	return false, nil
}

func (c *camera) EnumFrameSizes(pixFormat PixFmt) ([]*FrameSizeEnum, error) {
	return EnumFrameSizes(c.fd, pixFormat)
}

func (c *camera) SupportsFrameSize(pixFormat PixFmt, width, height uint32) (bool, error) {
	frameSizeEnums, err := c.EnumFrameSizes(pixFormat)
	if err != nil {
		return false, err
	}
	for _, frameSizeEnum := range frameSizeEnums {
		if frameSizeEnum.Type == FrmSizeTypeDiscrete {
			discrete := (*FrameSizeDiscrete)(unsafe.Pointer(&frameSizeEnum.M))
			if discrete.Width == width && discrete.Height == height {
				return true, nil
			}
		} else {
			stepwise := (*FrameSizeStepwise)(unsafe.Pointer(&frameSizeEnum.M))
			if width < stepwise.MinWidth || width > stepwise.MaxWidth || height < stepwise.MinHeight || height > stepwise.MaxHeight {
				continue
			}
			if frameSizeEnum.Type == FrmSizeTypeContinuous {
				return true, nil
			} else {
				deltaWidth := width - stepwise.MinWidth
				deltaHeight := height - stepwise.MinHeight
				if deltaWidth%stepwise.StepWidth != 0 || deltaHeight%stepwise.StepHeight != 0 {
					continue
				}
				return true, nil
			}
		}
	}
	return false, nil
}

func (c *camera) QueryControls() ([]*QueryCtrl, error) {
	return QueryControls(c.fd)
}

func (c *camera) GetControl(id CtrlID) (*Control, error) {
	return GetControl(c.fd, id)
}

func (c *camera) SetControl(control *Control) error {
	return SetControl(c.fd, control)
}

func (c *camera) QueryMenus(id CtrlID) ([]*QueryMenu, error) {
	return QueryMenus(c.fd, id)
}

func (c *camera) StreamOn() error {
	return StreamOn(c.fd, c.bufType)
}

func (c *camera) StreamOff() error {
	return StreamOff(c.fd, c.bufType)
}

func (c *camera) GrabFrame() ([]byte, error) {
	return GrabFrame(c.fd, c.bufType, c.memory, c.buffers)
}

func NewCamera(config *CameraConfig) (Camera, error) {
	var err error
	fd, err := unix.Open(config.Path, unix.O_RDWR, 0)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			unix.Close(fd)
		}
	}()
	capabilities, err := QueryCapabilities(fd)
	if err != nil {
		return nil, err
	}
	driver := BytesToString(capabilities.Driver[:])
	card := BytesToString(capabilities.Card[:])
	busInfo := BytesToString(capabilities.BusInfo[:])
	width, height, err := SetFormat(fd, config.BufType, config.PixFormat, config.Width, config.Height)
	if err != nil {
		return nil, err
	}
	count, err := RequestDriverBuffers(fd, config.BufCount, config.BufType, config.Memory)
	if err != nil {
		return nil, err
	}
	var buffers [][]byte
	if config.Memory == MemoryMmap {
		buffers, err = MmapBuffers(fd, count, config.BufType)
		if err != nil {
			return nil, err
		}
	}
	return &camera{
		path:      config.Path,
		fd:        fd,
		driver:    driver,
		card:      card,
		busInfo:   busInfo,
		bufType:   config.BufType,
		pixFormat: config.PixFormat,
		memory:    config.Memory,
		width:     width,
		height:    height,
		buffers:   buffers,
	}, nil
}
