//+build DoDeprecated

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"unsafe"
)

// AvcodecGetChromaSubSample deprecated
//
//Utility function to access log2_chroma_w log2_chroma_h from the pixel format AvPixFmtDescriptor.
func (p PixelFormat) AvcodecGetChromaSubSample(h, v *int) {
	C.avcodec_get_chroma_sub_sample((C.enum_AVPixelFormat)(p), (*C.int)(unsafe.Pointer(h)), (*C.int)(unsafe.Pointer(v)))
}
