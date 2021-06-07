//+build DoDeprecated

package avcodec

//#cgo pkg-config: libavformat libavcodec libavutil libswresample
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
import "C"

// AvCodecGetMaxLowRes deprecated
func (c *Codec) AvCodecGetMaxLowRes() int {
	return int(C.av_codec_get_max_lowres((*C.struct_AVCodec)(c)))
}

// AvCodecNext deprecated
// AvCodecNext If c is NULL, returns the first registered codec, if c is non-NULL,
func (c *Codec) AvCodecNext() *Codec {
	return (*Codec)(C.av_codec_next((*C.struct_AVCodec)(c)))
}

// AvcodecRegister deprecated
// Register the codec codec and initialize libavcodec.
func (c *Codec) AvcodecRegister() {
	C.avcodec_register((*C.struct_AVCodec)(c))
}

// AvcodecRegisterAll deprecated
//Register all the codecs, parsers and bitstream filters which were enabled at configuration time.
func AvcodecRegisterAll() {
	C.av_register_all()
	C.avcodec_register_all()
	// C.av_log_set_level(0xffff)
}

// AvGetCodecTagString deprecated
//Put a string representing the codec tag codec_tag in buf.
func AvGetCodecTagString(b string, bf uintptr, c uint) uintptr {
	return uintptr(C.av_get_codec_tag_string(C.CString(b), C.size_t(bf), C.uint(c)))
}

// AvHwaccelNext deprecated
//If hwaccel is NULL, returns the first registered hardware accelerator, if hwaccel is non-NULL,
//returns the next registered hardware accelerator after hwaccel, or NULL if hwaccel is the last one.
func (a *AvHWAccel) AvHwaccelNext() *AvHWAccel {
	return (*AvHWAccel)(C.av_hwaccel_next((*C.struct_AVHWAccel)(a)))
}
