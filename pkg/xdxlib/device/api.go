package device

import (
	"github.com/Longpai/xdxml-lib/pkg/xdxlib/xdxml"
)

// Interface provides the API to the 'device' package
type Interface interface {
	GetDevices() ([]Device, error)
	VisitDevices(func(i int, d Device) error) error
}

type devicelib struct {
	xdxml xdxml.Interface
}

var _ Interface = &devicelib{}

func New(opts ...Option) Interface {
	d := &devicelib{}
	for _, opt := range opts {
		opt(d)
	}
	if d.xdxml == nil {
		d.xdxml = xdxml.New()
	}
	return d
}

func WithXdxml(xdxml xdxml.Interface) Option {
	return func(d *devicelib) {
		d.xdxml = xdxml
	}
}

// Option defines a function for passing options to the New() call
type Option func(*devicelib)
