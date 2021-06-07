//+build DoDeprecated

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"github.com/aeotheme/goav/avcodec"
	"github.com/aeotheme/goav/avutil"
	"unsafe"
)

// AvStreamGetRFrameRate deprecated
//Rational av_stream_get_r_frame_rate (const Stream *s)
func (avs *Stream) AvStreamGetRFrameRate() avutil.Rational {
	return newRational(C.av_stream_get_r_frame_rate((*C.struct_AVStream)(avs)))
}

// AvStreamSetRFrameRate deprecated
//void av_stream_set_r_frame_rate (Stream *s, Rational r)
func (avs *Stream) AvStreamSetRFrameRate(r avutil.Rational) {
	rat := C.struct_AVRational{
		num: C.int(r.Num()),
		den: C.int(r.Den()),
	}

	C.av_stream_set_r_frame_rate((*C.struct_AVStream)(avs), rat)
}

func (avs *Stream) Codec() *avcodec.Context {
	return (*C.struct_AVCodecContext)(unsafe.Pointer(avs.codec))
}

func (avs *Stream) RecommendedEncoderConfiguration() string {
	return C.GoString(avs.recommended_encoder_configuration)
}

//char * av_stream_get_recommended_encoder_configuration (const Stream *s)
func (avs *Stream) AvStreamGetRecommendedEncoderConfiguration() string {
	return C.GoString(C.av_stream_get_recommended_encoder_configuration((*C.struct_AVStream)(avs)))
}

//void av_stream_set_recommended_encoder_configuration (Stream *s, char *configuration)
func (avs *Stream) AvStreamSetRecommendedEncoderConfiguration(c string) {
	C.av_stream_set_recommended_encoder_configuration((*C.struct_AVStream)(avs), C.CString(c))
}

// func (avs *Stream) PrivPts() *FFFrac {
// 	return (*FFFrac)(unsafe.Pointer(avs.priv_pts))
// }
