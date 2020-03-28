// Copyright (c) 2020 Peter Hagelund
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

package v4l2

import "syscall"

const (
	BufCapSupportsMMap uint32 = 1 << iota
	BufCapSupportsUserPtr
	BufCapSupportsDMABuf
	BufCapSupportsRequests
	BufCapSupportsOrphanedBufs
	BufCapSupportsM2MHoldCaptureBuf
)

const (
	BufTypeVideoCapture uint32 = iota + 1
	BufTypeVideoOutput
	BufTypeVideoOverlay
	BufTypeVBICapture
	BufTypeVBIOutput
	BufTypeSlicedVBICapture
	BufTypeSlicedVBIOutput
	BufTypeVideoOutputOverlay
	BufTypeVideCaptureMPlane
	BufTypeVideOutputMPlane
	BufTypeSDRCapture
	BufTypeSDROutput
	BufTypeMetaCapture
	BufTypeMetaOutput
)

const (
	CapVideoCapture uint32 = 1 << iota
	CapVideoOutput
	CapVideoOverlay
	_
	CapVBICapture
	CapVBIOutput
	CapSlicedVBICapture
	CapSlicedVBIOutput
	CapRDSCapture
	CapVideoOutputOverlay
	CapHwFreqSeek
	CapRDSOutput
	CapVideoCaptureMPlane
	CapVideoOutputMPlane
	CapVideoM2M
	CapVideoM2MMPlane
	CapTuner
	CapAudio
	CapRadio
	CapModulator
	CapSDRCapture
	CapExtPixFormat
	CapSDROutput
	CapMetaCapture
	capReadWrite
	CapAsyncIO
	CapStreaming
	CapMetaOutput
	CapTouch
	_
	_
	CapDeviceCaps
)

const (
	FieldAny uint32 = iota
	FieldNone
	FieldTop
	FieldBottom
	FieldInterlaced
	FieldSeqTB
	FieldSeqBT
	FieldAlternate
	FieldInterlacedTB
	FieldInterlacedBT
)

const (
	FrmSizeTypeDiscrete uint32 = iota + 1
	FrmSizeTypeContinuous
	FrmSizeTypeStepwise
)

const (
	// InputTypeTuner designates the input as a tuner.
	InputTypeTuner uint32 = iota + 1
	// InputTypeCamera designates the input as a camera.
	InputTypeCamera
	// InputTypeTouch designates the input as a touch device.
	InputTypeTouch
)

const (
	MemoryMMap uint32 = iota + 1
	MemoryUserPtr
	MemoryOverlay
	MemoryDMABuf
)

const (
	TcFlagUserDefined uint32 = 1 << iota
	TcFlagDropFrame
	// UserBits ??? Flags or ...?
)

const (
	TcType24Fps uint32 = iota + 1
	TcType25Fps
	TcType30Fps
	TcType50Fps
	TcType60Fps
)

const (
	// VidIocQueryCap queries the capabilities.
	VidIocQueryCap uint32 = 0x80685600
	// VidIocReserved is a reserved operation.
	VidIocReserved = 0x00005601
	// VidIocEnumFmt enumerates formats.
	VidIocEnumFmt = 0xc0405602
	// VidIocGFmt gets a specific format.
	VidIocGFmt = 0xc0d05604
	// VidIocSFmt sets a specific format.
	VidIocSFmt      = 0xc0d05605
	VidIocReqBufs   = 0xc0145608
	VidIocQueryBuf  = 0xc0585609
	VidIocGFBuf     = 0x8030560a
	VidIocSFbuf     = 0x4030560b
	VidIocOverlay   = 0x4004560e
	VidIocQBuf      = 0xc058560f
	VidIocExpBuf    = 0xc0405610
	VidIocDQBUF     = 0xc0585611
	VidIocSTREAMON  = 0x40045612
	VidIocSTREAMOFF = 0x40045613
	VidIocGParm     = 0xc0cc5615
	VidIocSParm     = 0xc0cc5616
	VidIocGStd      = 0x80085617
	VidIocSStd      = 0x40085618
	VidIocEnumStd   = 0xc0485619
	// VidIocEnumInput enumerates inputs.
	VidIocEnumInput          = 0xc050561a
	VidIocGCtrl              = 0xc008561b
	VidIocSCtrl              = 0xc008561c
	VidIocGTuner             = 0xc054561d
	VidIocSTuner             = 0x4054561e
	VidIocGAudio             = 0x80345621
	VidIocSAudio             = 0x40345622
	VidIocQueryCtrl          = 0xc0445624
	VidIocQueryMenu          = 0xc02c5625
	VidIocGInput             = 0x80045626
	VidIocSInput             = 0xc0045627
	VidIocGEDID              = 0xc0285628
	VidIocSEDID              = 0xc0285629
	VidIocGOutput            = 0x8004562e
	VidIocSOutput            = 0xc004562f
	VidIocEnumOutput         = 0xc0485630
	VidIocGAudOut            = 0x80345631
	VidIocSAudOut            = 0x40345632
	VidIocGModulator         = 0xc0445636
	VidIocSModulator         = 0x40445637
	VidIocGFrequency         = 0xc02c5638
	VidIocSFrequency         = 0x402c5639
	VidIocCropCap            = 0xc02c563a
	VidIocGCcrop             = 0xc014563b
	VidIocSCcrop             = 0x4014563c
	VidIocGJpegComp          = 0x808c563d
	VidIocSJpegComp          = 0x408c563e
	VidIocQueryStd           = 0x8008563f
	VidIocTryFmt             = 0xc0d05640
	VidIocEnumAudio          = 0xc0345641
	VidIocEnumAudOut         = 0xc0345642
	VidIocGPriority          = 0x80045643
	VidIocSPriority          = 0x40045644
	VidIocGSlicdeVBICap      = 0xc0745645
	VidIocLogStatus          = 0x5646
	VidIocGExtCtrls          = 0xc0205647
	VidIocSExtCtrls          = 0xc0205648
	VidIocTryExtCtrls        = 0xc0205649
	VidIocEnumFrameSizes     = 0xc02c564a
	VidIocEnumFrameIntervals = 0xc034564b
	VidIocGEncIndex          = 0x8818564c
	VidIocEncoderCmd         = 0xc028564d
	VidIocTryEncoderCmd      = 0xc028564e
)

// StdID is the standard ID type.
type StdID uint64

// Audio is an encapsulation of a set of audio attributes.
type Audio struct {
	Index      uint32
	Name       [32]byte
	Capability uint32
	Mode       uint32
	Reserved   [2]uint32
}

// Buffer is an encapsulation of a buffer.
type Buffer struct {
	Index     uint32
	Type      uint32
	BytesUsed uint32
	Flags     uint32
	Field     uint32
	Timestamp syscall.Timeval
	Timecode  Timecode
	Sequence  uint32
	Memory    uint32
	M         uint32 // Union
	Length    uint32
	Reserved2 uint32
	RequestFD uint32
}

// Capabilities is an encapsulation of device capabilities.
type Capabilities struct {
	Driver       [16]byte
	Card         [32]byte
	BusInfo      [32]byte
	Version      uint32
	Capabilities uint32
	DeviceCaps   uint32
	Reserved     [3]uint32
}

type CtrlFwhtparams struct {
	BackwardRefTS uint64
	Version       uint32
	Width         uint32
	Height        uint32
	Flags         uint32
	ColorSpace    uint32
	XferFunc      uint32
	YCbCrEnc      uint32
	Quantization  uint32
}

type CtrlH264DecodeParams struct {
	DPD                 [16]H264DPDEntry
	NumSlices           uint16
	NALRefIDC           uint16
	TopFieldOrderCnt    uint32
	BottomFieldOrderCnt uint32
	Flags               uint32
}

type CtrlHevcPps struct {
	NumExtraSliceHeaderBits      uint8
	InitQpMinus26                int8
	DiffCuQpDeltaDepth           uint8
	PpsCbQpOffset                int8
	PpsCrQpOffset                int8
	NumTileColumnsMinus1         uint8
	NumTileRowsMinus1            uint8
	ColumnWidthMinus1            [20]uint8
	RowHeightMinus1              [22]uint8
	PpsBetaOffsetDiv2            int8
	PpsTcOffsetDiv2              int8
	Log2ParallelMergeLevelMinus2 uint8
	Padding                      [4]uint8
	Flags                        uint64
}

type CtrlHevcSps struct {
	PicWidthInLumaSamples                uint16
	PicHeightInLumaSamples               uint16
	BitDepthLumaMinus8                   uint8
	BitDepthChromaMinus8                 uint8
	Log2MaxPicOrderCntLSBMinus4          uint8
	SpsMaxDecPicBufferingMinus1          uint8
	SpsMaxNumReorderPics                 uint8
	SpsMaxLatencyIncreasePlus1           uint8
	Log2MinLumaCodingBlockSizeMinus3     uint8
	Log2DiffMaxMinLumaCodingBlockSize    uint8
	Log2MinLumaTransformBlockSizeMinus2  uint8
	Log2DiffMaxMinLumaTransformBlockSize uint8
	MaxTransformHierarchyDepthInter      uint8
	MaxTransformHierarchyDepthIntra      uint8
	PCMSampleBitDepthLumaMinus1          uint8
	PCMSampleBitDepthChromaMinus1        uint8
	Log2MinPCMLumaCodingBlockSizeMinus3  uint8
	Log2DiffMaxMinPCMLumaCodingBlockSize uint8
	NumShortTermRefPicSets               uint8
	NumLongTermRefPicsSps                uint8
	ChromaFormatIDC                      uint8
	Flags                                uint64
}

type CtrlMpeg2Quantization struct {
	LoadIntraQuantiserMatrix          uint8
	LoadNonIntraQuantiserMatrix       uint8
	LoadChromaIntraQuantiserMatrix    uint8
	LoadChromaNonIntraQuantiserMatrix uint8
	IntraQuantiserMatrix              [64]uint8
	NonIntraQuantiserMatrix           [64]uint8
	ChromaIntraQuantiserMatrix        [64]uint8
	ChromaNonIntraQuantiserMatrix     [64]uint8
}

type CtrlMpeg2SliceParams struct {
	BitSize            uint32
	DataBitOffset      uint32
	Sequence           Mpeg2Sequence
	Picture            Mpeg2Picture
	ForwardRefTS       uint64
	QuantiserScaleCode uint32
}

type CtrlVP8FrameHeader struct {
	SegmentHeader         VP8SegmentHeader
	LoopfilterHeader      VP8LoopfilterHeader
	QuantHeader           VP8QuantizationHeader
	EntropyHeader         VP8EntropyHeader
	CoderState            VP8EntropyCoderState
	Width                 uint16
	Height                uint16
	HorizontalScale       uint8
	VerticalScalingFactor uint8
	Version               uint8
	PropSkipFalse         uint8
	PropIntra             uint8
	PropLast              uint8
	PropGF                uint8
	NumDCTParts           uint8
	FirstPartSize         uint32
	FirstPartHeaderBits   uint32
	DCTPartSize           [8]uint32
	LastFrameTS           uint64
	GoldenFrameTS         uint64
	AltFrameTS            uint64
	Flags                 uint64
}

// FmtDesc is an encapsulation of a format descriptor.
type FmtDesc struct {
	Index       uint32
	Type        uint32
	Flags       uint32
	Description [32]byte
	PixFormat   uint32
}

// Format is an encapsulation of the various formats.
type Format struct {
	Type    uint32
	RawData [256]byte
}

// FrameSizeDiscrete is an encapsulation of a discrete frame size.
type FrameSizeDiscrete struct {
	Width  uint32
	Height uint32
}

// FrameSizeEnum is an emcapsulation of frame size information.
type FrameSizeEnum struct {
	Index     uint32
	PixFormat uint32
	Type      uint32
	Stepwise  FrameSizeStepwise // Union with FrameSizeDiscrete
	Reserved  [2]uint32
}

// FrameSizeStepwise is an encapsulation of valid frame sizes in steps from min to max.
type FrameSizeStepwise struct {
	MinWidth   uint32
	MaxWidth   uint32
	StepWidth  uint32
	MinHeight  uint32
	MaxHeight  uint32
	StepHeight uint32
}

// Frequency is an encapsulation of a set of frquency attributes.
type Frequency struct {
	Tuner     uint32
	Type      uint32
	Frequency uint32
	Reserved  [4]uint32
}

type H264DPDEntry struct {
	ReferenceTS         uint64
	FrameNum            uint16
	PicNum              uint16
	TopFieldOrderCnt    int32
	BottomFieldOrderCnt int32
	Flags               uint32
}

type H264PredWeightTable struct {
	LumaLog2WeightDenom   uint16
	ChromaLog2WeightDenom uint16
	Weightfactors         [2]H264WeightFactors
}

type H264WeightFactors struct {
	LumaWeight   [32]int16
	LumaOffset   [32]int16
	ChromaWeight [32]int16
	ChromaOffset [32]int16
}

// Input is an encapsulation of a set of input attributes.
type Input struct {
	Index        uint32
	Name         [32]byte
	Type         uint32
	AudioSet     uint32
	Tuner        uint32
	Standard     StdID
	Status       uint32
	Capabilities uint32
	Reserved     [3]uint32
}

// Modulator is an encapsulation of a set of modulator attributes.
type Modulator struct {
	Index      uint32
	Name       [32]byte
	Capability uint32
	RangeLow   uint32
	RangeHigh  uint32
	TXSubChans uint32
	Type       uint32
	Reserved   [3]uint32
}

type Mpeg2Picture struct {
	PictureCodingType        uint8
	FCode                    [2][2]uint8
	IntraDCPrecision         uint8
	PictureStructure         uint8
	TopFieldFirst            uint8
	FramePredFrameDCT        uint8
	ConcealmentMotionVectors uint8
	QScaleType               uint8
	IntraVLCFormat           uint8
	AlternateScan            uint8
	RepeatFirstField         uint8
	ProgressiveFrame         uint16
}

type Mpeg2Sequence struct {
	HorizontalSize            uint16
	VerticalSize              uint16
	VBVBufferSize             uint32
	ProfileAndLevelIndication uint16
	ProgressiveSequence       uint8
	ChromaFormat              uint8
}

// Output is an encapsulation of a set of output attributes.
type Output struct {
	Index        uint32
	Name         [32]byte
	Type         uint32
	AudioSet     uint32
	Modulator    uint32
	Standard     StdID
	Capabilities uint32
	Reserved     [3]uint32
}

// PixFormat is an encapsulation of a single-planar format.
type PixFormat struct {
	Width        uint32
	Height       uint32
	PixelFormat  uint32
	Field        uint32
	BytesPerLine uint32
	SizeImage    uint32
	ColorSpace   uint32
	Priv         uint32
	Flags        uint32
	Enc          uint32 // Anonymous union of YCbCr and HSV
	Quantization uint32
	XferFunc     uint32
}

// PixFormatMPlane is an encapsulation of a multi-planar format.
type PixFormatMPlane struct {
	Width        uint32
	Height       uint32
	PixelFormat  uint32
	Field        uint32
	ColorSpace   uint32
	PlaneFmt     [8]PlanePixFormat
	NumPlanes    uint8
	Flags        uint8
	Enc          uint8 // Anonymous union of YCbCr and HSV
	Quantization uint8
	XferFunc     uint8
	Reserved     [7]uint8
}

// Plane is an encapsulation of a single plane.
type Plane struct {
	BytesUsed  uint32
	Length     uint32
	M          uint32 // Union
	DataOffset uint32
	Reserved   [11]uint32
}

// PlanePixFormat is an encapsulation of a plane format.
type PlanePixFormat struct {
	SizeImage    uint32
	BytesPerLine uint32
	Reserved     [6]uint16
}

// QueryCtrl is an encapsulation of a control.
type QueryCtrl struct {
	ID           uint32
	Type         uint32
	Name         [32]byte
	Minimum      int32
	Maximum      int32
	Step         int32
	DefaultValue int32
	Flags        uint32
	Rserved      [2]uint32
}

// QueryExtCtrl is an encapsulation of an extended control.
type QueryExtCtrl struct {
	ID           uint32
	Type         uint32
	Name         [32]byte
	Minimum      int32
	Maximum      int32
	Step         int32
	DefaultValue int32
	Flags        uint32
	ElemSize     uint32
	Elems        uint32
	NrOfDims     uint32
	Dims         [4]uint32
	Rserved      [32]uint32
}

// QueryMenu is an encapsulation of a menu.
type QueryMenu struct {
	ID    uint32
	Index uint32
	Name  [32]byte
	// Value int64 unioned with Name
	Reserved uint32
}

// RequestBuffers is an encapsulates a buffer request.
type RequestBuffers struct {
	Count        uint32
	Type         uint32
	Memory       uint32
	Capabilities uint32
	Reserved     [1]uint32
}

// Timecode is an encapsulation of a timecode.
type Timecode struct {
	Type     uint32
	Flags    uint32
	Frames   uint8
	Seconds  uint8
	Minutes  uint8
	Hours    uint8
	UserBits [4]uint8
}

// Tuner is an encapsulation of a set of tuner attributes.
type Tuner struct {
	Index      uint32
	Name       [32]byte
	Type       uint32
	Capability uint32
	RangeLow   uint32
	RangeHigh  uint32
	RXSubChans uint32
	AudMode    uint32
	Signal     uint32
	AFC        int32
	Reserved   [4]uint32
}

type VP8EntropyCoderState struct {
	Range    uint8
	Value    uint8
	BitCount uint8
	Padding  uint8
}

type VP8EntropyHeader struct {
	CoeffProbs  [4][8][3][11]uint8
	YModeProbs  [4]uint8
	UVModeProbs [4]uint8
	MVProb      [2][19]uint8
	Padding     [3]uint8
}

type VP8LoopfilterHeader struct {
	RefFrmDelta    [4]int8
	MBModeDelta    [4]int8
	SharpnessLevel uint8
	Padding        uint16
	Flags          uint32
}

type VP8QuantizationHeader struct {
	YACQi     uint8
	YDCDelta  int8
	Y2DCDelta int8
	Y2ACDelta int8
	UVDCDelta int8
	UVACDelta int8
	Padding   uint16
}

type VP8SegmentHeader struct {
	QuantUpdate  [4]int8
	LFUpdate     [4]int8
	SegmentProbs [3]uint8
	Padding      uint8
	Flags        uint32
}

// BytesToString converts a low-level, null-terminated C-string to a string.
func BytesToString(b []byte) string {
	var n int
	for n = 0; n < len(b); n++ {
		if b[n] == 0 {
			break
		}
	}
	return string(b[:n+1])
}

// Ioctl performs a low-level ioctl operation.
func Ioctl(fd int, op uint32, arg uintptr) error {
	_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(op), arg)
	if err == 0 {
		return nil
	}
	return err
}
