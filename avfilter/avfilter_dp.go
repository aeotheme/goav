//+build DoDeprecated

package avfilter

//#cgo pkg-config: libavfilter
//#include <libavfilter/avfilter.h>
import "C"

// AvfilterLinkGetChannels deprecated
//Get the number of channels of a link.
func AvfilterLinkGetChannels(l *Link) int {
	return int(C.avfilter_link_get_channels((*C.struct_AVFilterLink)(l)))
}

// AvfilterLinkSetClosed deprecated
//Set the closed field of a link.
func AvfilterLinkSetClosed(l *Link, c int) {
	C.avfilter_link_set_closed((*C.struct_AVFilterLink)(l), C.int(c))
}

// AvfilterRegisterAll deprecated
//Initialize the filter system.
func AvfilterRegisterAll() {
	C.avfilter_register_all()
}

// AvfilterRegister deprecated
//Register a filter.
func (f *Filter) AvfilterRegister() int {
	return int(C.avfilter_register((*C.struct_AVFilter)(f)))
}

// AvfilterNext deprecated
//Iterate over all registered filters.
func (f *Filter) AvfilterNext() *Filter {
	return (*Filter)(C.avfilter_next((*C.struct_AVFilter)(f)))
}
