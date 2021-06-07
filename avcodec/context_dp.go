//+build DoDeprecated

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"
import (
	"github.com/aeotheme/goav/avutil"
	"unsafe"
)

// deprecated
// AvCodecGetPktTimebase2
// AvCodecGetPktTimebase2 returns the timebase rational number as numerator and denominator
func (ctx *Context) AvCodecGetPktTimebase2() avutil.Rational {
	return ctx.AvCodecGetPktTimebase()
}

// AvCodecGetPktTimebase deprecated
func (ctx *Context) AvCodecGetPktTimebase() avutil.Rational {
	return C.av_codec_get_pkt_timebase((*C.struct_AVCodecContext)(ctx))
}

// AvCodecSetPktTimebase deprecated
func (ctx *Context) AvCodecSetPktTimebase(r avutil.Rational) {
	C.av_codec_set_pkt_timebase((*C.struct_AVCodecContext)(ctx), (C.struct_AVRational)(r))
}

// AvCodecGetCodecDescriptor deprecated
func (ctx *Context) AvCodecGetCodecDescriptor() *Descriptor {
	return (*Descriptor)(C.av_codec_get_codec_descriptor((*C.struct_AVCodecContext)(ctx)))
}

// AvCodecSetCodecDescriptor deprecated
func (ctx *Context) AvCodecSetCodecDescriptor(d *Descriptor) {
	C.av_codec_set_codec_descriptor((*C.struct_AVCodecContext)(ctx), (*C.struct_AVCodecDescriptor)(d))
}

// AvCodecGetLowres deprecated
func (ctx *Context) AvCodecGetLowres() int {
	return int(C.av_codec_get_lowres((*C.struct_AVCodecContext)(ctx)))
}

// AvCodecSetLowres deprecated
func (ctx *Context) AvCodecSetLowres(i int) {
	C.av_codec_set_lowres((*C.struct_AVCodecContext)(ctx), C.int(i))
}

// AvCodecGetSeekPreroll deprecated
func (ctx *Context) AvCodecGetSeekPreroll() int {
	return int(C.av_codec_get_seek_preroll((*C.struct_AVCodecContext)(ctx)))
}

// AvCodecSetSeekPreroll deprecated
func (ctx *Context) AvCodecSetSeekPreroll(i int) {
	C.av_codec_set_seek_preroll((*C.struct_AVCodecContext)(ctx), C.int(i))
}

// AvCodecGetChromaIntraMatrix deprecated
func (ctx *Context) AvCodecGetChromaIntraMatrix() *uint16 {
	return (*uint16)(C.av_codec_get_chroma_intra_matrix((*C.struct_AVCodecContext)(ctx)))
}

// AvCodecSetChromaIntraMatrix deprecated
func (ctx *Context) AvCodecSetChromaIntraMatrix(t *uint16) {
	C.av_codec_set_chroma_intra_matrix((*C.struct_AVCodecContext)(ctx), (*C.uint16_t)(t))
}

// AvcodecCopyContext deprecated
//Copy the settings of the source Context into the destination Context.
func (ctx *Context) AvcodecCopyContext(ctxt2 *Context) int {
	return int(C.avcodec_copy_context((*C.struct_AVCodecContext)(ctx), (*C.struct_AVCodecContext)(ctxt2)))
}

// AvcodecDecodeAudio4 deprecated
//
//Decode the audio frame of size avpkt->size from avpkt->data into frame.
func (ctx *Context) AvcodecDecodeAudio4(f *Frame, g *int, a *Packet) int {
	return int(C.avcodec_decode_audio4((*C.struct_AVCodecContext)(ctx), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

// AvcodecDecodeVideo2 deprecated
//
//Decode the video frame of size avpkt->size from avpkt->data into picture.
func (ctx *Context) AvcodecDecodeVideo2(p *Frame, g *int, a *Packet) int {
	return int(C.avcodec_decode_video2((*C.struct_AVCodecContext)(ctx), (*C.struct_AVFrame)(p), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

// AvcodecEncodeAudio2 deprecated
//
//Encode a frame of audio.
func (ctx *Context) AvcodecEncodeAudio2(p *Packet, f *Frame, gp *int) int {
	return int(C.avcodec_encode_audio2((*C.struct_AVCodecContext)(ctx), (*C.struct_AVPacket)(p), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(gp))))
}

// AvcodecEncodeVideo2 deprecated
//
//Encode a frame of video
func (ctx *Context) AvcodecEncodeVideo2(p *Packet, f *Frame, gp *int) int {
	return int(C.avcodec_encode_video2((*C.struct_AVCodecContext)(ctx), (*C.struct_AVPacket)(p), (*C.struct_AVFrame)(f), (*C.int)(unsafe.Pointer(gp))))
}

// AvParserNext deprecated
func (p *Parser) AvParserNext() *Parser {
	return (*Parser)(C.av_parser_next((*C.struct_AVCodecParser)(p)))
}

// AvRegisterCodecParser deprecated
func (p *Parser) AvRegisterCodecParser() {
	C.av_register_codec_parser((*C.struct_AVCodecParser)(p))
}
