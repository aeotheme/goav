package exam

import (
	"fmt"
	"github.com/aeotheme/goav/avcodec"
	"github.com/aeotheme/goav/avformat"
	"github.com/aeotheme/goav/avutil"
	"github.com/aeotheme/goav/swscale"
	"log"
	"os"
	"unsafe"
)

const pf = avcodec.AV_PIX_FMT_RGB24
const contextFlag = avcodec.SWS_BILINEAR

//const contextFlag = avcodec.SWS_BICUBIC

func ExamFfmpeg() {
	//if len(os.Args) < 2 {
	//	fmt.Println("Please provide a movie file")
	//	return
	//}
	dumpFrames(`sample.mp4`)
	//dumpFrames(`dump_mpegts_01.data`)
}

type paraCodec struct {
	index  int
	codeId avcodec.CodecId
	ctx    *avcodec.Context
	c      *avcodec.Codec
}

func dumpFrames(file string) {

	// Open video file
	pFormatContext := avformat.AvFormatAllocContext()
	defer func() {
		pFormatContext.AvFormatFreeContext()
	}()

	var opts *avutil.Dictionary
	if avformat.AvFormatOpenInput(&pFormatContext, file, nil, &opts) != 0 {
		fmt.Printf("Unable to open file %s\n", file)
		return
	}
	defer func() {
		//pFormatContext.AvformatCloseInput()
	}()

	// Retrieve stream information
	if pFormatContext.AvformatFindStreamInfo(nil) < 0 {
		fmt.Println("Couldn't find stream information")
		return
	}

	// Dump information about file onto standard error
	//pFormatContext.AvDumpFormat(0, file, false)

	// Allocate video frame
	pFrame := avutil.AvFrameAlloc()
	// Allocate an AVFrame structure
	pFrameRGB := avutil.AvFrameAlloc()
	if pFrame == nil || pFrameRGB == nil {
		fmt.Println("Unable to allocate Frame")
		return
	}
	defer func() {
		// Free the YUV frame
		avutil.AvFrameFree(pFrameRGB)
		avutil.AvFrameFree(pFrame)
	}()

	var codecMaps map[int]paraCodec
	var codecSlices []paraCodec
	codecMaps = make(map[int]paraCodec)
	codecSlices = make([]paraCodec, 0)

	var pCodecCtx *avcodec.Context
	var pCodec *avcodec.Codec
	// Find the first video stream
	for i := 0; i < int(pFormatContext.NbStreams()); i++ {
		cp := pFormatContext.Streams()[i].CodecParameters()
		codeId := cp.AvCodecGetId()
		fmt.Printf("stream %d codeId %d\n", i, int(codeId))
		switch cp.AvCodecGetType() {
		case avformat.AVMEDIA_TYPE_VIDEO:
			if _, ok := codecMaps[int(codeId)]; ok {
				continue
			}
			codec := paraCodec{
				index:  i,
				codeId: codeId,
				//ctx: pFormatContext.Streams()[i].Codec(),
				c: avcodec.AvCodecFindDecoder(codeId),
			}
			Tc := (*avcodec.Codec)(nil)
			codec.ctx = Tc.AvCodecAllocContext3()
			r := codec.ctx.FromParameters(cp)
			if r >= 0 {
				codecMaps[int(codeId)] = codec
				codecSlices = append(codecSlices, codec)
			}
		}
	}
	for k := range codecSlices {
		defer func(i int) {
			codecSlices[i].ctx.AvcodecFreeContext()
		}(k)
	}

	if len(codecSlices) == 0 {
		return
	}
	mainStreamId := codecSlices[0].index
	pCodecCtx = codecSlices[0].ctx
	pCodec = codecSlices[0].c
	if pCodec == nil {
		fmt.Println("Unsupported codec!")
		return
	}
	// Open codec
	if pCodecCtx.AvcodecOpen2(pCodec, nil) < 0 {
		fmt.Println("Could not open codec")
		return
	}
	defer func() {
		// Close the codecs
		pCodecCtx.AvcodecClose()
	}()

	// Determine required buffer size and allocate buffer
	numBytes := uintptr(avutil.AvPictureGetBufferSize((avutil.PixelFormat)(avcodec.AV_PIX_FMT_RGB24),
		pCodecCtx.Width(), pCodecCtx.Height(), 1))
	if numBytes <= 0 {
		fmt.Println("unexpect size", numBytes)
		return
	}
	buffer := avutil.AvMalloc(numBytes)
	if buffer == nil {
		fmt.Println("Could not AvMalloc")
		return
	}
	defer func() {
		// Free the RGB image
		avutil.AvFree(buffer)
	}()

	// Assign appropriate parts of buffer to image planes in pFrameRGB
	// Note that pFrameRGB is an AVFrame, but AVFrame is a superset
	// of AVPicture
	a1 := avutil.DataPoint(pFrameRGB)
	a2 := avutil.LineSizePoint(pFrameRGB)
	a11 := pCodecCtx.Width()
	a12 := pCodecCtx.Height()
	avutil.AvPictureFill(a1, a2,
		(*uint8)(buffer), pf,
		a11, a12, 1)
	//avutil.AvPictureFill2(pFrameRGB,
	//	(*uint8)(buffer), pf,
	//	a11, a12, 1)

	// initialize SWS context for software scaling
	var swsCtx *swscale.Context
	a13 := pCodecCtx.PixFmt()
	swsCtx = swscale.SwsGetCachedContext(swsCtx,
		a11, a12, a13, a11, a12,
		pf, contextFlag, nil, nil, nil)
	if swsCtx == nil {
		fmt.Println("Could not SwsGetCachedContext")
		return
	}

	// Read frames and save first five frames to disk
	frameNumber := 0
	packet := avcodec.AvPacketAlloc()
	if packet == nil {
		fmt.Println("Could not AvPacketAlloc")
		return
	}
	defer func() {
		// Free the packet that was allocated by av_read_frame
		avutil.AvFree(unsafe.Pointer(packet))
	}()

loopS:
	for {
		for {
			if pFormatContext.AvReadFrame(packet) < 0 {
				// needRestart
				break loopS
			}
			// Is this a packet from the video stream?
			if packet.StreamIndex() == mainStreamId {
				break
			}
		}
		// Decode video frame
		response := pCodecCtx.AvcodecSendPacket(packet)
		if response < 0 {
			fmt.Printf("Error while sending a packet to the decoder: %s\n", avutil.ErrorFromCode(response))
		}
		for response >= 0 {
			response = pCodecCtx.AvcodecReceiveFrame(pFrame)
			if response < 0 {
				//fmt.Printf("Error while receiving a frame from the decoder: %s\n", avutil.ErrorFromCode(response))
				break
			}
			if response == avutil.AvErrorEAGAIN || response == avutil.AvErrorEOF {
				break
			}

			frameNumber++
			if frameNumber > 5 {
				break
			}
			// Convert the image from its native format to RGB
			swscale.SwsScale2(swsCtx, avutil.Data(pFrame),
				avutil.LineSize(pFrame), 0, pCodecCtx.Height(),
				avutil.Data(pFrameRGB), avutil.LineSize(pFrameRGB))
			//swscale.SwsScale3(swsCtx,pFrame,
			//	0, pCodecCtx.Height(),pFrameRGB)

			// Save the frame to disk
			fmt.Printf("Writing frame %d\n", frameNumber)
			SaveFrame(pFrameRGB, pCodecCtx.Width(), pCodecCtx.Height(), frameNumber)
		}
		packet.AvPacketUnref()
	}

}

// SaveFrame writes a single frame to disk as a PPM file
func SaveFrame(frame *avutil.Frame, width, height, frameNumber int) {
	// Open file
	fileName := fmt.Sprintf("frame%d.ppm", frameNumber)
	file, err := os.Create(fileName)
	if err != nil {
		log.Println("Error Reading")
	}
	defer func() {
		_ = file.Close()
	}()

	// Write header
	header := fmt.Sprintf("P6\n%d %d\n255\n", width, height)
	_, _ = file.Write([]byte(header))

	data0 := avutil.Data(frame)[0]
	if data0 == nil {
		return
	}

	// Write pixel data
	for y := 0; y < height; y++ {
		buf := make([]byte, width*3)
		startPos := uintptr(unsafe.Pointer(data0)) + uintptr(y)*uintptr(avutil.LineSize(frame)[0])
		_ = startPos
		for i := 0; i < width*3; i++ {
			element := *(*uint8)(unsafe.Pointer(startPos + uintptr(i)))
			buf[i] = element
		}
		_, _ = file.Write(buf)
	}
}
