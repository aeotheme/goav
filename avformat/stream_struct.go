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

func (avs *Stream) CodecParameters() *avcodec.Parameters {
	cS := (*C.struct_AVStream)(unsafe.Pointer(avs))
	return (*avcodec.Parameters)(unsafe.Pointer(cS.codecpar))
}

func (avs *Stream) Metadata() *Dictionary {
	cS := (*C.struct_AVStream)(unsafe.Pointer(avs))
	return (*Dictionary)(unsafe.Pointer(cS.metadata))
}

func (avs *Stream) IndexEntries() *AvIndexEntry {
	cS := (*C.struct_AVStream)(unsafe.Pointer(avs))
	return (*AvIndexEntry)(unsafe.Pointer(cS.index_entries))
}

func (avs *Stream) SideData() *AvPacketSideData {
	cS := (*C.struct_AVStream)(unsafe.Pointer(avs))
	return (*AvPacketSideData)(unsafe.Pointer(cS.side_data))
}

func (avs *Stream) ProbeData() AvProbeData {
	cS := (*C.struct_AVStream)(unsafe.Pointer(avs))
	return *(*AvProbeData)(unsafe.Pointer(&cS.probe_data))
}

func (avs *Stream) DisplayAspectRatio() (r avutil.Rational) {
	cS := (*C.struct_AVStream)(unsafe.Pointer(avs))
	return *(*avutil.Rational)(unsafe.Pointer(&cS.display_aspect_ratio))
}

func (avs *Stream) StreamInternal() *AVStreamInternal {
	cS := (*C.struct_AVStream)(unsafe.Pointer(avs))
	return (*AVStreamInternal)(unsafe.Pointer(cS.internal))
}

func (avs *Stream) Discard() AvDiscard {
	cS := (*C.struct_AVStream)(unsafe.Pointer(avs))
	return (AvDiscard)(cS.discard)
}

func (avs *Stream) NeedParsing() AvStreamParseType {
	cS := (*C.struct_AVStream)(unsafe.Pointer(avs))
	return (AvStreamParseType)(cS.need_parsing)
}

func (avs *Stream) CodecInfoNbFrames() int {
	return int(avs.codec_info_nb_frames)
}

func (avs *Stream) Disposition() int {
	return int(avs.disposition)
}

func (avs *Stream) EventFlags() int {
	return int(avs.event_flags)
}

func (avs *Stream) PtsWrapBits() int {
	return int(avs.pts_wrap_bits)
}

func (avs *Stream) Id() int {
	return int(avs.id)
}

func (avs *Stream) Index() int {
	return int(avs.index)
}

func (avs *Stream) InjectGlobalSideData() int {
	return int(avs.inject_global_side_data)
}

func (avs *Stream) LastIpDuration() int {
	return int(avs.last_IP_duration)
}

func (avs *Stream) NbDecodedFrames() int {
	return int(avs.nb_decoded_frames)
}

func (avs *Stream) NbIndexEntries() int {
	return int(avs.nb_index_entries)
}

func (avs *Stream) NbSideData() int {
	return int(avs.nb_side_data)
}

func (avs *Stream) ProbePackets() int {
	return int(avs.probe_packets)
}

func (avs *Stream) PtsWrapBehavior() int {
	return int(avs.pts_wrap_behavior)
}

func (avs *Stream) RequestProbe() int {
	return int(avs.request_probe)
}

func (avs *Stream) SkipSamples() int {
	return int(avs.skip_samples)
}

func (avs *Stream) SkipToKeyframe() int {
	return int(avs.skip_to_keyframe)
}

func (avs *Stream) StreamIdentifier() int {
	return int(avs.stream_identifier)
}

func (avs *Stream) ProgramNum() int {
	return int(avs.program_num)
}

func (avs *Stream) PmtVersion() int {
	return int(avs.pmt_version)
}

func (avs *Stream) PmtStreamIndex() int {
	return int(avs.pmt_stream_idx)
}

func (avs *Stream) UpdateInitialDurationsDone() int {
	return int(avs.update_initial_durations_done)
}

func (avs *Stream) CurDts() int64 {
	return int64(avs.cur_dts)
}

func (avs *Stream) Duration() int64 {
	return int64(avs.duration)
}

func (avs *Stream) FirstDiscardSample() int64 {
	return int64(avs.first_discard_sample)
}

func (avs *Stream) FirstDts() int64 {
	return int64(avs.first_dts)
}

func (avs *Stream) InterleaverChunkDuration() int64 {
	return int64(avs.interleaver_chunk_duration)
}

func (avs *Stream) InterleaverChunkSize() int64 {
	return int64(avs.interleaver_chunk_size)
}

func (avs *Stream) LastDiscardSample() int64 {
	return int64(avs.last_discard_sample)
}

func (avs *Stream) LastDtsForOrderCheck() int64 {
	return int64(avs.last_dts_for_order_check)
}

func (avs *Stream) LastIpPts() int64 {
	return int64(avs.last_IP_pts)
}

func (avs *Stream) MuxTsOffset() int64 {
	return int64(avs.mux_ts_offset)
}

func (avs *Stream) NbFrames() int64 {
	return int64(avs.nb_frames)
}

func (avs *Stream) PtsBuffer() int64 {
	return int64(avs.pts_buffer[0])
}

func (avs *Stream) PtsReorderError() int64 {
	return int64(avs.pts_reorder_error[0])
}

func (avs *Stream) PtsWrapReference() int64 {
	return int64(avs.pts_wrap_reference)
}

func (avs *Stream) StartSkipSamples() int64 {
	return int64(avs.start_skip_samples)
}

func (avs *Stream) StartTime() int64 {
	return int64(avs.start_time)
}

func (avs *Stream) Parser() *CodecParserContext {
	cpCtx := (*C.struct_AVCodecParserContext)(unsafe.Pointer(avs.parser))
	return (*CodecParserContext)(cpCtx)
}

func (avs *Stream) LastInPacketBuffer() *AvPacketList {
	pl := (*C.struct_AVPacketList)(unsafe.Pointer(avs.last_in_packet_buffer))
	return (*AvPacketList)(pl)
}

func (avs *Stream) DtsMisOrdered() uint8 {
	return uint8(avs.dts_misordered)
}

func (avs *Stream) DtsOrdered() uint8 {
	return uint8(avs.dts_ordered)
}

func (avs *Stream) PtsReorderErrorCount() uint8 {
	return uint8(avs.pts_reorder_error_count[0])
}

func (avs *Stream) IndexEntriesAllocatedSize() uint {
	return uint(avs.index_entries_allocated_size)
}

func (avs *Stream) PrivateData() uintptr {
	return uintptr(unsafe.Pointer(avs.priv_data))
}
