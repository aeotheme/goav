// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

//Package swscale performs highly optimized image scaling and colorspace and pixel format conversion operations.
//Rescaling: is the process of changing the video size. Several rescaling options and algorithms are available.
//Pixel format conversion: is the process of converting the image format and colorspace of the image.
package swscale

//#cgo pkg-config: libswscale libavutil
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <string.h>
//#include <stdint.h>
//#include <libswscale/swscale.h>
//#include <libavutil/frame.h>
import "C"
import (
	"github.com/aeotheme/goav/avutil"
	"unsafe"
)

type (
	Context     C.struct_SwsContext
	Filter      C.struct_SwsFilter
	Vector      C.struct_SwsVector
	Class       C.struct_AVClass
	PixelFormat C.enum_AVPixelFormat
	Frame       C.struct_AVFrame
)

//Return the LIBSWSCALE_VERSION_INT constant.
func Version() uint {
	return uint(C.swscale_version())
}

//Return the libswscale build-time configuration.
func SwscaleConfiguration() string {
	return C.GoString(C.swscale_configuration())
}

//Return the libswscale license.
func SwscaleLicense() string {
	return C.GoString(C.swscale_license())
}

//Return a pointer to yuv<->rgb coefficients for the given colorspace suitable for sws_setColorspaceDetails().
func SwsGetcoefficients(c int) *int {
	return (*int)(unsafe.Pointer(C.sws_getCoefficients(C.int(c))))
}

//Return a positive value if pix_fmt is a supported input format, 0 otherwise.
func SwsIssupportedinput(p PixelFormat) int {
	return int(C.sws_isSupportedInput((C.enum_AVPixelFormat)(p)))
}

//Return a positive value if pix_fmt is a supported output format, 0 otherwise.
func SwsIssupportedoutput(p PixelFormat) int {
	return int(C.sws_isSupportedOutput((C.enum_AVPixelFormat)(p)))
}

func SwsIssupportedendiannessconversion(p PixelFormat) int {
	return int(C.sws_isSupportedEndiannessConversion((C.enum_AVPixelFormat)(p)))
}

////Scale the image slice in srcSlice and put the resulting scaled slice in the image in dst.
func SwsScale(ctx *Context, src *uint8, srcStride int32, y, h int, dst *uint8, dstStride int32) int {
	cCtx := (*C.struct_SwsContext)(unsafe.Pointer(ctx))
	cSrc := (*C.uint8_t)(unsafe.Pointer(src))
	cSrcStride := (*C.int)(unsafe.Pointer(&srcStride))
	cDst := (*C.uint8_t)(unsafe.Pointer(dst))
	cDstStride := (*C.int)(unsafe.Pointer(&dstStride))
	return int(C.sws_scale(cCtx, &cSrc, cSrcStride, C.int(y), C.int(h), &cDst, cDstStride))
}

func SwsScale2(ctxt *Context, srcData [8]*uint8, srcStride [8]int32, y, h int, dstData [8]*uint8, dstStride [8]int32) int {
	cCtx := (*C.struct_SwsContext)(unsafe.Pointer(ctxt))
	cSrc := (**C.uint8_t)(unsafe.Pointer(&srcData[0]))
	cSrcStride := (*C.int)(unsafe.Pointer(&srcStride[0]))
	cDst := (**C.uint8_t)(unsafe.Pointer(&dstData[0]))
	cDstStride := (*C.int)(unsafe.Pointer(&dstStride))
	return int(C.sws_scale(cCtx, cSrc, cSrcStride, C.int(y), C.int(h), cDst, cDstStride))
}

func SwsScale3(ctxt *Context, fSrc *avutil.Frame, y, h int, fDst *avutil.Frame) int {
	//srcData [8]*uint8, srcStride [8]int32
	//dstData [8]*uint8, dstStride [8]int32
	//csrc := (**C.uint8_t)(unsafe.Pointer(&srcData[0]))
	//cstr := (*C.int)(unsafe.Pointer(&srcStride[0]))
	//cd := (**C.uint8_t)(unsafe.Pointer(&dstData[0]))
	//cds := (*C.int)(unsafe.Pointer(&dstStride))
	cSwsCtx := (*C.struct_SwsContext)(unsafe.Pointer(ctxt))
	cfSrc := (*C.struct_AVFrame)(unsafe.Pointer(fSrc))
	cfDst := (*C.struct_AVFrame)(unsafe.Pointer(fDst))
	cdSrc := (**C.uint8_t)(&cfSrc.data[0])
	cdLineSizeSrc := (*C.int)(&cfSrc.linesize[0])
	cdDst := (**C.uint8_t)(&cfDst.data[0])
	cdLineSizeDst := (*C.int)(&cfDst.linesize[0])
	return int(C.sws_scale(cSwsCtx, cdSrc, cdLineSizeSrc, C.int(y), C.int(h), cdDst, cdLineSizeDst))
}

func SwsSetcolorspacedetails(ctxt *Context, it *int, sr int, t *int, dr, b, c, s int) int {
	cit := (*C.int)(unsafe.Pointer(it))
	ct := (*C.int)(unsafe.Pointer(t))
	return int(C.sws_setColorspaceDetails((*C.struct_SwsContext)(ctxt), cit, C.int(sr), ct, C.int(dr), C.int(b), C.int(c), C.int(s)))
}

func SwsGetcolorspacedetails(ctxt *Context, it, sr, t, dr, b, c, s *int) int {
	cit := (**C.int)(unsafe.Pointer(it))
	csr := (*C.int)(unsafe.Pointer(sr))
	ct := (**C.int)(unsafe.Pointer(t))
	cdr := (*C.int)(unsafe.Pointer(dr))
	cb := (*C.int)(unsafe.Pointer(b))
	cc := (*C.int)(unsafe.Pointer(c))
	cs := (*C.int)(unsafe.Pointer(s))
	return int(C.sws_getColorspaceDetails((*C.struct_SwsContext)(ctxt), cit, csr, ct, cdr, cb, cc, cs))
}

func SwsGetdefaultfilter(lb, cb, ls, cs, chs, cvs float32, v int) *Filter {
	return (*Filter)(unsafe.Pointer(C.sws_getDefaultFilter(C.float(lb), C.float(cb), C.float(ls), C.float(cs), C.float(chs), C.float(cvs), C.int(v))))
}

func SwsFreefilter(f *Filter) {
	C.sws_freeFilter((*C.struct_SwsFilter)(f))
}

//Convert an 8-bit paletted frame into a frame with a color depth of 32 bits.
func SwsConvertpalette8topacked32(s, d *uint8, px int, p *uint8) {
	C.sws_convertPalette8ToPacked32((*C.uint8_t)(s), (*C.uint8_t)(d), C.int(px), (*C.uint8_t)(p))
}

//Convert an 8-bit paletted frame into a frame with a color depth of 24 bits.
func SwsConvertpalette8topacked24(s, d *uint8, px int, p *uint8) {
	C.sws_convertPalette8ToPacked24((*C.uint8_t)(s), (*C.uint8_t)(d), C.int(px), (*C.uint8_t)(p))
}

//Get the Class for swsContext.
func SwsGetClass() *Class {
	return (*Class)(C.sws_get_class())
}
