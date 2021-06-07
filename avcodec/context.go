// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

//#cgo pkg-config: libavcodec libavutil
//#include <libavcodec/avcodec.h>
import "C"
import (
	"github.com/aeotheme/goav/avutil"
	"unsafe"
)

//Free the codec context and everything associated with it and write NULL to the provided pointer.
func (ctx *Context) AvcodecFreeContext() {
	C.avcodec_free_context((**C.struct_AVCodecContext)(unsafe.Pointer(&ctx)))
}

//Set the fields of the given Context to default values corresponding to the given codec (defaults may be codec-dependent).
func (ctx *Context) AvcodecGetContextDefaults3(c *Codec) int {
	return int(C.avcodec_get_context_defaults3((*C.struct_AVCodecContext)(ctx), (*C.struct_AVCodec)(c)))
}

//Initialize the Context to use the given Codec
func (ctx *Context) AvcodecOpen2(c *Codec, d **Dictionary) int {
	return int(C.avcodec_open2((*C.struct_AVCodecContext)(ctx), (*C.struct_AVCodec)(c), (**C.struct_AVDictionary)(unsafe.Pointer(d))))
}

//Close a given Context and free all the data associated with it (but not the Context itself).
func (ctx *Context) AvcodecClose() int {
	return int(C.avcodec_close((*C.struct_AVCodecContext)(ctx)))
}

//The default callback for Context.get_buffer2().
func (ctx *Context) AvcodecDefaultGetBuffer2(f *Frame, l int) int {
	return int(C.avcodec_default_get_buffer2((*C.struct_AVCodecContext)(ctx), (*C.struct_AVFrame)(f), C.int(l)))
}

//Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you do not use any horizontal padding.
func (ctx *Context) AvcodecAlignDimensions(w, h *int) {
	C.avcodec_align_dimensions((*C.struct_AVCodecContext)(ctx), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)))
}

//Modify width and height values so that they will result in a memory buffer that is acceptable for the codec if you also ensure that all line sizes are a multiple of the respective linesize_align[i].
func (ctx *Context) AvcodecAlignDimensions2(w, h *int, l int) {
	C.avcodec_align_dimensions2((*C.struct_AVCodecContext)(ctx), (*C.int)(unsafe.Pointer(w)), (*C.int)(unsafe.Pointer(h)), (*C.int)(unsafe.Pointer(&l)))
}

//Decode a subtitle message.
func (ctx *Context) AvcodecDecodeSubtitle2(s *AvSubtitle, g *int, a *Packet) int {
	return int(C.avcodec_decode_subtitle2((*C.struct_AVCodecContext)(ctx), (*C.struct_AVSubtitle)(s), (*C.int)(unsafe.Pointer(g)), (*C.struct_AVPacket)(a)))
}

func (ctx *Context) AvcodecEncodeSubtitle(b *uint8, bs int, s *AvSubtitle) int {
	return int(C.avcodec_encode_subtitle((*C.struct_AVCodecContext)(ctx), (*C.uint8_t)(b), C.int(bs), (*C.struct_AVSubtitle)(s)))
}

func (ctx *Context) AvcodecDefaultGetFormat(f *PixelFormat) PixelFormat {
	return (PixelFormat)(C.avcodec_default_get_format((*C.struct_AVCodecContext)(ctx), (*C.enum_AVPixelFormat)(f)))
}

//Reset the internal decoder state / flush internal buffers.
func (ctx *Context) AvcodecFlushBuffers() {
	C.avcodec_flush_buffers((*C.struct_AVCodecContext)(ctx))
}

//Return audio frame duration.
func (ctx *Context) AvGetAudioFrameDuration(f int) int {
	return int(C.av_get_audio_frame_duration((*C.struct_AVCodecContext)(ctx), C.int(f)))
}

func (ctx *Context) AvcodecIsOpen() int {
	return int(C.avcodec_is_open((*C.struct_AVCodecContext)(ctx)))
}

//Parse a packet.
func (ctx *Context) AvParserParse2(ctxtp *ParserContext, p **uint8, ps *int, b *uint8, bs int, pt, dt, po int64) int {
	return int(C.av_parser_parse2((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(ctx), (**C.uint8_t)(unsafe.Pointer(p)), (*C.int)(unsafe.Pointer(ps)), (*C.uint8_t)(b), C.int(bs), (C.int64_t)(pt), (C.int64_t)(dt), (C.int64_t)(po)))
}

func (ctx *Context) AvParserChange(ctxtp *ParserContext, pb **uint8, pbs *int, b *uint8, bs, k int) int {
	return int(C.av_parser_change((*C.struct_AVCodecParserContext)(ctxtp), (*C.struct_AVCodecContext)(ctx), (**C.uint8_t)(unsafe.Pointer(pb)), (*C.int)(unsafe.Pointer(pbs)), (*C.uint8_t)(b), C.int(bs), C.int(k)))
}

func AvParserInit(c int) *ParserContext {
	return (*ParserContext)(C.av_parser_init(C.int(c)))
}

func AvParserClose(ctxtp *ParserContext) {
	C.av_parser_close((*C.struct_AVCodecParserContext)(ctxtp))
}

func (ctx *Context) SetTimebase(num1 int, den1 int) {
	ctx.time_base.num = C.int(num1)
	ctx.time_base.den = C.int(den1)
}

func (ctx *Context) SetEncodeParams2(width int, height int, pxlFmt PixelFormat, hasBframes bool, gopSize int) {
	ctx.width = C.int(width)
	ctx.height = C.int(height)
	// ctx.bit_rate = 1000000
	ctx.gop_size = C.int(gopSize)
	// ctx.max_b_frames = 2
	if hasBframes {
		ctx.has_b_frames = 1
	} else {
		ctx.has_b_frames = 0
	}
	// ctx.extradata = nil
	// ctx.extradata_size = 0
	// ctx.channels = 0
	ctx.pix_fmt = int32(pxlFmt)
	// C.av_opt_set(ctx.priv_data, "preset", "ultrafast", 0)
}

func (ctx *Context) SetEncodeParams(width int, height int, pxlFmt PixelFormat) {
	ctx.SetEncodeParams2(width, height, pxlFmt, false /*no b frames*/, 10)
}

func (ctx *Context) AvcodecSendPacket(packet *Packet) int {
	return (int)(C.avcodec_send_packet((*C.struct_AVCodecContext)(ctx), (*C.struct_AVPacket)(packet)))
}

func (ctx *Context) AvcodecReceiveFrame(frame *avutil.Frame) int {
	frame0 := (*Frame)(unsafe.Pointer(frame))
	return (int)(C.avcodec_receive_frame((*C.struct_AVCodecContext)(ctx), (*C.struct_AVFrame)(frame0)))
}
