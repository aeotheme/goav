//+build DoDeprecated

package swscale

//#cgo pkg-config: libswscale libavutil
//#include <libswscale/swscale.h>
import "C"
import (
	"unsafe"
)

// SwsConvVec deprecated
func (a *Vector) SwsConvVec(b *Vector) {
	C.sws_convVec((*C.struct_SwsVector)(a), (*C.struct_SwsVector)(b))
}

// SwsAddVec deprecated
func (a *Vector) SwsAddVec(b *Vector) {
	C.sws_addVec((*C.struct_SwsVector)(a), (*C.struct_SwsVector)(b))
}

// SwsSubVec deprecated
func (a *Vector) SwsSubVec(b *Vector) {
	C.sws_subVec((*C.struct_SwsVector)(a), (*C.struct_SwsVector)(b))
}

// SwsShiftVec deprecated
func (a *Vector) SwsShiftVec(s int) {
	C.sws_shiftVec((*C.struct_SwsVector)(a), C.int(s))
}

// SwsCloneVec deprecated
//Allocate and return a clone of the vector a, that is a vector with the same coefficients as a.
func (a *Vector) SwsCloneVec() *Vector {
	return (*Vector)(unsafe.Pointer(C.sws_cloneVec((*C.struct_SwsVector)(a))))
}

// SwsPrintVec2 deprecated
//Print with av_log() a textual representation of the vector a if log_level <= av_log_level.
func (a *Vector) SwsPrintVec2(lctx *Class, l int) {
	C.sws_printVec2((*C.struct_SwsVector)(a), (*C.struct_AVClass)(lctx), C.int(l))
}

// SwsGetConstVec deprecated
//Allocate and return a vector with length coefficients, all with the same value c.
func SwsGetConstVec(c float64, l int) *Vector {
	return (*Vector)(unsafe.Pointer(C.sws_getConstVec(C.double(c), C.int(l))))
}

// SwsGetIdentityVec is deprecated
//Allocate and return a vector with just one coefficient, with value 1.0.
func SwsGetIdentityVec() *Vector {
	return (*Vector)(unsafe.Pointer(C.sws_getIdentityVec()))
}
