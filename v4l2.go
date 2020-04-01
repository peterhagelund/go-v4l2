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
	// VidIocQueryCap is VIDIOC_QUERYCAP (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-querycap.html).
	VidIocQueryCap uint32 = 0x80685600
	// VidIocReserved is VIDIOC_RESERVED.
	VidIocReserved = 0x00005601
	// VidIocEnumFmt is (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-enum-fmt.html).
	VidIocEnumFmt = 0xc0405602
	// VidIocGFmt is VIDIOC_G_FMT (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-fmt.html).
	VidIocGFmt = 0xc0d05604
	// VidIocSFmt is VIDIOC_S_FMT (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-fmt.html).
	VidIocSFmt = 0xc0cc5605
	// VidIocReqBufs is VIDIOC_REQBUFS (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-reqbufs.html).
	VidIocReqBufs = 0xc0145608
	// VidIocQueryBuf is VIDIOC_QUERYBUF (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-querybuf.html).
	VidIocQueryBuf = 0xc0445609
	// VidIocGFBuf is VIDIOC_G_FBUF (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-fbuf.html).
	VidIocGFBuf = 0x8030560a
	// VidIocSFBuf is VIDIOC_S_FBUF (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-fbuf.html).
	VidIocSFBuf = 0x4030560b
	// VidIocOverlay is VIDIOC_OVERLAY (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-overlay.html).
	VidIocOverlay = 0x4004560e
	// VidIocQBuf is VIDIOC_QBUF (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-qbuf.html).
	VidIocQBuf = 0xc044560f
	// VidIocExpBuf is VIDIOC_EXPBUF (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-expbuf.html).
	VidIocExpBuf = 0xc0405610
	// VidIocDQBuf is VIDIOC_DQBUF (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-qbuf.html).
	VidIocDQBuf = 0xc0445611
	// VidIocStreamOn is VIDIOC_STREAMON (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-streamon.html).
	VidIocStreamOn = 0x40045612
	// VidIocStreamOff is VIDIOC_STREAMOFF (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-streamon.html).
	VidIocStreamOff = 0x40045613
	// VidIocGParm is VIDIOC_G_PARM (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-parm.html).
	VidIocGParm = 0xc0cc5615
	// VidIocSParm is VIDIOC_S_PARM (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-parm.html).
	VidIocSParm = 0xc0cc5616
	// VidIocGStd is VIDIOC_G_STD (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-std.html).
	VidIocGStd = 0x80085617
	// VidIocSStd is VIDIOC_S_STD (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-std.html).
	VidIocSStd = 0x40085618
	// VidIocEnumStd is VIDIOC_ENUMSTD (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-enumstd.html).
	VidIocEnumStd = 0xc0485619
	// VidIocEnumInput is VIDIOC_ENUMINPUT (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-enuminput.html).
	VidIocEnumInput = 0xc050561a
	// VidIocGCtrl is VIDIOC_G_CTRL (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-ctrl.html).
	VidIocGCtrl = 0xc008561b
	// VidIocSCtrl is VIDIOC_S_CTRL (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-ctrl.html).
	VidIocSCtrl = 0xc008561c
	// VidIocGTuner is VIDIOC_G_TUNER (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-tuner.html).
	VidIocGTuner = 0xc054561d
	// VidIocSTuner is VIDIOC_S_TUNER (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-tuner.html).
	VidIocSTuner = 0x4054561e
	// VidIocGAudio is VIDIOC_G_AUDIO (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-audio.html).
	VidIocGAudio = 0x80345621
	// VidIocSAudio is VIDIOC_S_AUDIO (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-audio.html).
	VidIocSAudio = 0x40345622
	// VidIocQueryCtrl is VIDIOC_QUERYCTRL (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-queryctrl.html).
	VidIocQueryCtrl = 0xc0445624
	// VidIocQueryMenu is VIDIOC_QUERYMENU (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-queryctrl.html)
	VidIocQueryMenu = 0xc02c5625
	// VidIocGInput is VIDIOC_G_INPUT (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-input.html).
	VidIocGInput = 0x80045626
	// VidIocSInput is VIDIOC_S_INPUT (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-input.html).
	VidIocSInput = 0xc0045627
	// VidIocGEDID is VIDIOC_G_EDID (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-edid.html).
	VidIocGEDID = 0xc0285628
	// VidIocSEDID is VIDIOC_S_EDID (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-edid.html).
	VidIocSEDID = 0xc0285629
	// VidIocGOutput is VIDIOC_G_OUTPUT (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-output.html).
	VidIocGOutput = 0x8004562e
	// VidIocSOutput is VIDIOC_S_OUTPUT (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-output.html).
	VidIocSOutput = 0xc004562f
	// VidIocEnumOutput is VIDIOC_ENUMOUTPUT (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-enumoutput.html).
	VidIocEnumOutput = 0xc0485630
	// VidIocGAudOut is VIDIOC_G_AUDOUT (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-audioout.html).
	VidIocGAudOut = 0x80345631
	// VidIocSAudOut is VIDIOC_S_AUDOUT (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-audioout.html).
	VidIocSAudOut = 0x40345632
	// VidIocGModulator is VIDIOC_G_MODULATOR (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-modulator.html).
	VidIocGModulator = 0xc0445636
	// VidIocSModulator is VIDIOC_S_MODULATOR (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-modulator.html).
	VidIocSModulator = 0x40445637
	// VidIocGFrequency is VIDIOC_G_FREQUENCY (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-frequency.html).
	VidIocGFrequency = 0xc02c5638
	// VidIocSFrequency is VIDIOC_S_FREQUENCY (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-frequency.html).
	VidIocSFrequency = 0x402c5639
	// VidIocCropCap is VIDIOC_CROPCAP (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-cropcap.html).
	VidIocCropCap = 0xc02c563a
	// VidIocGCrop is VIDIOC_G_CROP (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-crop.html).
	VidIocGCrop = 0xc014563b
	// VidIocSCrop is VIDIOC_S_CROP (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-crop.html).
	VidIocSCrop = 0x4014563c
	// VidIocGJpegComp is VIDIOC_G_JPEGCOMP (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-jpegcomp.html).
	VidIocGJpegComp = 0x808c563d
	// VidIocSJpegComp is VIDIOC_S_JPEGCOMP (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-jpegcomp.html).
	VidIocSJpegComp = 0x408c563e
	// VidIocQueryStd is VIDIOC_QUERYSTD (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-querystd.html).
	VidIocQueryStd = 0x8008563f
	// VidIocTryFmt is VIDIOC_TRY_FMT (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-fmt.html).
	VidIocTryFmt = 0xc0d05640
	// VidIocEnumAudio is VIDIOC_ENUMAUDIO (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-enumaudio.html).
	VidIocEnumAudio = 0xc0345641
	// VidIocEnumAudOut is VIDIOC_ENUMAUDOUT (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-enumaudioout.html).
	VidIocEnumAudOut = 0xc0345642
	// VidIocGPriority is VIDIOC_G_PRIORITY (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-priority.html).
	VidIocGPriority = 0x80045643
	// VidIocSPriority is VIDIOC_S_PRIORITY (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-priority.html).
	VidIocSPriority = 0x40045644
	// VidIocGSlicedVBICap is VIDIOC_G_SLICED_VBI_CAP (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-sliced-vbi-cap.html).
	VidIocGSlicedVBICap = 0xc0745645
	// VidIocLogStatus is VIDIOC_LOG_STATUS (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-log-status.html).
	VidIocLogStatus = 0x00005646
	// VidIocGExtCtrls is VIDIOC_G_EXT_CTRLS (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-ext-ctrls.html).
	VidIocGExtCtrls = 0xc0205647
	// VidIocSExtCtrls is VIDIOC_S_EXT_CTRLS (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-ext-ctrls.html).
	VidIocSExtCtrls = 0xc0205648
	// VidIocTryExtCtrls is VIDIOC_TRY_EXT_CTRLS (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-ext-ctrls.html).
	VidIocTryExtCtrls = 0xc0205649
	// VidIocEnumFrameSizes is VIDIOC_ENUM_FRAMESIZES (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-enum-framesizes.html).
	VidIocEnumFrameSizes = 0xc02c564a
	// VidIocEnumFrameIntervals is VIDIOC_ENUM_FRAMEINTERVALS (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-enum-frameintervals.html).
	VidIocEnumFrameIntervals = 0xc034564b
	// VidIocGEncIndex is VIDIOC_G_ENC_INDEX (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-enc-index.html).
	VidIocGEncIndex = 0x8818564c
	// VidIocEncoderCmd is VIDIOC_ENCODER_CMD (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-encoder-cmd.html).
	VidIocEncoderCmd = 0xc028564d
	// VidIocTryEncoderCmd is VIDIOC_TRY_ENCODER_CMD (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-encoder-cmd.html).
	VidIocTryEncoderCmd = 0xc028564e
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
	Reserved    [4]uint32
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

// FrameSizeEnum is an encapsulation of frame size information.
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
	PixFormat    uint32
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
	PixFormat    uint32
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
	return string(b[:n])
}
