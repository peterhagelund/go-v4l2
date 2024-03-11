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

package v4l2

import (
	"bytes"
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

// AudCap is the audio capability type.
type AudCap uint32

// Audio capability flags.
const (
	AudCapStereo AudCap = 1 << iota
	AudCapAVL
)

// AudMode is the audio mode type.
type AudMode uint32

// Audio mode flags.
const (
	AudModeAVL = 1 << iota
)

// ColorSpace is the the color space type.
type ColorSpace uint32

// Color spaces.
const (
	ColorSpaceDefault ColorSpace = iota // TODO Find actual values
	ColorSpaceSMPTE170M
	ColorSpaceRec709
	ColorSpaceSRGB
	ColorSpaceOPRGB
	ColorSpaceBT2020
	ColorSpaceDCIP3
	ColorSpaceSMPTE240M
	ColorSpace470SystemM
	ColorSpace470SystemBG
	ColorSpaceJPEG
	ColorSpaceRaw
)

// Buffer capabilities.
const (
	BufCapSupportsMMap uint32 = 1 << iota
	BufCapSupportsUserPtr
	BufCapSupportsDMABuf
	BufCapSupportsRequests
	BufCapSupportsOrphanedBufs
	BufCapSupportsM2MHoldCaptureBuf
)

// BufFlag is the buffer flag type.
type BufFlag uint32

// Buffer flags.
const (
	BufFlagMapped BufFlag = 1 << iota
	BufFlagQueued
	BufFlagDone
	BufFlagKeyFrame
	BufFlagPFrame
	BufFlagBFrame
	BufFlagError
	BufFlagInRequest
	BufFlagTimecode
	BufFlagM2MHoldCaptureBuf
	BufFlagPrepared
	BufFlagNoCacheInvalidate
	BufFlagNoCacheClean
	BufFlagTimestampMonotonic
	BufFlagTimestampCopy
	_
	BufFlagTstampSOE
	BufFlagTstampSrcEOF  BufFlag = 0x00000000
	BufFlagTimestampMask BufFlag = 0x0000e000
	BufFlagTstampSrcMask BufFlag = 0x00070000
	BufFlagLast          BufFlag = 0x00100000
	BufFlagRequestFD     BufFlag = 0x00800000
)

// BufType is the buffer type type.
type BufType uint32

// Buffer types.
const (
	BufTypeVideoCapture BufType = iota + 1
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

// Cap is the device capability type.
type Cap uint32

// Device capability flags.
const (
	CapVideoCapture Cap = 1 << iota
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
	CapReadWrite
	CapAsyncIO
	CapStreaming
	CapMetaOutput
	CapTouch
	_
	_
	CapDeviceCaps
)

// Field is the field type.
type Field uint32

// Field types.
const (
	FieldAny Field = iota
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

// FmtFlag is the format flag type.
type FmtFlag uint32

// Format flags.
const (
	FmtFlagCompressed FmtFlag = 1 << iota
	FmtFlagEmulated
	FmtFlagContinuousByteStream
	FmtFlagDynResolution
)

// FrmSizeType is the frame size type type.
type FrmSizeType uint32

// Frame size types.
const (
	FrmSizeTypeDiscrete FrmSizeType = iota + 1
	FrmSizeTypeContinuous
	FrmSizeTypeStepwise
)

// InputCap is the input capabilities type.
type InputCap uint32

// Input capabilities.
const (
	_ InputCap = 1 << iota
	InputCapDVTimings
	InputCapStd
	InputCapNativeSize
)

// InputStatus is the input status type.
type InputStatus uint32

// Input statuses.
const (
	InputStatusNoPower InputStatus = 1 << iota
	InputStatusNoSignal
	InputStatusNoColor
	_
	InputStatusHFlip
	InputStatusVFlip
	_
	_
	InputStatusNoHLock
	InputStatusColorKill
	InputStatusNoVLock
	InputStatusNoStdLock
	_
	_
	_
	_
	InputStatusNoSync
	InputStatusNoEqu
	InputStatusNoCarrier
	_
	_
	_
	_
	_
	InputStatusMacroVision
	InputStatusNoAccess
	InputStatusVTR
)

// InputType is the input type type.
type InputType uint32

// Input types.
const (
	InputTypeTuner InputType = iota + 1
	InputTypeCamera
	InputTypeTouch
)

// Memory is the memory type.
type Memory uint32

// Memory types.
const (
	MemoryMmap Memory = iota + 1
	MemoryUserPtr
	MemoryOverlay
	MemoryDMABuf
)

// OutputCap is output capabilities type.
type OutputCap uint32

// Output capabilities.
const (
	_ OutputCap = 1 << iota
	OutputCapDVTimings
	OutputCapStd
	OutputCapNativeSize
)

// OutputType is output type type.
type OutputType uint32

// Output types.
const (
	OutputTypeModulator = 1 + iota
	OutputTypeAnalog
	OutputTypeAnalogVGAOverlay
)

// PixFmt is the pixel format type.
type PixFmt uint32

// Pixel formats.
const (
	PixFmtRGB332         PixFmt = 'R' | 'G'<<8 | 'B'<<16 | '1'<<24
	PixFmtARGB444        PixFmt = 'A' | 'R'<<8 | '1'<<16 | '2'<<24
	PixFmtXRGB444        PixFmt = 'X' | 'R'<<8 | '1'<<16 | '2'<<24
	PixFmtRGBA444        PixFmt = 'R' | 'A'<<8 | '1'<<16 | '2'<<24
	PixFmtRGBX444        PixFmt = 'R' | 'X'<<8 | '1'<<16 | '2'<<24
	PixFmtABGR444        PixFmt = 'A' | 'B'<<8 | '1'<<16 | '2'<<24
	PixFmtXBGR444        PixFmt = 'X' | 'B'<<8 | '1'<<16 | '2'<<24
	PixFmtBGRA444        PixFmt = 'B' | 'A'<<8 | '1'<<16 | '2'<<24
	PixFmtBGRX444        PixFmt = 'B' | 'X'<<8 | '1'<<16 | '2'<<24
	PixFmtARGB555        PixFmt = 'A' | 'R'<<8 | '1'<<16 | '5'<<24
	PixFmtXRGB555        PixFmt = 'X' | 'R'<<8 | '1'<<16 | '5'<<24
	PixFmtRGBA555        PixFmt = 'R' | 'A'<<8 | '1'<<16 | '5'<<24
	PixFmtRGBX555        PixFmt = 'R' | 'X'<<8 | '1'<<16 | '5'<<24
	PixFmtABGR555        PixFmt = 'A' | 'B'<<8 | '1'<<16 | '5'<<24
	PixFmtXBGR555        PixFmt = 'X' | 'B'<<8 | '1'<<16 | '5'<<24
	PixFmtBGRA555        PixFmt = 'B' | 'A'<<8 | '1'<<16 | '5'<<24
	PixFmtBGRX555        PixFmt = 'B' | 'X'<<8 | '1'<<16 | '5'<<24
	PixFmtRGB565         PixFmt = 'R' | 'G'<<8 | 'B'<<16 | 'P'<<24
	PixFmtARGB555X       PixFmt = PixFmtARGB555 | 1<<31
	PixFmtXRGB555X       PixFmt = PixFmtXRGB555 | 1<<31
	PixFmtRGB565X        PixFmt = 'R' | 'G'<<8 | 'B'<<16 | 'R'<<24
	PixFmtBGR24          PixFmt = 'B' | 'G'<<8 | 'R'<<16 | '3'<<24
	PixFmtRGB24          PixFmt = 'R' | 'G'<<8 | 'B'<<16 | '3'<<24
	PixFmtBGR666         PixFmt = 'B' | 'G'<<8 | 'R'<<16 | 'H'<<24 // Format from hell?
	PixFmtABGR32         PixFmt = 'A' | 'R'<<8 | '2'<<16 | '4'<<24
	PixFmtXBGR32         PixFmt = 'X' | 'R'<<8 | '2'<<16 | '4'<<24
	PixFmtBGRA32         PixFmt = 'R' | 'A'<<8 | '2'<<16 | '4'<<24
	PixFmtBGRX32         PixFmt = 'R' | 'X'<<8 | '2'<<16 | '4'<<24
	PixFmtRGBA32         PixFmt = 'A' | 'B'<<8 | '2'<<16 | '4'<<24
	PixFmtRGBX32         PixFmt = 'X' | 'B'<<8 | '2'<<16 | '4'<<24
	PixFmtARGB32         PixFmt = 'B' | 'A'<<8 | '2'<<16 | '4'<<24
	PixFmtXRGB32         PixFmt = 'B' | 'X'<<8 | '2'<<16 | '4'<<24
	PixFmtGrey           PixFmt = 'G' | 'R'<<8 | 'E'<<16 | 'Y'<<24
	PixFmtY10            PixFmt = 'Y' | '1'<<8 | '0'<<16 | ' '<<24
	PixFmtY12            PixFmt = 'Y' | '1'<<8 | '2'<<16 | ' '<<24
	PixFmtY10BPack       PixFmt = 'Y' | '1'<<8 | '0'<<16 | 'B'<<24
	PixFmtY10P           PixFmt = 'Y' | '1'<<8 | '0'<<16 | 'P'<<24
	PixFmtY16            PixFmt = 'Y' | '1'<<8 | '6'<<16 | ' '<<24
	PixFmtY16BE          PixFmt = PixFmtY16 | 1<<31
	PixFmtY8I            PixFmt = 'Y' | '8'<<8 | 'I'<<16 | ' '<<24
	PixFmtY12I           PixFmt = 'Y' | '1'<<8 | '2'<<16 | 'I'<<24
	PixFmtUV8            PixFmt = 'U' | 'V'<<8 | '8'<<16 | ' '<<24
	PixFmtYUYV           PixFmt = 'Y' | 'U'<<8 | 'Y'<<16 | 'V'<<24
	PixFmtUYVY           PixFmt = 'U' | 'Y'<<8 | 'V'<<16 | 'Y'<<24
	PixFmtYVYU           PixFmt = 'Y' | 'V'<<8 | 'Y'<<16 | 'U'<<24
	PixFmtVYUY           PixFmt = 'V' | 'Y'<<8 | 'U'<<16 | 'Y'<<24
	PixFmtY41P           PixFmt = 'Y' | '4'<<8 | '1'<<16 | 'P'<<24
	PixFmtYVU420         PixFmt = 'Y' | 'V'<<8 | '1'<<16 | '2'<<24
	PixFmtYUV420         PixFmt = 'Y' | 'U'<<8 | '1'<<16 | '2'<<24
	PixFmtYUV420M        PixFmt = 'Y' | 'M'<<8 | '1'<<16 | '2'<<24
	PixFmtYVU420M        PixFmt = 'Y' | 'M'<<8 | '2'<<16 | '1'<<24
	PixFmtYUV422M        PixFmt = 'Y' | 'M'<<8 | '1'<<16 | '6'<<24
	PixFmtYVU422M        PixFmt = 'Y' | 'M'<<8 | '6'<<16 | '1'<<24
	PixFmtYUV444M        PixFmt = 'Y' | 'M'<<8 | '2'<<16 | '4'<<24
	PixFmtYVU444M        PixFmt = 'Y' | 'M'<<8 | '4'<<16 | '2'<<24
	PixFmtYVU410         PixFmt = 'Y' | 'V'<<8 | 'U'<<16 | '9'<<24
	PixFmtYUV410         PixFmt = 'Y' | 'U'<<8 | 'V'<<16 | '9'<<24
	PixFmtYUV422P        PixFmt = '4' | '2'<<8 | '2'<<16 | 'P'<<24
	PixFmtYUV411P        PixFmt = '4' | '1'<<8 | '1'<<16 | 'P'<<24
	PixFmtNV12           PixFmt = 'N' | 'V'<<8 | '1'<<16 | '2'<<24
	PixFmtNV21           PixFmt = 'N' | 'V'<<8 | '2'<<16 | '1'<<24
	PixFmtNV12M          PixFmt = 'N' | 'M'<<8 | '1'<<16 | '2'<<24
	PixFmtNV21M          PixFmt = 'N' | 'M'<<8 | '2'<<16 | '1'<<24
	PixFmtNV12MT         PixFmt = 'T' | 'M'<<8 | '1'<<16 | '2'<<24
	PixFmtNV16           PixFmt = 'N' | 'V'<<8 | '1'<<16 | '6'<<24
	PixFmtNV61           PixFmt = 'N' | 'V'<<8 | '6'<<16 | '1'<<24
	PixFmtNV16M          PixFmt = 'N' | 'M'<<8 | '1'<<16 | '6'<<24
	PixFmtNV61M          PixFmt = 'N' | 'M'<<8 | '6'<<16 | '1'<<24
	PixFmtNV24           PixFmt = 'N' | 'V'<<8 | '2'<<16 | '4'<<24
	PixFmtNV42           PixFmt = 'N' | 'V'<<8 | '4'<<16 | '2'<<24
	PixFmtM420           PixFmt = 'M' | '4'<<8 | '2'<<16 | '0'<<24
	PixFmtDV             PixFmt = 'd' | 'v'<<8 | 's'<<16 | 'd'<<24
	PixFmtET61X251       PixFmt = 'E' | '6'<<8 | '2'<<16 | '5'<<24
	PixFmtHI240          PixFmt = 'H' | 'I'<<8 | '2'<<16 | '4'<<24
	PixFmtHM12           PixFmt = 'H' | 'M'<<8 | '1'<<16 | '2'<<24
	PixFmtCPIA1          PixFmt = 'C' | 'P'<<8 | 'I'<<16 | 'A'<<24
	PixFmtJPGL           PixFmt = 'J' | 'P'<<8 | 'G'<<16 | 'L'<<24
	PixFmtSPCA501        PixFmt = 'S' | '5'<<8 | '0'<<16 | '1'<<24
	PixFmtSPCA505        PixFmt = 'S' | '5'<<8 | '0'<<16 | '5'<<24
	PixFmtSPCA508        PixFmt = 'S' | '5'<<8 | '0'<<16 | '8'<<24
	PixFmtSPCA561        PixFmt = 'S' | '5'<<8 | '6'<<16 | '1'<<24
	PixFmtPAC207         PixFmt = 'P' | '2'<<8 | '0'<<16 | '7'<<24
	PixFmtMR97310A       PixFmt = 'M' | '3'<<8 | '1'<<16 | '0'<<24
	PixFmtJL2005BCD      PixFmt = 'J' | 'L'<<8 | '2'<<16 | '0'<<24
	PixFmtOV511          PixFmt = 'O' | '5'<<8 | '1'<<16 | '1'<<24
	PixFmtOV518          PixFmt = 'O' | '5'<<8 | '1'<<16 | '8'<<24
	PixFmtPJPG           PixFmt = 'P' | 'J'<<8 | 'P'<<16 | 'G'<<24
	PixFmtSE401          PixFmt = 'S' | '4'<<8 | '0'<<16 | '1'<<24
	PixFmtSQ90C          PixFmt = '9' | '0'<<8 | '5'<<16 | 'C'<<24
	PixFmtMJPEG          PixFmt = 'M' | 'J'<<8 | 'P'<<16 | 'G'<<24
	PixFmtPWC1           PixFmt = 'P' | 'W'<<8 | 'C'<<16 | '1'<<24
	PixFmtPWC2           PixFmt = 'P' | 'W'<<8 | 'C'<<16 | '2'<<24
	PixFmtSN9C10X        PixFmt = 'S' | '9'<<8 | '1'<<16 | '0'<<24
	PixFmtN9C20XI420     PixFmt = 'S' | '9'<<8 | '2'<<16 | '0'<<24
	PixFmtSN9C2028       PixFmt = 'S' | 'O'<<8 | 'N'<<16 | 'X'<<24
	PixFmtSTV0680        PixFmt = 'S' | '6'<<8 | '8'<<16 | '0'<<24
	PixFmtWNVA           PixFmt = 'W' | 'N'<<8 | 'V'<<16 | 'A'<<24
	PixFmtTM6000         PixFmt = 'T' | 'M'<<8 | '6'<<16 | '0'<<24
	PixFmtCITYYVYUY      PixFmt = 'C' | 'I'<<8 | 'T'<<16 | 'V'<<24
	PixFmtKonica420      PixFmt = 'K' | 'O'<<8 | 'N'<<16 | 'I'<<24
	PixFmtYYUV           PixFmt = 'Y' | 'Y'<<8 | 'U'<<16 | 'V'<<24
	PixFmtY4             PixFmt = 'Y' | '0'<<8 | '4'<<16 | ' '<<24
	PixFmtY6             PixFmt = 'Y' | '0'<<8 | '6'<<16 | ' '<<24
	PixFmtS5CUYVYJPG     PixFmt = 'S' | '5'<<8 | 'C'<<16 | 'I'<<24
	PixFmtMT21C          PixFmt = 'M' | 'T'<<8 | '2'<<16 | '1'<<24
	PixFmtSunXITiledNV12 PixFmt = 'S' | 'T'<<8 | '1'<<16 | '2'<<24
	PixFmtPAL8           PixFmt = 'P' | 'A'<<8 | 'L'<<16 | '8'<<24
	PixFmtJPEG           PixFmt = 'J' | 'P'<<8 | 'E'<<16 | 'G'<<24
	PixFmtMPEG           PixFmt = 'M' | 'P'<<8 | 'E'<<16 | 'G'<<24
	PixFmtH264           PixFmt = 'H' | '2'<<8 | '6'<<16 | '4'<<24
	PixFmtH264NoSC       PixFmt = 'A' | 'V'<<8 | 'C'<<16 | '1'<<24
	PixFmtH264MVC        PixFmt = 'M' | '2'<<8 | '6'<<16 | '4'<<24
	PixFmtH264Slice      PixFmt = 'S' | '2'<<8 | '6'<<16 | '4'<<24
	PixFmtH263           PixFmt = 'H' | '2'<<8 | '6'<<16 | '3'<<24
	PixFmtMPEG1          PixFmt = 'M' | 'P'<<8 | 'G'<<16 | '1'<<24
	PixFmtMPEG2          PixFmt = 'M' | 'P'<<8 | 'G'<<16 | '2'<<24
	PixFmtMPEG2Slice     PixFmt = 'M' | 'G'<<8 | '2'<<16 | 'S'<<24
	PixFmtMPEG4          PixFmt = 'M' | 'P'<<8 | 'G'<<16 | '4'<<24
	PixFmtXVID           PixFmt = 'X' | 'V'<<8 | 'I'<<16 | 'D'<<24
	PixFmtVC1AnnexG      PixFmt = 'V' | 'C'<<8 | '1'<<16 | 'G'<<24
	PixFmtVC1AnnexL      PixFmt = 'V' | 'C'<<8 | '1'<<16 | 'L'<<24
	PixFmtVP8            PixFmt = 'V' | 'P'<<8 | '8'<<16 | '0'<<24
	PixFmtVP8Frame       PixFmt = 'V' | 'P'<<8 | '8'<<16 | 'F'<<24
	PixFmtVP9            PixFmt = 'V' | 'P'<<8 | '9'<<16 | '0'<<24
	PixFmtHEVC           PixFmt = 'H' | 'E'<<8 | 'V'<<16 | 'C'<<24
	PixFmtHEVCSlice      PixFmt = 'S' | '2'<<8 | '6'<<16 | '5'<<24
	PixFmtFWHT           PixFmt = 'F' | 'W'<<8 | 'H'<<16 | 'T'<<24
	PixFmtFWHTStateless  PixFmt = 'S' | 'F'<<8 | 'W'<<16 | 'H'<<24
	// TODO There are more...
)

// PixFmtFlag is the pixel format flag type.
type PixFmtFlag uint32

// Pixel format flags.
const (
	PixFmtFlagPremulAlpha = 1 << iota
)

// Quantization is the quantization type.
type Quantization uint32

// The quantization flags.
const (
	QuantizationDefault Quantization = iota
	QuantizationFullRange
	QuantizationLimRange
)

// SlicedVBIService is the sliced VBI service type.
type SlicedVBIService uint16

// The sliced VBI service.
const (
	SlicedVBIServiceTeleTextB SlicedVBIService = 0x0001
)

// TcFlag is the timecode flag type.
type TcFlag uint32

// Timecode flags.
const (
	TcFlagUserDefined TcFlag = 1 << iota
	TcFlagDropFrame
	_
	TcUserBits8BitChars
	TcUserBitsField       TcFlag = 0x0000000C
	TcUserBitsUserDefined TcFlag = 0x00000000
)

// TcType is the timecode type type.
type TcType uint32

// Timecode types.
const (
	TcType24Fps TcType = iota + 1
	TcType25Fps
	TcType30Fps
	TcType50Fps
	TcType60Fps
)

// CtrlType is the control type type.
type CtrlType uint32

// The control types.
const (
	CtrlTypeInteger               CtrlType = 1
	CtrlTypeBoolean               CtrlType = 2
	CtrlTypeMenu                  CtrlType = 3
	CtrlTypeButton                CtrlType = 4
	CtrlTypeInteger64             CtrlType = 5
	CtrlTypeCtrlClass             CtrlType = 6
	CtrlTypeString                CtrlType = 7
	CtrlTypeBitMask               CtrlType = 8
	CtrlTypeIntegerMenu           CtrlType = 9
	CtrlCompundTypes              CtrlType = 0x0100
	CtrlTypeU8                    CtrlType = 0x0100
	CtrlTypeU16                   CtrlType = 0x0101
	CtrlTypeU32                   CtrlType = 0x0102
	CtrlTypeArea                  CtrlType = 0x0106
	CtrlTypeHdr10CllInfo          CtrlType = 0x0110
	CtrlTypeHdr10MasteringDisplay CtrlType = 0x0111
	CtrlTypeH264SPS               CtrlType = 0x0200
	CtrlTypeH264PPS               CtrlType = 0x0201
	CtrlTypeH264ScalingMatrix     CtrlType = 0x0202
	CtrlTypeH264SliceParams       CtrlType = 0x0203
	CtrlTypeH264DecordeParams     CtrlType = 0x0204
	CtrlTypeH264PredWeights       CtrlType = 0x0205
	CtrlTypeFwhtParams            CtrlType = 0x0220
	CtrlTypeVP8Frame              CtrlType = 0x0240
	CtrlTypeMPEG2Quatisation      CtrlType = 0x0250
	CtrlTypeMPEG2Sequence         CtrlType = 0x0251
	CtrlTypeMPEG2Picture          CtrlType = 0x0252
)

// CtrlType is the control flag type.
type CtrlFlag uint32

// The control flags.
const (
	CtrlFlagDisable        CtrlFlag = 0x0001
	CtrlFlagGrabbed        CtrlFlag = 0x0002
	CtrlFlagReadOnly       CtrlFlag = 0x0004
	CtrlFlagUpdate         CtrlFlag = 0x0008
	CtrlFlagInactive       CtrlFlag = 0x0010
	CtrlFlagSlider         CtrlFlag = 0x0020
	CtrlFlagWriteOnly      CtrlFlag = 0x0040
	CtrlFlagVolatile       CtrlFlag = 0x0080
	CtrlFlagHasPayload     CtrlFlag = 0x0100
	CtrlFlagExecuteOnWrite CtrlFlag = 0x0200
	CtrlFlagModifyLayout   CtrlFlag = 0x0400
	CtrlFlagNextCtrl       CtrlFlag = 0x80000000
	CtrlFlagNextCompound   CtrlFlag = 0x40000000
	CidMaxCtrls            CtrlFlag = 1024
	CidPrivateBase         CtrlFlag = 0x08000000
)

// CtrlClass is the control class type.
type CtrlClass uint32

// The control classes.
const (
	CtrlClassUser        CtrlClass = 0x00980000
	CtrlClassCodec       CtrlClass = 0x00990000
	CtrlClassCamera      CtrlClass = 0x009a0000
	CtrlClassFMTX        CtrlClass = 0x009b0000
	CtrlClassFlash       CtrlClass = 0x009c0000
	CtrlClassJPEG        CtrlClass = 0x009d0000
	CtrlClassImageSource CtrlClass = 0x009e0000
	CtrlClassImageProc   CtrlClass = 0x009f0000
	CtrlClassDV          CtrlClass = 0x00a00000
	CtrlClassFMRX        CtrlClass = 0x00a10000
	CtrlClassRFTuner     CtrlClass = 0x00a20000
	CtrlClassDetect      CtrlClass = 0x00a30000
	CtrlClassColorimetry CtrlClass = 0x00a50000
)

// CtrlID is the control ID type.
type CtrlID uint32

// The control IDs.
const (
	CidBase                                           CtrlID = CtrlID(CtrlClassUser | 0x900)
	CidUserBase                                       CtrlID = CidBase
	CidUserClass                                      CtrlID = CtrlID(CtrlClassUser | 1)
	CidBrightness                                     CtrlID = CidBase + 0
	CidContrast                                       CtrlID = CidBase + 1
	CidSaturation                                     CtrlID = CidBase + 2
	CidHue                                            CtrlID = CidBase + 3
	CidAudioVolume                                    CtrlID = CidBase + 5
	CidAudioBalance                                   CtrlID = CidBase + 6
	CidAudioBass                                      CtrlID = CidBase + 7
	CidAudioTreble                                    CtrlID = CidBase + 8
	CidAudioMute                                      CtrlID = CidBase + 9
	CidAudioLoudness                                  CtrlID = CidBase + 10
	CidBlackLevel                                     CtrlID = CidBase + 11
	CidAutoWhiteBalance                               CtrlID = CidBase + 12
	CidDoWhiteBalance                                 CtrlID = CidBase + 13
	CidRedBalance                                     CtrlID = CidBase + 14
	CidBlueBalance                                    CtrlID = CidBase + 15
	CidGamma                                          CtrlID = CidBase + 16
	CidWhiteness                                      CtrlID = CidGamma
	CidExposure                                       CtrlID = CidBase + 17
	CidAutogain                                       CtrlID = CidBase + 18
	CidGain                                           CtrlID = CidBase + 19
	CidHFlip                                          CtrlID = CidBase + 20
	CidVFlip                                          CtrlID = CidBase + 21
	CidPowerLineFrequency                             CtrlID = CidBase + 24
	CidHueAuto                                        CtrlID = CidBase + 25
	CidWhiteBalanceTemperature                        CtrlID = CidBase + 26
	CidSharpness                                      CtrlID = CidBase + 27
	CidBacklightCompensation                          CtrlID = CidBase + 28
	CidChromaAgc                                      CtrlID = CidBase + 29
	CidColorKiller                                    CtrlID = CidBase + 30
	CidColorFX                                        CtrlID = CidBase + 31
	CidAutobrightness                                 CtrlID = CidBase + 32
	CidBandStopFilter                                 CtrlID = CidBase + 33
	CidRotate                                         CtrlID = CidBase + 34
	CidBgColor                                        CtrlID = CidBase + 35
	CidChromaGain                                     CtrlID = CidBase + 36
	CidIlluminators1                                  CtrlID = CidBase + 37
	CidIlluminators2                                  CtrlID = CidBase + 38
	CidMinBuffersForCapture                           CtrlID = CidBase + 39
	CidMinBuffersForOutput                            CtrlID = CidBase + 40
	CidAlphaComponent                                 CtrlID = CidBase + 41
	CidColorrFXCBCR                                   CtrlID = CidBase + 42
	CidLastP1                                         CtrlID = CidBase + 43
	CidUserMEYEBase                                   CtrlID = CidUserBase + 0x1000
	CidUserBTTVBase                                   CtrlID = CidUserBase + 0x1010
	CidUserS2255Base                                  CtrlID = CidUserBase + 0x1030
	CidUserSI476XBase                                 CtrlID = CidUserBase + 0x1040
	CidUserTIVPEBase                                  CtrlID = CidUserBase + 0x1050
	CidUserSAA7134Base                                CtrlID = CidUserBase + 0x1060
	CidUserADV7180Base                                CtrlID = CidUserBase + 0x1070
	CidUserTC358743Base                               CtrlID = CidUserBase + 0x1080
	CidUserMAX217XBase                                CtrlID = CidUserBase + 0x1090
	CidUserIMXBase                                    CtrlID = CidUserBase + 0x10b0
	CidUserAtmelISCBase                               CtrlID = CidUserBase + 0x10c0
	CidUserCODABase                                   CtrlID = CidUserBase + 0x10e0
	CidUserCCSBase                                    CtrlID = CidUserBase + 0x10f0
	CidCodecBase                                      CtrlID = CtrlID(CtrlClassCodec | 0x900)
	CidCodecClass                                     CtrlID = CtrlID(CtrlClassCodec | 1)
	CidMPEGStreamType                                 CtrlID = CidCodecBase + 0
	CidMPEGStreamPidPmt                               CtrlID = CidCodecBase + 1
	CidMPEGStreamPidAudio                             CtrlID = CidCodecBase + 2
	CidMPEGStreamPidVideo                             CtrlID = CidCodecBase + 3
	CidMPEGStreamPidPCR                               CtrlID = CidCodecBase + 4
	CidMPEGStreamPesIDAudio                           CtrlID = CidCodecBase + 5
	CidMPEGStreamPesIDVideo                           CtrlID = CidCodecBase + 6
	CidMPEGStreamVBIFmt                               CtrlID = CidCodecBase + 7
	CidMPEGAudioSamplingFreq                          CtrlID = CidCodecBase + 100
	CidMPEGAudioEnCoding                              CtrlID = CidCodecBase + 101
	CidMPEGAudioL1Bitrate                             CtrlID = CidCodecBase + 102
	CidMPEGAudioL2Bitrate                             CtrlID = CidCodecBase + 103
	CidMPEGAudioL3Bitrate                             CtrlID = CidCodecBase + 104
	CidMPEGAudioMode                                  CtrlID = CidCodecBase + 105
	CidMPEGAudioModeExtension                         CtrlID = CidCodecBase + 106
	CidMPEGAudioEmphasis                              CtrlID = CidCodecBase + 107
	CidMPEGAudioCRC                                   CtrlID = CidCodecBase + 108
	CidMPEGAudioMute                                  CtrlID = CidCodecBase + 109
	CidMPEGAudioAACBitrate                            CtrlID = CidCodecBase + 110
	CidMPEGAudioAC3Bitrate                            CtrlID = CidCodecBase + 111
	CidMPEGAudioDecPlayback                           CtrlID = CidCodecBase + 112
	CidMPEGAudioDecMultilingualPlayback               CtrlID = CidCodecBase + 113
	CidMPEGVideoEnCoding                              CtrlID = CidCodecBase + 200
	CidMPEGVideoAspect                                CtrlID = CidCodecBase + 201
	CidMPEGVideoBFrames                               CtrlID = CidCodecBase + 202
	CidMPEGVideoGOPSize                               CtrlID = CidCodecBase + 203
	CidMPEGVideoGOPClosure                            CtrlID = CidCodecBase + 204
	CidMPEGVideoPulldown                              CtrlID = CidCodecBase + 205
	CidMPEGVideoBitrateMode                           CtrlID = CidCodecBase + 206
	CidMPEGVideoBitrate                               CtrlID = CidCodecBase + 207
	CidMPEGVideoBitratePeak                           CtrlID = CidCodecBase + 208
	CidMPEGVideoTemporalDecimation                    CtrlID = CidCodecBase + 209
	CidMPEGVideoMute                                  CtrlID = CidCodecBase + 210
	CidMPEGVideoMuteYUV                               CtrlID = CidCodecBase + 211
	CidMPEGVideoDecoderSliceInterface                 CtrlID = CidCodecBase + 212
	CidMPEGVideoDecoderMPEG4DeblockFilter             CtrlID = CidCodecBase + 213
	CidMPEGVideoCyclicIntraRefreshMB                  CtrlID = CidCodecBase + 214
	CidMPEGVideoFeameRCEnable                         CtrlID = CidCodecBase + 215
	CidMPEGVideoHeaderMode                            CtrlID = CidCodecBase + 216
	CidMPEGVideoMaxRefPic                             CtrlID = CidCodecBase + 217
	CidMPEGVideoMBRCEnable                            CtrlID = CidCodecBase + 218
	CidMPEGVideoMultiSliceMaxBytes                    CtrlID = CidCodecBase + 219
	CidMPEGVideoMultiSliceMaxMB                       CtrlID = CidCodecBase + 220
	CidMPEGVideoMultiSliceMode                        CtrlID = CidCodecBase + 221
	CidMPEGVideoVBVSize                               CtrlID = CidCodecBase + 222
	CidMPEGVideoDecPTS                                CtrlID = CidCodecBase + 223
	CidMPEGVideoDecFrame                              CtrlID = CidCodecBase + 224
	CidMPEGVideoVBVDelay                              CtrlID = CidCodecBase + 225
	CidMPEGVideoREPEAT_SEQ_HEADER                     CtrlID = CidCodecBase + 226
	CidMPEGVideoMV_HSearchRange                       CtrlID = CidCodecBase + 227
	CidMPEGVideoMV_VSearchRange                       CtrlID = CidCodecBase + 228
	CidMPEGVideoForceKeyFrame                         CtrlID = CidCodecBase + 229
	CidMPEGVideoBaselayerPriorityID                   CtrlID = CidCodecBase + 230
	CidMPEGVideoAUDelimiter                           CtrlID = CidCodecBase + 231
	CidMPEGVideoLTRCount                              CtrlID = CidCodecBase + 232
	CidMPEGVideoFrameLTRIndex                         CtrlID = CidCodecBase + 233
	CidMPEGVideoUseLTRFrameS                          CtrlID = CidCodecBase + 234
	CidMPEGVideoDecConcealColor                       CtrlID = CidCodecBase + 235
	CidMPEGVideoIntraRefreshPeriod                    CtrlID = CidCodecBase + 236
	CidMPEGVideoMPEG2Level                            CtrlID = CidCodecBase + 270
	CidMPEGVideoMPEG2Profile                          CtrlID = CidCodecBase + 271
	CidFwhtIFrameQP                                   CtrlID = CidCodecBase + 290
	CidFwhtPFrameQP                                   CtrlID = CidCodecBase + 291
	CidMPEGVideoH263IFrameQP                          CtrlID = CidCodecBase + 300
	CidMPEGVideoH263PFrameQP                          CtrlID = CidCodecBase + 301
	CidMPEGVideoH263BFrameQP                          CtrlID = CidCodecBase + 302
	CidMPEGVideoH263MinQP                             CtrlID = CidCodecBase + 303
	CidMPEGVideoH263MaxQP                             CtrlID = CidCodecBase + 304
	CidMPEGVideoH264IFrameQP                          CtrlID = CidCodecBase + 350
	CidMPEGVideoH264PFrameQP                          CtrlID = CidCodecBase + 351
	CidMPEGVideoH264BFrameQP                          CtrlID = CidCodecBase + 352
	CidMPEGVideoH264MinQP                             CtrlID = CidCodecBase + 353
	CidMPEGVideoH264MaxQP                             CtrlID = CidCodecBase + 354
	CidMPEGVideoH2648X8Transform                      CtrlID = CidCodecBase + 355
	CidMPEGVideoH264CPBSize                           CtrlID = CidCodecBase + 356
	CidMPEGVideoH264EntropyMode                       CtrlID = CidCodecBase + 357
	CidMPEGVideoH264IPeriod                           CtrlID = CidCodecBase + 358
	CidMPEGVideoH264Level                             CtrlID = CidCodecBase + 359
	CidMPEGVideoH264LoopFilterAlpha                   CtrlID = CidCodecBase + 360
	CdMPEGVideoH264LoopFilterBeta                     CtrlID = CidCodecBase + 361
	CdMPEGVideoH264LoopFilterMode                     CtrlID = CidCodecBase + 362
	CidMPEGVideoH264Profile                           CtrlID = CidCodecBase + 363
	CidMPEGVideoH264VUI_EXT_SAR_HEIGHT                CtrlID = CidCodecBase + 364
	CidMPEGVideoH264VUI_EXT_SAR_WIDTH                 CtrlID = CidCodecBase + 365
	CidMPEGVideoH264VUI_SAR_ENABLE                    CtrlID = CidCodecBase + 366
	CidMPEGVideoH264VUI_SAR_IDC                       CtrlID = CidCodecBase + 367
	CidMPEGVideoH264SEIFramePACKING                   CtrlID = CidCodecBase + 368
	CidMPEGVideoH264SEI_FPCurrentFrame0               CtrlID = CidCodecBase + 369
	CidMPEGVideoH264SEI_FPArrangementType             CtrlID = CidCodecBase + 370
	CidMPEGVideoH264FMO                               CtrlID = CidCodecBase + 371
	CidMPEGVideoH264FMO_MapType                       CtrlID = CidCodecBase + 372
	CidMPEGVideoH264FMO_SliceGroup                    CtrlID = CidCodecBase + 373
	CidMPEGVideoH264FMOChangeDirection                CtrlID = CidCodecBase + 374
	CidMPEGVideoH264FMOChangeRate                     CtrlID = CidCodecBase + 375
	CidMPEGVideoH264FMORunLength                      CtrlID = CidCodecBase + 376
	CidMPEGVideoH264ASO                               CtrlID = CidCodecBase + 377
	CidMPEGVideoH264ASOSliceOrder                     CtrlID = CidCodecBase + 378
	CidMPEGVideoH264HierarchicalCoding                CtrlID = CidCodecBase + 379
	CidMPEGVideoH264HierarchicalCodingType            CtrlID = CidCodecBase + 380
	CidMPEGVideoH264HierarchicalCodingLayer           CtrlID = CidCodecBase + 381
	CidMPEGVideoH264HierarchicalCodingLaYERQP         CtrlID = CidCodecBase + 382
	CidMPEGVideoH264ConstrainedIntraPrediction        CtrlID = CidCodecBase + 383
	CidMPEGVideoH264ChromaQPIndexOffset               CtrlID = CidCodecBase + 384
	CidMPEGVideoH264IFrameMinQP                       CtrlID = CidCodecBase + 385
	CidMPEGVideoH264IFrameMaxQP                       CtrlID = CidCodecBase + 386
	CidMPEGVideoH264PFrameMinQP                       CtrlID = CidCodecBase + 387
	CidMPEGVideoH264PFrameMaxQP                       CtrlID = CidCodecBase + 388
	CidMPEGVideoH264BFrameMinQP                       CtrlID = CidCodecBase + 389
	CidMPEGVideoH264BFrameMaxQP                       CtrlID = CidCodecBase + 390
	CidMPEGVideoH264HierCodingL0BR                    CtrlID = CidCodecBase + 391
	CidMPEGVideoH264HierCodingL1BR                    CtrlID = CidCodecBase + 392
	CidMPEGVideoH264HierCodingL2BR                    CtrlID = CidCodecBase + 393
	CidMPEGVideoH264HierCodingL3BR                    CtrlID = CidCodecBase + 394
	CidMPEGVideoH264HierCodingL4BR                    CtrlID = CidCodecBase + 395
	CidMPEGVideoH264HierCodingL5BR                    CtrlID = CidCodecBase + 396
	CidMPEGVideoH264HierCodingL6BR                    CtrlID = CidCodecBase + 397
	CidMPEGVideoMPEG4IFrameQP                         CtrlID = CidCodecBase + 400
	CidMPEGVideoMPEG4PFrameQP                         CtrlID = CidCodecBase + 401
	CidMPEGVideoMPEG4BFrameQP                         CtrlID = CidCodecBase + 402
	CidMPEGVideoMPEG4MinQP                            CtrlID = CidCodecBase + 403
	CidMPEGVideoMPEG4MaxQP                            CtrlID = CidCodecBase + 404
	CidMPEGVideoMPEG4Level                            CtrlID = CidCodecBase + 405
	CidMPEGVideoMPEG4Profile                          CtrlID = CidCodecBase + 406
	CidMPEGVideoMPEG4QPEL                             CtrlID = CidCodecBase + 407
	CidMPEGVideoVPXNumPartitions                      CtrlID = CidCodecBase + 500
	CidMPEGVideoVPX_IMD_DISABLE_4X4                   CtrlID = CidCodecBase + 501
	CidMPEGVideoVPXNum_REFFrameS                      CtrlID = CidCodecBase + 502
	CidMPEGVideoVPXFilterLevel                        CtrlID = CidCodecBase + 503
	CidMPEGVideoVPXFilterSharpness                    CtrlID = CidCodecBase + 504
	CidMPEGVideoVPXGoldenFrameRefPeriod               CtrlID = CidCodecBase + 505
	CidMPEGVideoVPXGoldenFrameSel                     CtrlID = CidCodecBase + 506
	CidMPEGVideoVPXMinQP                              CtrlID = CidCodecBase + 507
	CidMPEGVideoVPMaxQP                               CtrlID = CidCodecBase + 508
	CidMPEGVideoVPIFrameQP                            CtrlID = CidCodecBase + 509
	CidMPEGVideoVPXPFrameQP                           CtrlID = CidCodecBase + 510
	CidMPEGVideoVP8Profile                            CtrlID = CidCodecBase + 511
	CidMPEGVideoVPXProfile                            CtrlID = CidMPEGVideoVP8Profile
	CidMPEGVideoVP9Profile                            CtrlID = CidCodecBase + 512
	CidMPEGVideoVP9Level                              CtrlID = CidCodecBase + 513
	CidMPEGVideoHEVC_MinQP                            CtrlID = CidCodecBase + 600
	CidMPEGVideoHEVC_MaxQP                            CtrlID = CidCodecBase + 601
	CidMPEGVideoHEVC_IFrameQP                         CtrlID = CidCodecBase + 602
	CidMPEGVideoHEVC_PFrameQP                         CtrlID = CidCodecBase + 603
	CidMPEGVideoHEVC_BFrameQP                         CtrlID = CidCodecBase + 604
	CidMPEGVideoHEVC_HierQP                           CtrlID = CidCodecBase + 605
	CidMPEGVideoHEVC_HierCodingTYPE                   CtrlID = CidCodecBase + 606
	CidMPEGVideoHEVC_HierCodingLAYER                  CtrlID = CidCodecBase + 607
	CidMPEGVideoHEVC_HierCodingL0_QP                  CtrlID = CidCodecBase + 608
	CidMPEGVideoHEVC_HierCodingL1_QP                  CtrlID = CidCodecBase + 609
	CidMPEGVideoHEVC_HierCodingL2_QP                  CtrlID = CidCodecBase + 610
	CidMPEGVideoHEVC_HierCodingL3_QP                  CtrlID = CidCodecBase + 611
	CidMPEGVideoHEVC_HierCodingL4_QP                  CtrlID = CidCodecBase + 612
	CidMPEGVideoHEVC_HierCodingL5_QP                  CtrlID = CidCodecBase + 613
	CidMPEGVideoHEVC_HierCodingL6_QP                  CtrlID = CidCodecBase + 614
	CidMPEGVideoHEVCProfile                           CtrlID = CidCodecBase + 615
	CidMPEGVideoHEVCLevel                             CtrlID = CidCodecBase + 616
	CidMPEGVideoHEVCFrameRATE_RESOLUTION              CtrlID = CidCodecBase + 617
	CidMPEGVideoHEVC_TIER                             CtrlID = CidCodecBase + 618
	CidMPEGVideoHEVC_MaxPARTITION_DEPTH               CtrlID = CidCodecBase + 619
	CidMPEGVideoHEVC_LoopFILTERMode                   CtrlID = CidCodecBase + 620
	CidMPEGVideoHEVC_LF_BETAOffset_DIV2               CtrlID = CidCodecBase + 621
	CidMPEGVideoHEVC_LF_TCOffset_DIV2                 CtrlID = CidCodecBase + 622
	CidMPEGVideoHEVC_RefreshTYPE                      CtrlID = CidCodecBase + 623
	CidMPEGVideoHEVC_RefreshPeriod                    CtrlID = CidCodecBase + 624
	CidMPEGVideoHEVC_LOSSLESS_CU                      CtrlID = CidCodecBase + 625
	CidMPEGVideoHEVC_CONST_IntraPRED                  CtrlID = CidCodecBase + 626
	CidMPEGVdeoHEVC_WAVEFRONT                         CtrlID = CidCodecBase + 627
	CidMPEGVdeoHEVC_GENERAL_PB                        CtrlID = CidCodecBase + 628
	CidMPEGVideoHEVC_TEMPORAL_ID                      CtrlID = CidCodecBase + 629
	CidMPEGVideoHEVC_STRONG_SMOOTHING                 CtrlID = CidCodecBase + 630
	CidMPEGVideoHEVC_MaxNUM_MERGE_MV_MINUS1           CtrlID = CidCodecBase + 631
	CidMPEGVideoHEVC_IntraPU_SPLIT                    CtrlID = CidCodecBase + 632
	CidMPEGVideoHEVC_TMV_PREDICTION                   CtrlID = CidCodecBase + 633
	CidMPEGVideoHEVC_WITHOUT_STARTCODE                CtrlID = CidCodecBase + 634
	CidMPEGVideoHEVCSize_OFLength_FIELD               CtrlID = CidCodecBase + 635
	CidMPEGVideoHEVC_HierCodingL0_BR                  CtrlID = CidCodecBase + 636
	CidMPEGVideoHEVC_HierCodingL1_BR                  CtrlID = CidCodecBase + 637
	CidMPEGVideoHEVC_HierCodingL2_BR                  CtrlID = CidCodecBase + 638
	CidMPEGVideoHEVC_HierCodingL3_BR                  CtrlID = CidCodecBase + 639
	CidMPEGVideoHEVC_HierCodingL4_BR                  CtrlID = CidCodecBase + 640
	CidMPEGVideoHEVC_HierCodingL5_BR                  CtrlID = CidCodecBase + 641
	CidMPEGVideoHEVC_HierCodingL6_BR                  CtrlID = CidCodecBase + 642
	CidMPEGVideoREFNumBER_FOR_PFRAMES                 CtrlID = CidCodecBase + 643
	CidMPEGVideoPREPEND_SPSPPS_TO_IDR                 CtrlID = CidCodecBase + 644
	CidMPEGVideoCONSTANT_QUALITY                      CtrlID = CidCodecBase + 645
	CidMPEGVideoFrameSKIPMode                         CtrlID = CidCodecBase + 646
	CidMPEGVideoHEVC_IFrameMinQP                      CtrlID = CidCodecBase + 647
	CidMPEGVideoHEVC_IFrameMaxQP                      CtrlID = CidCodecBase + 648
	CidMPEGVideoHEVC_PFrameMinQP                      CtrlID = CidCodecBase + 649
	CidMPEGVideoHEVC_PFrameMaxQP                      CtrlID = CidCodecBase + 650
	CidMPEGVideoHEVC_BFrameMinQP                      CtrlID = CidCodecBase + 651
	CidMPEGVideoHEVC_BFrameMaxQP                      CtrlID = CidCodecBase + 652
	CidMPEGVideoDEC_DISPLAY_DELAY                     CtrlID = CidCodecBase + 653
	CidMPEGVideoDEC_DISPLAY_DELAY_ENABLE              CtrlID = CidCodecBase + 654
	CidCodecCX2341X_BASE                              CtrlID = CtrlID(CtrlClassCodec | 0x1000)
	CidMPEGCX2341XVideoSPATIALFilterMode              CtrlID = CidCodecCX2341X_BASE + 0
	CidMPEGCX2341XVideoSPATIALFilter                  CtrlID = CidCodecCX2341X_BASE + 1
	CidMPEGCX2341XVideoLUMA_SPATIALFilterTYPE         CtrlID = CidCodecCX2341X_BASE + 2
	CidMPEGCX2341XVideoChromaSPATIALFilterTYPE        CtrlID = CidCodecCX2341X_BASE + 3
	CidMPEGCX2341XVideoTEMPORALFilterMode             CtrlID = CidCodecCX2341X_BASE + 4
	CidMPEGCX2341XVideoTEMPORALFilter                 CtrlID = CidCodecCX2341X_BASE + 5
	CidMPEGCX2341XVideoMEDIANFilterTYPE               CtrlID = CidCodecCX2341X_BASE + 6
	CidMPEGCX2341XVideoLUMA_MEDIANFilterBOTTOM        CtrlID = CidCodecCX2341X_BASE + 7
	CidMPEGCX2341XVideoLUMA_MEDIANFilterTOP           CtrlID = CidCodecCX2341X_BASE + 8
	CidMPEGCX2341XVideoChromaMEDIANFilterBOTTOM       CtrlID = CidCodecCX2341X_BASE + 9
	CidMPEGCX2341XVideoChromaMEDIANFilterTOP          CtrlID = CidCodecCX2341X_BASE + 10
	CidMPEGCX2341X_STREAM_INSERT_NAV_PACKETS          CtrlID = CidCodecCX2341X_BASE + 11
	CidCodecMFC51_BASE                                CtrlID = CtrlID(CtrlClassCodec | 0x1100)
	CidMPEGMFC51VideoDecoderH264_DISPLAY_DELAY        CtrlID = CidCodecMFC51_BASE + 0
	CidMPEGMFC51VideoDecoderH264_DISPLAY_DELAY_ENABLE CtrlID = CidCodecMFC51_BASE + 1
	CidMPEGMFC51VideoFrameSKIPMode                    CtrlID = CidCodecMFC51_BASE + 2
	CidMPEGMFC51VideoForceFrameTYPE                   CtrlID = CidCodecMFC51_BASE + 3
	CidMPEGMFC51VideoPADDING                          CtrlID = CidCodecMFC51_BASE + 4
	CidMPEGMFC51VideoPADDING_YUV                      CtrlID = CidCodecMFC51_BASE + 5
	CidMPEGMFC51VideoRC_FIXED_TARGET_BIT              CtrlID = CidCodecMFC51_BASE + 6
	CidMPEGMFC51VideoRC_REACTION_COEFF                CtrlID = CidCodecMFC51_BASE + 7
	CidMPEGMFC51VideoH264ADAPTIVE_RC_ACTIVITY         CtrlID = CidCodecMFC51_BASE + 50
	CidMPEGMFC51VideoH264ADAPTIVE_RC_DARK             CtrlID = CidCodecMFC51_BASE + 51
	CidMPEGMFC51VideoH264ADAPTIVE_RC_SMOOTH           CtrlID = CidCodecMFC51_BASE + 52
	CidMPEGMFC51VideoH264ADAPTIVE_RC_STATIC           CtrlID = CidCodecMFC51_BASE + 53
	CidMPEGMFC51VideoH264NUM_REF_PIC_FOR_P            CtrlID = CidCodecMFC51_BASE + 54
)

// The video ioctl values.
const (
	VidIocQueryCap           uint32 = 0x80685600
	VidIocReserved           uint32 = 0x00005601
	VidIocEnumFmt            uint32 = 0xc0405602
	VidIocGFmt               uint32 = 0xc0d05604
	VidIocSFmt               uint32 = 0xc0d05605
	VidIocReqBufs            uint32 = 0xc0145608
	VidIocQueryBuf           uint32 = 0xc0585609
	VidIocGFBuf              uint32 = 0x8030560a
	VidIocSFBuf              uint32 = 0x4030560b
	VidIocOverlay            uint32 = 0x4004560e
	VidIocQBuf               uint32 = 0xc058560f
	VidIocExpBuf             uint32 = 0xc0405610
	VidIocDQBuf              uint32 = 0xc0585611
	VidIocStreamOn           uint32 = 0x40045612
	VidIocStreamOff          uint32 = 0x40045613
	VidIocGParm              uint32 = 0xc0cc5615
	VidIocSParm              uint32 = 0xc0cc5616
	VidIocGStd               uint32 = 0x80085617
	VidIocSStd               uint32 = 0x40085618
	VidIocEnumStd            uint32 = 0xc0485619
	VidIocEnumInput          uint32 = 0xc050561a
	VidIocGCtrl              uint32 = 0xc008561b
	VidIocSCtrl              uint32 = 0xc008561c
	VidIocGTuner             uint32 = 0xc054561d
	VidIocSTuner             uint32 = 0x4054561e
	VidIocGAudio             uint32 = 0x80345621
	VidIocSAudio             uint32 = 0x40345622
	VidIocQueryCtrl          uint32 = 0xc0445624
	VidIocQueryMenu          uint32 = 0xc02c5625
	VidIocGInput             uint32 = 0x80045626
	VidIocSInput             uint32 = 0xc0045627
	VidIocGEDID              uint32 = 0xc0285628
	VidIocSEDID              uint32 = 0xc0285629
	VidIocGOutput            uint32 = 0x8004562e
	VidIocSOutput            uint32 = 0xc004562f
	VidIocEnumOutput         uint32 = 0xc0485630
	VidIocGAudOut            uint32 = 0x80345631
	VidIocSAudOut            uint32 = 0x40345632
	VidIocGModulator         uint32 = 0xc0445636
	VidIocSModulator         uint32 = 0x40445637
	VidIocGFrequency         uint32 = 0xc02c5638
	VidIocSFrequency         uint32 = 0x402c5639
	VidIocCropCap            uint32 = 0xc02c563a
	VidIocGCrop              uint32 = 0xc014563b
	VidIocSCrop              uint32 = 0x4014563c
	VidIocGJpegComp          uint32 = 0x808c563d
	VidIocSJpegComp          uint32 = 0x408c563e
	VidIocQueryStd           uint32 = 0x8008563f
	VidIocTryFmt             uint32 = 0xc0d05640
	VidIocEnumAudio          uint32 = 0xc0345641
	VidIocEnumAudOut         uint32 = 0xc0345642
	VidIocGPriority          uint32 = 0x80045643
	VidIocSPriority          uint32 = 0x40045644
	VidIocGSlicedVBICap      uint32 = 0xc0745645
	VidIocLogStatus          uint32 = 0x00005646
	VidIocGExtCtrls          uint32 = 0xc0205647
	VidIocSExtCtrls          uint32 = 0xc0205648
	VidIocTryExtCtrls        uint32 = 0xc0205649
	VidIocEnumFrameSizes     uint32 = 0xc02c564a
	VidIocEnumFrameIntervals uint32 = 0xc034564b
	VidIocGEncIndex          uint32 = 0x8818564c
	VidIocEncoderCmd         uint32 = 0xc028564d
	VidIocTryEncoderCmd      uint32 = 0xc028564e
)

// StdID is the standard ID type.
type StdID uint64

// TunerCap is the tuner capability type.
type TunerCap uint32

// The tuner capabilities.
const (
	TunerCapLow TunerCap = 1 << iota
	TunerCapNorm
	TunerCapHWSeekBounded
	TunerCapHWSeekWrap
	TunerCapStereo
	TunerCapLang2
	TunerCapLang1
	TunerCapRDS
	TunerCapRDSBlockIO
	TunerCapRDSControls
	TunerCapFreqBands
	TunerCapHWSeekProgLim
	TunerCap1Hz
	TunerCapSAP = TunerCapLang2
)

// TunerMode is the tuner mode type.
type TunerMode uint32

// The tuner modes.
const (
	TunerModeMono TunerMode = iota
	TunerModeStereo
	TunerModeLang2
	TunerModeLang1
	TunerModeLang1Lang2
	TunerModeSAP = TunerModeLang2
)

// TunerSub is the tuner RX sub-channel type.
type TunerSub uint32

// The tuner RX sub-channel flags.
const (
	TunerSubMono TunerSub = 1 << iota
	TunerSubStereo
	TunerSubLang2
	TunerSubLang1
	TunerSubRDS
	TunerSubSAP = TunerSubLang2
)

// TunerType is the tuner type type.
type TunerType uint32

// The tuner types.
const (
	TunerTypeRadio TunerType = iota + 1
	TunerTypeAnalogTV
	_
	TunerTypeSDR
	TunerTypeRF
)

// VBIFmtFlag is the VBI format flag type.
type VBIFmtFlag uint32

// VBI format flags.
const (
	VBIFmtFlagUnsync VBIFmtFlag = 1 << iota
	VBIFmtFlagInterlaces
)

// XferFunc is the transfer function type.
type XferFunc uint32

// The transfer functions.
const (
	XferFuncDefault XferFunc = iota
	XferFunc709
	XferFuncSRGB
	XferFuncOPRGB
	XferFuncSMPTE240M
	XferFuncNone
	XferFuncDCIP3
	XferFuncSMPTE2084
)

// Audio is the v4l2 audio struct.
type Audio struct {
	Index      uint32
	Name       [32]byte
	Capability AudCap
	Mode       AudMode
	Reserved   [2]uint32
}

// Buffer is the v4l2 buffer struct.
type Buffer struct {
	Index     uint32
	Type      BufType
	BytesUsed uint32
	Flags     BufFlag
	Field     Field
	Timestamp syscall.Timeval
	Timecode  Timecode
	Sequence  uint32
	Memory    Memory
	M         uintptr
	Length    uint32
	Reserved2 uint32
	RequestFD uint32
}

// Capability is the v4l2 capability struct.
type Capability struct {
	Driver       [16]byte
	Card         [32]byte
	BusInfo      [32]byte
	Version      uint32
	Capabilities Cap
	DeviceCaps   Cap
	Reserved     [3]uint32
}

// Clip is the v4l2 clip struct.
type Clip struct {
	C    Rect
	Next uintptr
}

// Control is the v4l2 control.
type Control struct {
	ID    CtrlID
	Value int32
}

// CtrlFwhtparams is the v4l2 TODO.
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

// CtrlH264DecodeParams is the v4l2 TODO.
type CtrlH264DecodeParams struct {
	DPD                 [16]H264DPDEntry
	NumSlices           uint16
	NALRefIDC           uint16
	TopFieldOrderCnt    uint32
	BottomFieldOrderCnt uint32
	Flags               uint32
}

// CtrlHevcPps is the v4l2 TODO.
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

// CtrlHevcSps is the v4l2 TODO.
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

// CtrlMpeg2Quantization is the v4l2 TODO.
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

// CtrlMpeg2SliceParams is the v4l2 TODO.
type CtrlMpeg2SliceParams struct {
	BitSize            uint32
	DataBitOffset      uint32
	Sequence           Mpeg2Sequence
	Picture            Mpeg2Picture
	ForwardRefTS       uint64
	QuantiserScaleCode uint32
}

// CtrlVP8FrameHeader is the v4l2 TODO.
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

// FmtDesc is the v4l2 fmtdesc.
type FmtDesc struct {
	Index       uint32
	Type        BufType
	Flags       FmtFlag
	Description [32]byte
	PixFormat   PixFmt
	Reserved    [4]uint32
}

// Format is the v4l2 format.
type Format struct {
	Type    BufType
	_       uint32
	RawData [200]byte // Union of several possible types.
}

// FrameSizeDiscrete is v4l2Framesize_discrete.
type FrameSizeDiscrete struct {
	Width  uint32
	Height uint32
}

// FrameSizeEnum is v4l2Framesizeenum.
type FrameSizeEnum struct {
	Index     uint32
	PixFormat PixFmt
	Type      FrmSizeType
	M         [24]byte // Union
	Reserved  [2]uint32
}

// FrameSizeStepwise is v4l2Framesize_stepwise.
type FrameSizeStepwise struct {
	MinWidth   uint32
	MaxWidth   uint32
	StepWidth  uint32
	MinHeight  uint32
	MaxHeight  uint32
	StepHeight uint32
}

// Frequency is v4l2Frequency.
type Frequency struct {
	Tuner     uint32
	Type      uint32
	Frequency uint32
	Reserved  [4]uint32
}

// H264DPDEntry is the v4l2 TODO.
type H264DPDEntry struct {
	ReferenceTS         uint64
	FrameNum            uint16
	PicNum              uint16
	TopFieldOrderCnt    int32
	BottomFieldOrderCnt int32
	Flags               uint32
}

// H264PredWeightTable is the v4l2 TODO.
type H264PredWeightTable struct {
	LumaLog2WeightDenom   uint16
	ChromaLog2WeightDenom uint16
	Weightfactors         [2]H264WeightFactors
}

// H264WeightFactors is the v4l2 TODO.
type H264WeightFactors struct {
	LumaWeight   [32]int16
	LumaOffset   [32]int16
	ChromaWeight [32]int16
	ChromaOffset [32]int16
}

// Input is the v4l2 input.
type Input struct {
	Index        uint32
	Name         [32]byte
	Type         InputType
	AudioSet     uint32
	Tuner        uint32
	Standard     StdID
	Status       InputStatus
	Capabilities InputCap
	Reserved     [3]uint32
}

// Modulator is the v4l2 modulator.
type Modulator struct {
	Index      uint32
	Name       [32]byte
	Capability TunerCap
	RangeLow   uint32
	RangeHigh  uint32
	TXSubChans TunerSub
	Type       TunerType
	Reserved   [3]uint32
}

// Mpeg2Picture is the v4l2 TODO.
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

// Mpeg2Sequence is the v4l2 TODO.
type Mpeg2Sequence struct {
	HorizontalSize            uint16
	VerticalSize              uint16
	VBVBufferSize             uint32
	ProfileAndLevelIndication uint16
	ProgressiveSequence       uint8
	ChromaFormat              uint8
}

// Output is the v4l2 output.
type Output struct {
	Index        uint32
	Name         [32]byte
	Type         OutputType
	AudioSet     uint32
	Modulator    uint32
	Standard     StdID
	Capabilities OutputCap
	Reserved     [3]uint32
}

// PixFormat is the v4l2 pix format.
type PixFormat struct {
	Width        uint32
	Height       uint32
	PixFormat    PixFmt
	Field        Field
	BytesPerLine uint32
	SizeImage    uint32
	ColorSpace   ColorSpace
	Priv         uint32
	Flags        PixFmtFlag
	M            uint32 // Anonymous union of YCbCr and HSV
	Quantization Quantization
	XferFunc     XferFunc
}

// PixFormatMPlane is the v4l2 pix format_mplane.
type PixFormatMPlane struct {
	Width        uint32
	Height       uint32
	PixFormat    PixFmt
	Field        Field
	ColorSpace   ColorSpace
	PlaneFmt     [8]PlanePixFormat
	NumPlanes    uint8
	Flags        PixFmtFlag
	M            uint32 // Anonymous union of YCbCr and HSV
	Quantization Quantization
	XferFunc     XferFunc
	Reserved     [7]uint8
}

// Plane is the v4l2 plane.
type Plane struct {
	BytesUsed  uint32
	Length     uint32
	M          uint32 // Union
	DataOffset uint32
	Reserved   [11]uint32
}

// PlanePixFormat is the v4l2 plane pix format.
type PlanePixFormat struct {
	SizeImage    uint32
	BytesPerLine uint32
	Reserved     [6]uint16
}

// QueryCtrl is the v4l2 query control struct.
type QueryCtrl struct {
	ID           CtrlID
	Type         CtrlType
	Name         [32]byte
	Minimum      int32
	Maximum      int32
	Step         int32
	DefaultValue int32
	Flags        uint32
	Reserved     [2]uint32
}

// QueryExtCtrl is the v4l2 query extended control struct.
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
	Reserved     [32]uint32
}

// QueryMenu is an encapsulation of a menu.
type QueryMenu struct {
	ID    CtrlID
	Index uint32
	Name  [32]byte
	// Value int64 unioned with Name
	Reserved uint32
}

// Rect is the v4l2 rect.
type Rect struct {
	Left   int32
	Top    int32
	Width  uint32
	Height uint32
}

// RequestBuffers is the v4l2 requestbuffers.
type RequestBuffers struct {
	Count        uint32
	Type         BufType
	Memory       Memory
	Capabilities Cap
	Reserved     [1]uint32
}

// SlicedVBIFormat is the v4l2 sliced_vbi_format.
type SlicedVBIFormat struct {
	ServiceSet   uint32
	ServiceLines [2][24]SlicedVBIService
	IOSize       uint32
	Reserved     [2]uint32
}

// Timecode is the v4l2 timecode.
type Timecode struct {
	Type     TcType
	Flags    TcFlag
	Frames   uint8
	Seconds  uint8
	Minutes  uint8
	Hours    uint8
	UserBits [4]uint8
}

// Tuner is the v4l2 tuner.
type Tuner struct {
	Index      uint32
	Name       [32]byte
	Type       TunerType
	Capability TunerCap
	RangeLow   uint32
	RangeHigh  uint32
	RXSubChans TunerSub
	AudMode    TunerMode
	Signal     uint32
	AFC        int32
	Reserved   [4]uint32
}

// VBIFormat is the v4l2 vbi_format.
type VBIFormat struct {
	SamplingRate   uint32
	Offset         uint32
	SamplesPerLine uint32
	Start          uint32
	Count          uint32
	Flags          VBIFmtFlag
	Reserved       uint32
}

// VP8EntropyCoderState is the v4l2 TODO.
type VP8EntropyCoderState struct {
	Range    uint8
	Value    uint8
	BitCount uint8
	Padding  uint8
}

// VP8EntropyHeader is the v4l2 TODO.
type VP8EntropyHeader struct {
	CoeffProbs  [4][8][3][11]uint8
	YModeProbs  [4]uint8
	UVModeProbs [4]uint8
	MVProb      [2][19]uint8
	Padding     [3]uint8
}

// VP8LoopfilterHeader is the v4l2 TODO.
type VP8LoopfilterHeader struct {
	RefFrmDelta    [4]int8
	MBModeDelta    [4]int8
	SharpnessLevel uint8
	Padding        uint16
	Flags          uint32
}

// VP8QuantizationHeader is the v4l2 TODO.
type VP8QuantizationHeader struct {
	YACQi     uint8
	YDCDelta  int8
	Y2DCDelta int8
	Y2ACDelta int8
	UVDCDelta int8
	UVACDelta int8
	Padding   uint16
}

// VP8SegmentHeader is the v4l2 TODO.
type VP8SegmentHeader struct {
	QuantUpdate  [4]int8
	LFUpdate     [4]int8
	SegmentProbs [3]uint8
	Padding      uint8
	Flags        uint32
}

// Window is the v4l2 window.
type Window struct {
	W           Rect
	Field       Field
	ChromaKey   uint32
	Clips       *Clip
	ClipCount   uint32
	Bitmap      *interface{}
	GlobalAlpha uint8
}

// QueryCapabilities queries the device capabilities.
func QueryCapabilities(fd int) (*Capability, error) {
	capability := &Capability{}
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(VidIocQueryCap), uintptr(unsafe.Pointer(capability))); err != 0 {
		return nil, err
	}
	return capability, nil
}

// EnumFormats enumerates the supported formats.
func EnumFormats(fd int, bufType BufType) ([]*FmtDesc, error) {
	var index uint32 = 0
	fmtDescs := make([]*FmtDesc, 0, 8)
	for {
		fmtDesc := &FmtDesc{}
		fmtDesc.Index = index
		fmtDesc.Type = bufType
		_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(VidIocEnumFmt), uintptr(unsafe.Pointer(fmtDesc)))
		if err != 0 {
			if err == syscall.EINVAL {
				break
			}
			return nil, err
		}
		fmtDescs = append(fmtDescs, fmtDesc)
		index++
	}
	return fmtDescs, nil
}

// EnumFrameSizes enumerates the available frame sizes for a pixel format.
func EnumFrameSizes(fd int, pixFormat PixFmt) ([]*FrameSizeEnum, error) {
	var index uint32 = 0
	frameSizeEnums := make([]*FrameSizeEnum, 0, 4)
	for {
		frameSizeEnum := &FrameSizeEnum{}
		frameSizeEnum.Index = index
		frameSizeEnum.PixFormat = pixFormat
		_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(VidIocEnumFrameSizes), uintptr(unsafe.Pointer(frameSizeEnum)))
		if err != 0 {
			if err == syscall.EINVAL {
				break
			}
			return nil, err
		}
		frameSizeEnums = append(frameSizeEnums, frameSizeEnum)
		index++
	}
	return frameSizeEnums, nil
}

// QueryControls queries the controls.
func QueryControls(fd int) ([]*QueryCtrl, error) {
	controls := make([]*QueryCtrl, 0, 4)
	id := CtrlID(CtrlFlagNextCtrl)
	for {
		queryCtrl := &QueryCtrl{}
		queryCtrl.ID = id
		_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(VidIocQueryCtrl), uintptr(unsafe.Pointer(queryCtrl)))
		if err != 0 {
			if err == syscall.EINVAL {
				break
			}
			return nil, err
		}
		controls = append(controls, queryCtrl)
		id = queryCtrl.ID | CtrlID(CtrlFlagNextCtrl)
	}
	return controls, nil
}

// QueryMenus queries all menus for a particular control ID.
func QueryMenus(fd int, id CtrlID) ([]*QueryMenu, error) {
	var index uint32 = 0
	menus := make([]*QueryMenu, 0, 4)
	for {
		queryMenu := &QueryMenu{}
		queryMenu.ID = id
		queryMenu.Index = index
		_, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(VidIocQueryMenu), uintptr(unsafe.Pointer(queryMenu)))
		if err != 0 {
			if err == syscall.EINVAL {
				break
			}
			return nil, err
		}
		menus = append(menus, queryMenu)
		index++
	}
	return menus, nil
}

// GetFormat returns the current format.
func GetFormat(fd int, bufType BufType) (*Format, error) {
	format := &Format{}
	format.Type = bufType
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(VidIocGFmt), uintptr(unsafe.Pointer(format))); err != 0 {
		return nil, err
	}
	return format, nil
}

// SetFormat sets the format and frame size.
func SetFormat(fd int, bufType BufType, pixFormat PixFmt, width uint32, height uint32) (uint32, uint32, error) {
	format := &Format{}
	format.Type = bufType
	pix := (*PixFormat)(unsafe.Pointer(&format.RawData[0]))
	pix.Width = width
	pix.Height = height
	pix.PixFormat = pixFormat
	pix.Field = FieldNone
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(VidIocSFmt), uintptr(unsafe.Pointer(format))); err != 0 {
		return 0, 0, err
	}
	return pix.Width, pix.Height, nil
}

func GetControl(fd int, id CtrlID) (*Control, error) {
	control := &Control{}
	control.ID = id
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(VidIocGCtrl), uintptr(unsafe.Pointer(control))); err != 0 {
		return nil, err
	}
	return control, nil
}

func SetControl(fd int, control *Control) error {
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(VidIocSCtrl), uintptr(unsafe.Pointer(control))); err != 0 {
		return err
	}
	return nil
}

// RequestDriverBuffers requests driver buffers.
func RequestDriverBuffers(fd int, count uint32, bufType BufType, memory Memory) (uint32, error) {
	requestBuffers := &RequestBuffers{}
	requestBuffers.Count = count
	requestBuffers.Type = bufType
	requestBuffers.Memory = memory
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(VidIocReqBufs), uintptr(unsafe.Pointer(requestBuffers))); err != 0 {
		return 0, err
	}
	return requestBuffers.Count, nil
}

// QueryBuffer queries a buffer.
func QueryBuffer(fd int, index uint32, bufType BufType, memory Memory) (*Buffer, error) {
	buffer := &Buffer{}
	buffer.Index = index
	buffer.Type = bufType
	buffer.Memory = memory
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(VidIocQueryBuf), uintptr(unsafe.Pointer(buffer))); err != 0 {
		return nil, err
	}
	return buffer, nil
}

// EnqueueBuffer enqueues a buffer.
func EnqueueBuffer(fd int, buffer *Buffer) error {
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(VidIocQBuf), uintptr(unsafe.Pointer(buffer))); err != 0 {
		return err
	}
	return nil
}

// DequeueBuffer dequeues a buffer.
func DequeueBuffer(fd int, bufType BufType, memory Memory) (*Buffer, error) {
	buffer := &Buffer{}
	buffer.Type = bufType
	buffer.Memory = memory
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(VidIocDQBuf), uintptr(unsafe.Pointer(buffer))); err != 0 {
		return nil, err
	}
	return buffer, nil
}

// StreamOn turns on Streaming for the specified buffer type.
func StreamOn(fd int, bufType BufType) error {
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(VidIocStreamOn), uintptr(unsafe.Pointer(&bufType))); err != 0 {
		return err
	}
	return nil
}

// StreamOff turns off Streaming for the specified buffer type.
func StreamOff(fd int, bufType BufType) error {
	if _, _, err := syscall.Syscall(syscall.SYS_IOCTL, uintptr(fd), uintptr(VidIocStreamOff), uintptr(unsafe.Pointer(&bufType))); err != 0 {
		return err
	}
	return nil
}

// GrabFrame grabs a single frame.
func GrabFrame(fd int, bufType BufType, memory Memory, buffers [][]byte) ([]byte, error) {
	fdSet := &syscall.FdSet{}
	fdSet.Bits[fd/32] |= 1 << (uint(fd) % 32)
	timeout := &syscall.Timeval{
		Sec:  2,
		Usec: 0,
	}
	if _, err := syscall.Select(fd+1, fdSet, nil, nil, timeout); err != nil {
		return nil, err
	}
	buffer, err := DequeueBuffer(fd, bufType, memory)
	if err != nil {
		return nil, err
	}
	data := buffers[buffer.Index]
	frame := make([]byte, buffer.BytesUsed)
	copy(frame, data)
	err = EnqueueBuffer(fd, buffer)
	if err != nil {
		return nil, err
	}
	return frame, nil
}

// MmapBuffers memory maps buffers.
// The buffers must have been requested with a memory type of MemoryMmap.
func MmapBuffers(fd int, count uint32, bufType BufType) ([][]byte, error) {
	var index uint32
	buffers := make([][]byte, 0)
	for index = 0; index < count; index++ {
		buffer, err := QueryBuffer(fd, index, bufType, MemoryMmap)
		if err != nil {
			return nil, err
		}
		offset := int64(buffer.M)
		length := int(buffer.Length)
		data, err := unix.Mmap(fd, offset, length, unix.PROT_READ|unix.PROT_WRITE, unix.MAP_SHARED)
		if err != nil {
			return nil, err
		}
		buffers = append(buffers, data)
		if err := EnqueueBuffer(fd, buffer); err != nil {
			return nil, err
		}
	}
	return buffers, nil
}

// MunmapBuffers memory unmaps previously mapped driver buffers.
func MunmapBuffers(buffers [][]byte) error {
	for _, data := range buffers {
		if err := unix.Munmap(data); err != nil {
			return err
		}
	}
	return nil
}

// BytesToString converts a low-level, null-terminated C-string to a string.
func BytesToString(b []byte) string {
	if n := bytes.IndexByte(b, 0); n <= 0 {
		return ""
	} else {
		return string(b[:n])
	}
}
