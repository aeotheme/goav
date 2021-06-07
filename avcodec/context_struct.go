// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
// Giorgis (habtom@giorgis.io)

package avcodec

import "github.com/aeotheme/goav/swscale"

func (ctx *Context) ActiveThreadType() int {
	return int(ctx.active_thread_type)
}

func (ctx *Context) BFrameStrategy() int {
	return int(ctx.b_frame_strategy)
}

func (ctx *Context) BQuantFactor() float64 {
	return float64(ctx.b_quant_factor)
}

func (ctx *Context) BQuantOffset() float64 {
	return float64(ctx.b_quant_offset)
}

func (ctx *Context) BSensitivity() int {
	return int(ctx.b_sensitivity)
}

func (ctx *Context) BidirRefine() int {
	return int(ctx.bidir_refine)
}

func (ctx *Context) BitRate() int {
	return int(ctx.bit_rate)
}

func (ctx *Context) BitRateTolerance() int {
	return int(ctx.bit_rate_tolerance)
}

func (ctx *Context) BitsPerCodedSample() int {
	return int(ctx.bits_per_coded_sample)
}

func (ctx *Context) BitsPerRawSample() int {
	return int(ctx.bits_per_raw_sample)
}

func (ctx *Context) BlockAlign() int {
	return int(ctx.block_align)
}

func (ctx *Context) BrdScale() int {
	return int(ctx.brd_scale)
}

func (ctx *Context) Channels() int {
	return int(ctx.channels)
}

func (ctx *Context) Chromaoffset() int {
	return int(ctx.chromaoffset)
}

func (ctx *Context) CodedHeight() int {
	return int(ctx.coded_height)
}

func (ctx *Context) CodedWidth() int {
	return int(ctx.coded_width)
}

func (ctx *Context) CoderType() int {
	return int(ctx.coder_type)
}

func (ctx *Context) CompressionLevel() int {
	return int(ctx.compression_level)
}

func (ctx *Context) ContextModel() int {
	return int(ctx.context_model)
}

func (ctx *Context) Cutoff() int {
	return int(ctx.cutoff)
}

func (ctx *Context) DarkMasking() float64 {
	return float64(ctx.dark_masking)
}

func (ctx *Context) DctAlgo() int {
	return int(ctx.dct_algo)
}

func (ctx *Context) Debug() int {
	return int(ctx.debug)
}

func (ctx *Context) DebugMv() int {
	return int(ctx.debug_mv)
}

func (ctx *Context) Delay() int {
	return int(ctx.delay)
}

func (ctx *Context) DiaSize() int {
	return int(ctx.dia_size)
}

func (ctx *Context) ErrRecognition() int {
	return int(ctx.err_recognition)
}

func (ctx *Context) ErrorConcealment() int {
	return int(ctx.error_concealment)
}

func (ctx *Context) ExtradataSize() int {
	return int(ctx.extradata_size)
}

func (ctx *Context) Flags() int {
	return int(ctx.flags)
}

func (ctx *Context) Flags2() int {
	return int(ctx.flags2)
}

func (ctx *Context) FrameBits() int {
	return int(ctx.frame_bits)
}

func (ctx *Context) FrameNumber() int {
	return int(ctx.frame_number)
}

func (ctx *Context) FrameSize() int {
	return int(ctx.frame_size)
}

func (ctx *Context) FrameSkipCmp() int {
	return int(ctx.frame_skip_cmp)
}

func (ctx *Context) FrameSkipExp() int {
	return int(ctx.frame_skip_exp)
}

func (ctx *Context) FrameSkipFactor() int {
	return int(ctx.frame_skip_factor)
}

func (ctx *Context) FrameSkipThreshold() int {
	return int(ctx.frame_skip_threshold)
}

func (ctx *Context) GlobalQuality() int {
	return int(ctx.global_quality)
}

func (ctx *Context) GopSize() int {
	return int(ctx.gop_size)
}

func (ctx *Context) HasBFrames() int {
	return int(ctx.has_b_frames)
}

func (ctx *Context) HeaderBits() int {
	return int(ctx.header_bits)
}

func (ctx *Context) Height() int {
	return int(ctx.height)
}

func (ctx *Context) ICount() int {
	return int(ctx.i_count)
}

func (ctx *Context) IQuantFactor() float64 {
	return float64(ctx.i_quant_factor)
}

func (ctx *Context) IQuantOffset() float64 {
	return float64(ctx.i_quant_offset)
}

func (ctx *Context) ITexBits() int {
	return int(ctx.i_tex_bits)
}

func (ctx *Context) IdctAlgo() int {
	return int(ctx.idct_algo)
}

func (ctx *Context) IldctCmp() int {
	return int(ctx.ildct_cmp)
}

func (ctx *Context) IntraDcPrecision() int {
	return int(ctx.intra_dc_precision)
}

func (ctx *Context) KeyintMin() int {
	return int(ctx.keyint_min)
}

func (ctx *Context) LastPredictorCount() int {
	return int(ctx.last_predictor_count)
}

func (ctx *Context) Level() int {
	return int(ctx.level)
}

func (ctx *Context) LogLevelOffset() int {
	return int(ctx.log_level_offset)
}

func (ctx *Context) Lowres() int {
	return int(ctx.lowres)
}

func (ctx *Context) LumiMasking() float64 {
	return float64(ctx.lumi_masking)
}

func (ctx *Context) MaxBFrames() int {
	return int(ctx.max_b_frames)
}

func (ctx *Context) MaxPredictionOrder() int {
	return int(ctx.max_prediction_order)
}

func (ctx *Context) MaxQdiff() int {
	return int(ctx.max_qdiff)
}

func (ctx *Context) MbCmp() int {
	return int(ctx.mb_cmp)
}

func (ctx *Context) MbDecision() int {
	return int(ctx.mb_decision)
}

func (ctx *Context) MbLmax() int {
	return int(ctx.mb_lmax)
}

func (ctx *Context) MbLmin() int {
	return int(ctx.mb_lmin)
}

func (ctx *Context) MeCmp() int {
	return int(ctx.me_cmp)
}

func (ctx *Context) MePenaltyCompensation() int {
	return int(ctx.me_penalty_compensation)
}

func (ctx *Context) MePreCmp() int {
	return int(ctx.me_pre_cmp)
}

func (ctx *Context) MeRange() int {
	return int(ctx.me_range)
}

func (ctx *Context) MeSubCmp() int {
	return int(ctx.me_sub_cmp)
}

func (ctx *Context) MeSubpelQuality() int {
	return int(ctx.me_subpel_quality)
}

func (ctx *Context) MinPredictionOrder() int {
	return int(ctx.min_prediction_order)
}

func (ctx *Context) MiscBits() int {
	return int(ctx.misc_bits)
}

func (ctx *Context) MpegQuant() int {
	return int(ctx.mpeg_quant)
}

func (ctx *Context) Mv0Threshold() int {
	return int(ctx.mv0_threshold)
}

func (ctx *Context) MvBits() int {
	return int(ctx.mv_bits)
}

func (ctx *Context) NoiseReduction() int {
	return int(ctx.noise_reduction)
}

func (ctx *Context) NsseWeight() int {
	return int(ctx.nsse_weight)
}

func (ctx *Context) PCount() int {
	return int(ctx.p_count)
}

func (ctx *Context) PMasking() float64 {
	return float64(ctx.p_masking)
}

func (ctx *Context) PTexBits() int {
	return int(ctx.p_tex_bits)
}

func (ctx *Context) PreDiaSize() int {
	return int(ctx.pre_dia_size)
}

func (ctx *Context) PreMe() int {
	return int(ctx.pre_me)
}

func (ctx *Context) PredictionMethod() int {
	return int(ctx.prediction_method)
}

func (ctx *Context) Profile() int {
	return int(ctx.profile)
}

func (ctx *Context) Qblur() float64 {
	return float64(ctx.qblur)
}

func (ctx *Context) Qcompress() float64 {
	return float64(ctx.qcompress)
}

func (ctx *Context) Qmax() int {
	return int(ctx.qmax)
}

func (ctx *Context) Qmin() int {
	return int(ctx.qmin)
}

func (ctx *Context) RcBufferSize() int {
	return int(ctx.rc_buffer_size)
}

func (ctx *Context) RcInitialBufferOccupancy() int {
	return int(ctx.rc_initial_buffer_occupancy)
}

func (ctx *Context) RcMaxAvailableVbvUse() float64 {
	return float64(ctx.rc_max_available_vbv_use)
}

func (ctx *Context) RcMaxRate() int {
	return int(ctx.rc_max_rate)
}

func (ctx *Context) RcMinRate() int {
	return int(ctx.rc_min_rate)
}

func (ctx *Context) RcMinVbvOverflowUse() float64 {
	return float64(ctx.rc_min_vbv_overflow_use)
}

func (ctx *Context) RcOverrideCount() int {
	return int(ctx.rc_override_count)
}

func (ctx *Context) RefcountedFrames() int {
	return int(ctx.refcounted_frames)
}

func (ctx *Context) Refs() int {
	return int(ctx.refs)
}

func (ctx *Context) RtpPayloadSize() int {
	return int(ctx.rtp_payload_size)
}

func (ctx *Context) SampleRate() int {
	return int(ctx.sample_rate)
}

func (ctx *Context) ScenechangeThreshold() int {
	return int(ctx.scenechange_threshold)
}

func (ctx *Context) SeekPreroll() int {
	return int(ctx.seek_preroll)
}

func (ctx *Context) SideDataOnlyPackets() int {
	return int(ctx.side_data_only_packets)
}

func (ctx *Context) SkipAlpha() int {
	return int(ctx.skip_alpha)
}

func (ctx *Context) SkipBottom() int {
	return int(ctx.skip_bottom)
}

func (ctx *Context) SkipCount() int {
	return int(ctx.skip_count)
}

func (ctx *Context) SkipTop() int {
	return int(ctx.skip_top)
}

func (ctx *Context) SliceCount() int {
	return int(ctx.slice_count)
}

func (ctx *Context) SliceFlags() int {
	return int(ctx.slice_flags)
}

func (ctx *Context) Slices() int {
	return int(ctx.slices)
}

func (ctx *Context) SpatialCplxMasking() float64 {
	return float64(ctx.spatial_cplx_masking)
}

func (ctx *Context) StrictStdCompliance() int {
	return int(ctx.strict_std_compliance)
}

func (ctx *Context) SubCharencMode() int {
	return int(ctx.sub_charenc_mode)
}

func (ctx *Context) SubtitleHeaderSize() int {
	return int(ctx.subtitle_header_size)
}

func (ctx *Context) TemporalCplxMasking() float64 {
	return float64(ctx.temporal_cplx_masking)
}

func (ctx *Context) ThreadCount() int {
	return int(ctx.thread_count)
}

func (ctx *Context) ThreadSafeCallbacks() int {
	return int(ctx.thread_safe_callbacks)
}

func (ctx *Context) ThreadType() int {
	return int(ctx.thread_type)
}

func (ctx *Context) TicksPerFrame() int {
	return int(ctx.ticks_per_frame)
}

func (ctx *Context) Trellis() int {
	return int(ctx.trellis)
}

func (ctx *Context) Width() int {
	return int(ctx.width)
}

func (ctx *Context) WorkaroundBugs() int {
	return int(ctx.workaround_bugs)
}

func (ctx *Context) AudioServiceType() AvAudioServiceType {
	return (AvAudioServiceType)(ctx.audio_service_type)
}

func (ctx *Context) ChromaSampleLocation() AvChromaLocation {
	return (AvChromaLocation)(ctx.chroma_sample_location)
}

func (ctx *Context) CodecDescriptor() *Descriptor {
	return (*Descriptor)(ctx.codec_descriptor)
}

func (ctx *Context) CodecId() CodecId {
	return (CodecId)(ctx.codec_id)
}

func (ctx *Context) CodecType() MediaType {
	return (MediaType)(ctx.codec_type)
}

func (ctx *Context) ColorPrimaries() AvColorPrimaries {
	return (AvColorPrimaries)(ctx.color_primaries)
}

func (ctx *Context) ColorRange() AvColorRange {
	return (AvColorRange)(ctx.color_range)
}

func (ctx *Context) ColorTrc() AvColorTransferCharacteristic {
	return (AvColorTransferCharacteristic)(ctx.color_trc)
}

func (ctx *Context) Colorspace() AvColorSpace {
	return (AvColorSpace)(ctx.colorspace)
}

func (ctx *Context) FieldOrder() AvFieldOrder {
	return (AvFieldOrder)(ctx.field_order)
}

func (ctx *Context) PixFmt() swscale.PixelFormat {
	return (swscale.PixelFormat)(ctx.pix_fmt)
}

func (ctx *Context) RequestSampleFmt() AvSampleFormat {
	return (AvSampleFormat)(ctx.request_sample_fmt)
}

func (ctx *Context) SampleFmt() AvSampleFormat {
	return (AvSampleFormat)(ctx.sample_fmt)
}

func (ctx *Context) SkipFrame() AvDiscard {
	return (AvDiscard)(ctx.skip_frame)
}

func (ctx *Context) SkipIdct() AvDiscard {
	return (AvDiscard)(ctx.skip_idct)
}

func (ctx *Context) SkipLoopFilter() AvDiscard {
	return (AvDiscard)(ctx.skip_loop_filter)
}
