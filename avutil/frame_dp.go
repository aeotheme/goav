//+build DoDeprecated

package avutil

//#cgo pkg-config: libavutil
//#include <libavutil/frame.h>
//#include <stdlib.h>
import "C"
import (
	"unsafe"
)

// AvFrameSetQpTable deprecated
func AvFrameSetQpTable(f *Frame, b *AvBufferRef, s, q int) int {
	return int(C.av_frame_set_qp_table((*C.struct_AVFrame)(unsafe.Pointer(f)), (*C.struct_AVBufferRef)(unsafe.Pointer(b)), C.int(s), C.int(q)))
}

// AvFrameGetQpTable deprecated
func AvFrameGetQpTable(f *Frame, s, t *int) int8 {
	return int8(*C.av_frame_get_qp_table((*C.struct_AVFrame)(unsafe.Pointer(f)), (*C.int)(unsafe.Pointer(s)), (*C.int)(unsafe.Pointer(t))))
}
