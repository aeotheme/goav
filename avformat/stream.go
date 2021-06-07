// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avformat

//#cgo pkg-config: libavformat
//#include <libavformat/avformat.h>
import "C"
import (
	"github.com/aeotheme/goav/avcodec"
	"github.com/aeotheme/goav/avutil"
	"unsafe"
)

//struct CodecParserContext * av_stream_get_parser (const Stream *s)
func (avs *Stream) AvStreamGetParser() *CodecParserContext {
	return (*CodecParserContext)(C.av_stream_get_parser((*C.struct_AVStream)(avs)))
}

//int64_t av_stream_get_end_pts (const Stream *st)
//Returns the pts of the last muxed packet + its duration.
func (avs *Stream) AvStreamGetEndPts() int64 {
	return int64(C.av_stream_get_end_pts((*C.struct_AVStream)(avs)))
}

//Get side information from stream.
func (avs *Stream) AvStreamGetSideData(t AvPacketSideDataType, z int) *uint8 {
	return (*uint8)(C.av_stream_get_side_data((*C.struct_AVStream)(avs), (C.enum_AVPacketSideDataType)(t), (*C.int)(unsafe.Pointer(&z))))
}

func (avs *Stream) Free() {
	C.av_freep(unsafe.Pointer(avs))
}

func (avs *Stream) RFrameRate() avutil.Rational {
	return newRational(avs.r_frame_rate)
}

func (avs *Stream) SampleAspectRatio() avutil.Rational {
	return newRational(avs.sample_aspect_ratio)
}

func (avs *Stream) TimeBase() avutil.Rational {
	return newRational(avs.time_base)
}

func (avs *Stream) AvgFrameRate() avutil.Rational {
	return newRational(avs.avg_frame_rate)
}

func (avs *Stream) AttachedPic() avcodec.Packet {
	return *fromCPacket(&avs.attached_pic)
	//return *(*avcodec.Packet)(unsafe.Pointer(&avs.attached_pic))
}
