//+build DoDeprecated

package avcodec

//#cgo pkg-config: libavcodec
//#include <libavcodec/avcodec.h>
import "C"

// AvDupPacket deprecated
func (p *Packet) AvDupPacket() int {
	return int(C.av_dup_packet((*C.struct_AVPacket)(p)))
}

// AvCopyPacket deprecated
//
//Copy packet, including contents.
func (p *Packet) AvCopyPacket(r *Packet) int {
	return int(C.av_copy_packet((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(r)))
}

// AvCopyPacketSideData deprecated
//
//Copy packet side data.
func (p *Packet) AvCopyPacketSideData(r *Packet) int {
	return int(C.av_copy_packet_side_data((*C.struct_AVPacket)(p), (*C.struct_AVPacket)(r)))
}

// AvFreePacket deprecated
//
//Free a packet.
func (p *Packet) AvFreePacket() {
	C.av_free_packet((*C.struct_AVPacket)(p))
}

// AvPacketMergeSideData deprecated
//int 	av_packet_merge_side_data (Packet *pkt)
func (p *Packet) AvPacketMergeSideData() int {
	return int(C.av_packet_merge_side_data((*C.struct_AVPacket)(p)))
}

// AvPacketSplitSideData deprecated
//int 	av_packet_split_side_data (Packet *pkt)
func (p *Packet) AvPacketSplitSideData() int {
	return int(C.av_packet_split_side_data((*C.struct_AVPacket)(p)))
}
