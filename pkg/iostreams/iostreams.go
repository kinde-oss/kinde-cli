// Package iostreams holds the input and output streams commands use, so the
// real CLI can talk to the terminal while tests capture output in buffers.
package iostreams

import (
	"bytes"
	"io"
	"os"

	"github.com/jwalton/go-supportscolor"
)

type IOStreams struct {
	In     io.Reader
	Out    io.Writer
	ErrOut io.Writer

	colorEnabled bool
}

// System returns streams connected to the process's stdin, stdout and stderr.
func System() *IOStreams {
	return &IOStreams{
		In:           os.Stdin,
		Out:          os.Stdout,
		ErrOut:       os.Stderr,
		colorEnabled: supportscolor.Stderr().SupportsColor,
	}
}

// Test returns streams backed by buffers, along with those buffers so tests
// can provide input and inspect output.
func Test() (ios *IOStreams, in *bytes.Buffer, out *bytes.Buffer, errOut *bytes.Buffer) {
	in, out, errOut = &bytes.Buffer{}, &bytes.Buffer{}, &bytes.Buffer{}
	return &IOStreams{In: in, Out: out, ErrOut: errOut}, in, out, errOut
}

// ColorEnabled reports whether output written to ErrOut may use color.
func (s *IOStreams) ColorEnabled() bool {
	return s.colorEnabled
}
