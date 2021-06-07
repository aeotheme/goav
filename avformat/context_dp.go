//+build DoDeprecated

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import "unsafe"

// AvFormatGetProbeScore deprecated
func (s *Context) AvFormatGetProbeScore() int {
	return int(C.av_format_get_probe_score((*C.struct_AVFormatContext)(s)))
}

// AvFormatGetVideoCodec deprecated
func (s *Context) AvFormatGetVideoCodec() *AvCodec {
	return (*AvCodec)(C.av_format_get_video_codec((*C.struct_AVFormatContext)(s)))
}

// AvFormatSetVideoCodec deprecated
func (s *Context) AvFormatSetVideoCodec(c *AvCodec) {
	C.av_format_set_video_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
}

// AvFormatGetAudioCodec deprecated
func (s *Context) AvFormatGetAudioCodec() *AvCodec {
	return (*AvCodec)(C.av_format_get_audio_codec((*C.struct_AVFormatContext)(s)))
}

// AvFormatSetAudioCodec deprecated
func (s *Context) AvFormatSetAudioCodec(c *AvCodec) {
	C.av_format_set_audio_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
}

// AvFormatGetSubtitleCodec deprecated
func (s *Context) AvFormatGetSubtitleCodec() *AvCodec {
	return (*AvCodec)(C.av_format_get_subtitle_codec((*C.struct_AVFormatContext)(s)))
}

// AvFormatSetSubtitleCodec deprecated
func (s *Context) AvFormatSetSubtitleCodec(c *AvCodec) {
	C.av_format_set_subtitle_codec((*C.struct_AVFormatContext)(s), (*C.struct_AVCodec)(c))
}

// AvFormatGetMetadataHeaderPadding deprecated
func (s *Context) AvFormatGetMetadataHeaderPadding() int {
	return int(C.av_format_get_metadata_header_padding((*C.struct_AVFormatContext)(s)))
}

// AvFormatSetMetadataHeaderPadding deprecated
func (s *Context) AvFormatSetMetadataHeaderPadding(c int) {
	C.av_format_set_metadata_header_padding((*C.struct_AVFormatContext)(s), C.int(c))
}

// AvFormatGetOpaque deprecated
func (s *Context) AvFormatGetOpaque() {
	C.av_format_get_opaque((*C.struct_AVFormatContext)(s))
}

// AvFormatSetOpaque deprecated
func (s *Context) AvFormatSetOpaque(o int) {
	C.av_format_set_opaque((*C.struct_AVFormatContext)(s), unsafe.Pointer(&o))
}
