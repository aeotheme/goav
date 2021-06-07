//+build DoDeprecated

package avformat

//#cgo pkg-config: libavformat libavcodec libavutil libavdevice libavfilter libswresample libswscale
//#include <stdio.h>
//#include <stdlib.h>
//#include <inttypes.h>
//#include <stdint.h>
//#include <string.h>
//#include <libavformat/avformat.h>
//#include <libavcodec/avcodec.h>
//#include <libavutil/avutil.h>
//#include <libavutil/opt.h>
//#include <libavdevice/avdevice.h>
import "C"

// AvRegisterInputFormat deprecated
func (f *InputFormat) AvRegisterInputFormat() {
	C.av_register_input_format((*C.struct_AVInputFormat)(f))
}

// AvRegisterOutputFormat deprecated
func (f *OutputFormat) AvRegisterOutputFormat() {
	C.av_register_output_format((*C.struct_AVOutputFormat)(f))
}

// AvIFormatNext deprecated
//
//If f is NULL, returns the first registered input format, if f is non-NULL, returns the next registered input format after f or NULL if f is the last one.
func (f *InputFormat) AvIFormatNext() *InputFormat {
	return (*InputFormat)(C.av_iformat_next((*C.struct_AVInputFormat)(f)))
}

// AvOFormatNext deprecated
//
//If f is NULL, returns the first registered output format, if f is non-NULL, returns the next registered output format after f or NULL if f is the last one.
func (f *OutputFormat) AvOFormatNext() *OutputFormat {
	return (*OutputFormat)(C.av_oformat_next((*C.struct_AVOutputFormat)(f)))
}

// AvRegisterAll deprecated
//Initialize libavformat and register all the muxers, demuxers and protocols.
func AvRegisterAll() {
	C.av_register_all()
}
