package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/imgutils.h>
import "C"
import (
	"unsafe"
)

// AvPictureGetBufferSize
//
//int av_image_get_buffer_size(enum AVPixelFormat pix_fmt, int width, int height, int align);
func AvPictureGetBufferSize(pf PixelFormat, w, h, align int) int {
	return int(C.av_image_get_buffer_size((C.enum_AVPixelFormat)(pf), C.int(w), C.int(h), C.int(align)))
}

// int av_image_fill_arrays(uint8_t *dst_data[4], int dst_linesize[4],
//   const uint8_t *src,enum AVPixelFormat pix_fmt, int width, int height, int align);
func AvPictureFill(dst uintptr, dstLineSize uintptr, src *uint8, pf PixelFormat, w, h, align int) int {
	//for k := range d1 {
	//	d1[k] = (*C.uint8_t)(dst[k])
	//	d2[k] = (C.int)(dstLineSize[k])
	//}
	//
	//cDst := (**C.uint8_t)(unsafe.Pointer(&d1[0]))
	//cDstLineSize := (*C.int)(unsafe.Pointer(&d2[0]))
	cDst := (**C.uint8_t)(unsafe.Pointer(dst))
	cDstLineSize := (*C.int)(unsafe.Pointer(dstLineSize))
	cSrc := (*C.uint8_t)(unsafe.Pointer(src))

	return int(C.av_image_fill_arrays(cDst, cDstLineSize, cSrc,
		(C.enum_AVPixelFormat)(pf), C.int(w), C.int(h), C.int(align)))
}

// int av_image_fill_arrays(uint8_t *dst_data[4], int dst_linesize[4],
//   const uint8_t *src,enum AVPixelFormat pix_fmt, int width, int height, int align);
func AvPictureFill2(f *Frame, src *uint8, pf PixelFormat, w, h, align int) int {
	//var dst unsafe.Pointer
	//var dstLineSize unsafe.Pointer
	cF := (*C.struct_AVFrame)(unsafe.Pointer(f))
	cDst := (**C.uint8_t)(&cF.data[0])
	cDstLineSize := (*C.int)(&cF.linesize[0])
	cSrc := (*C.uint8_t)(unsafe.Pointer(src))

	return int(C.av_image_fill_arrays(cDst, cDstLineSize, cSrc,
		(C.enum_AVPixelFormat)(pf), C.int(w), C.int(h), C.int(align)))
}
