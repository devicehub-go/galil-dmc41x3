package galildmc41x3

import (
	"github.com/devicehub-go/galil-dmc41x3/protocol"
	"github.com/devicehub-go/unicomm"
)

// Creates a new instance of Galil DMC41x3 Controller
func New(options unicomm.Options) *protocol.DMC41x3 {
	options.Delimiter = "\r"
	return &protocol.DMC41x3{
		Communication: unicomm.New(options),
	}
}
