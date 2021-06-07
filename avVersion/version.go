package avVersion

import (
	"fmt"
	"github.com/aeotheme/goav/avcodec"
	"github.com/aeotheme/goav/avdevice"
	"github.com/aeotheme/goav/avfilter"
	"github.com/aeotheme/goav/avformat"
	"github.com/aeotheme/goav/avutil"
	"github.com/aeotheme/goav/swresample"
	"github.com/aeotheme/goav/swscale"
)

// Register all formats and codecs
//avformat.AvRegisterAll()
//avcodec.AvcodecRegisterAll()

func AvVersionFormat(s string, v uint) (ver string) {
	return fmt.Sprintf("%-9s %2d.%02d.%04d.%04d", s, (v&0xFF0000)>>16, (v&0xFF00)>>8, (v&0xFF)>>0, (v&0xFF000000)>>24)
}

func AvVersion(v uint) (ver string) {
	return fmt.Sprintf("%02d.%02d.%04d.%04d", (v&0xFF0000)>>16, (v&0xFF00)>>8, (v&0xFF)>>0, (v&0xFF000000)>>24)
}

func String() (s string) {
	s = fmt.Sprintf("ffmpeg library Version:\n%s\n%s\n%s\n%s\n%s\n%s\n%s",
		AvVersionFormat(`AvDevice`, avdevice.Version()),
		AvVersionFormat(`AvFormat`, avformat.Version()),
		AvVersionFormat(`AvCodec`, avcodec.Version()),
		AvVersionFormat(`AvUtil`, avutil.Version()),
		AvVersionFormat(`Resample`, swresample.Version()),
		AvVersionFormat(`SWScale`, swscale.Version()),
		AvVersionFormat(`AvFilter`, avfilter.Version()),
	)
	return
}

func Println() {
	fmt.Printf("Version:\n%s\n", String())
}
