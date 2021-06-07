# goav
Golang binding for FFmpeg

A comprehensive binding to the ffmpeg video/audio manipulation library.

Forked from https://github.com/giorgisio/goav

## Libraries

* `avcodec` corresponds to the ffmpeg library: libavcodec [provides implementation of a wider range of codecs]
* `avformat` corresponds to the ffmpeg library: libavformat [implements streaming protocols, container formats and basic I/O access]
* `avutil` corresponds to the ffmpeg library: libavutil [includes hashers, decompressors and miscellaneous utility functions]
* `avfilter` corresponds to the ffmpeg library: libavfilter [provides a mean to alter decoded Audio and Video through chain of filters]
* `avdevice` corresponds to the ffmpeg library: libavdevice [provides an abstraction to access capture and playback devices]
* `swresample` corresponds to the ffmpeg library: libswresample [implements audio mixing and resampling routines]
* `swscale` corresponds to the ffmpeg library: libswscale [implements color conversion and scaling routines]


## Note
- Function names in Go are consistent with that of the libraries to help with easy search
- [cgo: Extending Go with C](http://blog.giorgis.io/cgo-examples)
- goav comes with absolutely no warranty.

## Contribute
- Fork this repo and create your own feature branch.
- Follow standard Go conventions
- Test your code.
- Create pull request

## TODO

- [ ] Returning Errors
- [ ] Garbage Collection
- [X] Review included/excluded functions from each library
- [ ] Go Tests
- [ ] Possible restructuring packages
- [x] Tutorial01.c
- [ ] More Tutorial

## License
This library is under the [MIT License](http://opensource.org/licenses/MIT)
