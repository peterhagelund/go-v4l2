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

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/unix"
)

// AudCap is the audio capability type.
type AudCap uint32

// Audio capability flags (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-audio.html#audio-capability).
const (
	AudCapStereo AudCap = 1 << iota
	AudCapAVL
)

// AudMode is the audio mode type.
type AudMode uint32

// Audio mode flags (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-audio.html#audio-mode).
const (
	AudModeAVL = 1 << iota
)

// ColorSpace is the the color space type.
type ColorSpace uint32

// Color spaces (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/colorspaces-defs.html#c.v4l2_colorspace).
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

// Buffer capabilities (TODO).
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

// Buffer flags (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/buffer.html#buffer-flags).
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

// Buffer types (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/buffer.html#c.v4l2_buf_type).
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

// Device capability flags (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-querycap.html#device-capabilities).
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
	capReadWrite
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

// Field types (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/field-order.html#c.v4l2_field).
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

// Format flags (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-enum-fmt.html#fmtdesc-flags).
const (
	FmtFlagCompressed FmtFlag = 1 << iota
	FmtFlagEmulated
	FmtFlagContinuousByteStream
	FmtFlagDynResolution
)

// Frame size types (TODO).
const (
	FrmSizeTypeDiscrete uint32 = iota + 1
	FrmSizeTypeContinuous
	FrmSizeTypeStepwise
)

// InputCap is the input capabilities type.
type InputCap uint32

// Input capabilities (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-enuminput.html#input-capabilities).
const (
	_ InputCap = 1 << iota
	InputCapDVTimings
	InputCapStd
	InputCapNativeSize
)

// InputStatus is the input status type.
type InputStatus uint32

// Input statuses (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-enuminput.html#input-status).
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

// Input types (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-enuminput.html#input-type).
const (
	InputTypeTuner InputType = iota + 1
	InputTypeCamera
	InputTypeTouch
)

// Memory is the memory type.
type Memory uint32

// Memory types (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/buffer.html#c.v4l2_memory).
const (
	MemoryMmap Memory = iota + 1
	MemoryUserPtr
	MemoryOverlay
	MemoryDMABuf
)

// OutputCap is output capabilities type.
type OutputCap uint32

// Output capabilities (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-enumoutput.html#output-capabilities).
const (
	_ OutputCap = 1 << iota
	OutputCapDVTimings
	OutputCapStd
	OutputCapNativeSize
)

// OutputType is output type type.
type OutputType uint32

// Output types (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-enumoutput.html#output-type).
const (
	OutputTypeModulator = 1 + iota
	OutputTypeAnalog
	OutputTypeAnalogVGAOverlay
)

// PixFmt is the pixel format type.
type PixFmt uint32

// Pixel formats
// (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/pixfmt-rgb.html#pixfmt-rgb
// https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/yuv-formats.html#yuv-formats
// https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/pixfmt-reserved.html#reserved-formats).
// Absent macros in Go and given consts must be compile-time constants, all values are coded with "|" and "<<".
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

// Pixel format flags (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/pixfmt-reserved.html#format-flags).
const (
	PixFmtFlagPremulAlpha = 1 << iota
)

// Quantization is the quantization type.
type Quantization uint32

// The quantization flags (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/colorspaces-defs.html#c.v4l2_quantization).
const (
	QuantizationDefault Quantization = iota
	QuantizationFullRange
	QuantizationLimRange
)

// SlicedVBIService is the sliced VBI service type.
type SlicedVBIService uint16

// The sliced VBI services (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/dev-sliced-vbi.html#vbi-services2).
const (
	SlicedVBIServiceTeleTextB SlicedVBIService = 0x0001
)

// TcFlag is the timecode flag type.
type TcFlag uint32

// Timecode flags (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/buffer.html#timecode-flags).
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

// Timecode types (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/buffer.html#timecode-type).
const (
	TcType24Fps TcType = iota + 1
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

// TunerCap is the tuner capability type.
type TunerCap uint32

// The tuner capabilities (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-tuner.html#tuner-capability).
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

// The tuner modes (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-tuner.html#tuner-audmode).
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

// The tuner RX sub-channel flags (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-tuner.html#tuner-rxsubchans).
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

// The tuner types (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-tuner.html#c.v4l2_tuner_type).
const (
	TunerTypeRadio TunerType = iota + 1
	TunerTypeAnalogTV
	_
	TunerTypeSDR
	TunerTypeRF
)

// VBIFmtFlag is the VBI format flag type.
type VBIFmtFlag uint32

// VBI format flags (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/dev-raw-vbi.html#vbifmt-flags).
const (
	VBIFmtFlagUnsync VBIFmtFlag = 1 << iota
	VBIFmtFlagInterlaces
)

// XferFunc is the transfer function type.
type XferFunc uint32

// The transfer functions (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/colorspaces-defs.html#c.v4l2_xfer_func).
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

// Audio is v4l2_audio (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-audio.html#c.v4l2_audio).
type Audio struct {
	Index      uint32
	Name       [32]byte
	Capability AudCap
	Mode       AudMode
	Reserved   [2]uint32
}

// Buffer is v4l2_buffer (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/buffer.html#c.v4l2_buffer).
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
	M         uint32 // Union - Note: might differ between 32-bit and 64-bit systems. Investigate.
	Length    uint32
	Reserved2 uint32
	RequestFD uint32
}

// Capability is v4l2_capability (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-querycap.html#c.v4l2_capability).
type Capability struct {
	Driver       [16]byte
	Card         [32]byte
	BusInfo      [32]byte
	Version      uint32
	Capabilities Cap
	DeviceCaps   Cap
	Reserved     [3]uint32
}

// Clip is v4l2_clip (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/dev-overlay.html?highlight=v4l2_clip#c.v4l2_clip).
type Clip struct {
	C    Rect
	Next unsafe.Pointer
}

// CtrlFwhtparams is v4l2_TODO
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

// CtrlH264DecodeParams is v4l2_TODO
type CtrlH264DecodeParams struct {
	DPD                 [16]H264DPDEntry
	NumSlices           uint16
	NALRefIDC           uint16
	TopFieldOrderCnt    uint32
	BottomFieldOrderCnt uint32
	Flags               uint32
}

// CtrlHevcPps is v4l2_TODO
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

// CtrlHevcSps is v4l2_TODO
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

// CtrlMpeg2Quantization is v4l2_TODO
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

// CtrlMpeg2SliceParams is v4l2_TODO
type CtrlMpeg2SliceParams struct {
	BitSize            uint32
	DataBitOffset      uint32
	Sequence           Mpeg2Sequence
	Picture            Mpeg2Picture
	ForwardRefTS       uint64
	QuantiserScaleCode uint32
}

// CtrlVP8FrameHeader is v4l2_TODO
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

// FmtDesc is v4l2_fmtdesc (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-enum-fmt.html#c.v4l2_fmtdesc).
type FmtDesc struct {
	Index       uint32
	Type        BufType
	Flags       FmtFlag
	Description [32]byte
	PixFormat   PixFmt
	Reserved    [4]uint32
}

// Format is v4l2_format (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-fmt.html#c.v4l2_format).
type Format struct {
	Type    BufType
	RawData [256]byte // Union of several possible types.
}

// FrameSizeDiscrete is v4l2_framesize_discrete (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-enum-framesizes.html#c.v4l2_frmsize_discrete)).
type FrameSizeDiscrete struct {
	Width  uint32
	Height uint32
}

// FrameSizeEnum is v4l2_framesizeenum (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-enum-framesizes.html#c.v4l2_frmsizeenum)).
type FrameSizeEnum struct {
	Index     uint32
	PixFormat PixFmt
	Type      uint32
	M         [24]byte // Union
	Reserved  [2]uint32
}

// FrameSizeStepwise is v4l2_framesize_stepwise (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-enum-framesizes.html#c.v4l2_frmsize_stepwise).
type FrameSizeStepwise struct {
	MinWidth   uint32
	MaxWidth   uint32
	StepWidth  uint32
	MinHeight  uint32
	MaxHeight  uint32
	StepHeight uint32
}

// Frequency is v4l2_frequency (TODO).
type Frequency struct {
	Tuner     uint32
	Type      uint32
	Frequency uint32
	Reserved  [4]uint32
}

// H264DPDEntry is v4l2_TODO
type H264DPDEntry struct {
	ReferenceTS         uint64
	FrameNum            uint16
	PicNum              uint16
	TopFieldOrderCnt    int32
	BottomFieldOrderCnt int32
	Flags               uint32
}

// H264PredWeightTable is v4l2_TODO
type H264PredWeightTable struct {
	LumaLog2WeightDenom   uint16
	ChromaLog2WeightDenom uint16
	Weightfactors         [2]H264WeightFactors
}

// H264WeightFactors is v4l2_TODO
type H264WeightFactors struct {
	LumaWeight   [32]int16
	LumaOffset   [32]int16
	ChromaWeight [32]int16
	ChromaOffset [32]int16
}

// Input is v4l2_input (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-enuminput.html#c.v4l2_input).
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

// Modulator is v4l2_modulator (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-modulator.html#c.v4l2_modulator).
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

// Mpeg2Picture is v4l2_TODO
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

// Mpeg2Sequence is v4l2_TODO
type Mpeg2Sequence struct {
	HorizontalSize            uint16
	VerticalSize              uint16
	VBVBufferSize             uint32
	ProfileAndLevelIndication uint16
	ProgressiveSequence       uint8
	ChromaFormat              uint8
}

// Output is v4l2_output (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-enumoutput.html#c.v4l2_output).
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

// PixFormat is v4l2_pix_format (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/pixfmt-v4l2.html#c.v4l2_pix_format).
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

// PixFormatMPlane is v4l2_pix_format_mplane (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/pixfmt-v4l2-mplane.html#c.v4l2_pix_format_mplane).
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

// Plane is v4l2_plane (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/buffer.html?highlight=v4l2_plane#c.v4l2_plane).
type Plane struct {
	BytesUsed  uint32
	Length     uint32
	M          uint32 // Union
	DataOffset uint32
	Reserved   [11]uint32
}

// PlanePixFormat is v4l2_plane_pix_format (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/pixfmt-v4l2-mplane.html#c.v4l2_plane_pix_format).
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

// Rect is v4l2_rect (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/dev-overlay.html?highlight=v4l2_rect#c.v4l2_rect).
type Rect struct {
	Left   int32
	Top    int32
	Width  uint32
	Height uint32
}

// RequestBuffers is v4l2_requestbuffers (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-reqbufs.html#c.v4l2_requestbuffers).
type RequestBuffers struct {
	Count        uint32
	Type         BufType
	Memory       Memory
	Capabilities Cap
	Reserved     [1]uint32
}

// SlicedVBIFormat is v4l2_sliced_vbi_format (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/dev-sliced-vbi.html#c.v4l2_sliced_vbi_format).
type SlicedVBIFormat struct {
	ServiceSet   uint32
	ServiceLines [2][24]SlicedVBIService
	IOSize       uint32
	Reserved     [2]uint32
}

// Timecode is v4l2_timecode (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/buffer.html#c.v4l2_timecode).
type Timecode struct {
	Type     TcType
	Flags    TcFlag
	Frames   uint8
	Seconds  uint8
	Minutes  uint8
	Hours    uint8
	UserBits [4]uint8
}

// Tuner is v4l2_tuner (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/vidioc-g-tuner.html#c.v4l2_tuner).
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

// VBIFormat is v4l2_vbi_format (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/dev-raw-vbi.html#c.v4l2_vbi_format).
type VBIFormat struct {
	SamplingRate   uint32
	Offset         uint32
	SamplesPerLine uint32
	Start          uint32
	Count          uint32
	Flags          VBIFmtFlag
	Reserved       uint32
}

// VP8EntropyCoderState is v4l2_TODO
type VP8EntropyCoderState struct {
	Range    uint8
	Value    uint8
	BitCount uint8
	Padding  uint8
}

// VP8EntropyHeader is v4l2_TODO
type VP8EntropyHeader struct {
	CoeffProbs  [4][8][3][11]uint8
	YModeProbs  [4]uint8
	UVModeProbs [4]uint8
	MVProb      [2][19]uint8
	Padding     [3]uint8
}

// VP8LoopfilterHeader is v4l2_TODO
type VP8LoopfilterHeader struct {
	RefFrmDelta    [4]int8
	MBModeDelta    [4]int8
	SharpnessLevel uint8
	Padding        uint16
	Flags          uint32
}

// VP8QuantizationHeader is v4l2_TODO
type VP8QuantizationHeader struct {
	YACQi     uint8
	YDCDelta  int8
	Y2DCDelta int8
	Y2ACDelta int8
	UVDCDelta int8
	UVACDelta int8
	Padding   uint16
}

// VP8SegmentHeader is v4l2_TODO
type VP8SegmentHeader struct {
	QuantUpdate  [4]int8
	LFUpdate     [4]int8
	SegmentProbs [3]uint8
	Padding      uint8
	Flags        uint32
}

// Window is v4l2_window (https://www.linuxtv.org/downloads/v4l-dvb-apis-new/uapi/v4l/dev-overlay.html#c.v4l2_window).
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
	if err := Ioctl(fd, VidIocQueryCap, uintptr(unsafe.Pointer(capability))); err != nil {
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
		err := Ioctl(fd, VidIocEnumFmt, uintptr(unsafe.Pointer(fmtDesc)))
		if err != nil {
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
		err := Ioctl(fd, VidIocEnumFrameSizes, uintptr(unsafe.Pointer(frameSizeEnum)))
		if err != nil {
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

// SetFormat sets the format and frame size.
func SetFormat(fd int, bufType BufType, pixFormat PixFmt, width uint32, height uint32) (uint32, uint32, error) {
	format := &Format{}
	format.Type = bufType
	pix := (*PixFormat)(unsafe.Pointer(&format.RawData[0]))
	pix.Width = width
	pix.Height = height
	pix.PixFormat = pixFormat
	pix.Field = FieldNone
	if err := Ioctl(fd, VidIocSFmt, uintptr(unsafe.Pointer(format))); err != nil {
		return 0, 0, err
	}
	return pix.Width, pix.Height, nil
}

// RequestDriverBuffers requests driver buffers.
func RequestDriverBuffers(fd int, count uint32, bufType BufType, memory Memory) (uint32, error) {
	requestBuffers := &RequestBuffers{}
	requestBuffers.Count = count
	requestBuffers.Type = bufType
	requestBuffers.Memory = memory
	if err := Ioctl(fd, VidIocReqBufs, uintptr(unsafe.Pointer(requestBuffers))); err != nil {
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
	if err := Ioctl(fd, VidIocQueryBuf, uintptr(unsafe.Pointer(buffer))); err != nil {
		return nil, err
	}
	return buffer, nil
}

// MmapBuffers memory maps buffers.
// The buffers must have been requested with a memory type of MemoryMMap.
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

// EnqueueBuffer enqueues a buffer.
func EnqueueBuffer(fd int, buffer *Buffer) error {
	return Ioctl(fd, VidIocQBuf, uintptr(unsafe.Pointer(buffer)))
}

// DequeueBuffer dequeues a buffer.
func DequeueBuffer(fd int, bufType BufType, memory Memory) (*Buffer, error) {
	buffer := &Buffer{}
	buffer.Type = bufType
	buffer.Memory = memory
	if err := Ioctl(fd, VidIocDQBuf, uintptr(unsafe.Pointer(buffer))); err != nil {
		return nil, err
	}
	return buffer, nil
}

// StreamOn turns on streaming for the specified buffer type.
func StreamOn(fd int, bufType BufType) error {
	return Ioctl(fd, VidIocStreamOn, uintptr(unsafe.Pointer(&bufType)))
}

// StreamOff turns off streaming for the specified buffer type.
func StreamOff(fd int, bufType BufType) error {
	return Ioctl(fd, VidIocStreamOff, uintptr(unsafe.Pointer(&bufType)))
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
