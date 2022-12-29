// Package common includes code used by multiple packages
package common

import (
	"image"
	"image/png"
	"io"
	"os"
)

// OS interface used for file action
//go:generate go run github.com/vektra/mockery/cmd/mockery -name OS -inpkg --filename os_mock.go
type OS interface {
	Create(name string) (*os.File, error)
	Encode(w io.Writer, m image.Image) error
}

// OSReal implement file actions
type OSReal struct{}

// Create creates or truncates the named file. If the file already exists,
// it is truncated. If the file does not exist, it is created with mode 0666
// (before umask). If successful, methods on the returned File can
// be used for I/O; the associated file descriptor has mode O_RDWR.
// If there is an error, it will be of type *PathError.
func (*OSReal) Create(name string) (*os.File, error) {
	return os.Create(name)
}

// Encode writes the Image m to w in PNG format. Any Image may be
// encoded, but images that are not image.NRGBA might be encoded lossily.
func (*OSReal) Encode(w io.Writer, m image.Image) error {
	return png.Encode(w, m)
}
