// Package iostreams holds the input and output streams commands use, so the
// real CLI can talk to the terminal while tests capture output in buffers.
package iostreams

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/jwalton/go-supportscolor"
	"golang.org/x/term"
)

type IOStreams struct {
	In     io.Reader
	Out    io.Writer
	ErrOut io.Writer

	colorEnabled bool
	stderrTTY    bool
	canPrompt    bool
	readSecret   func() (string, error)
}

// System returns streams connected to the process's stdin, stdout and stderr.
func System() *IOStreams {
	return &IOStreams{
		In:           os.Stdin,
		Out:          os.Stdout,
		ErrOut:       os.Stderr,
		colorEnabled: supportscolor.Stderr().SupportsColor,
		stderrTTY:    term.IsTerminal(int(os.Stderr.Fd())),
		// Prompts are written to stderr and answered on stdin, so both must be
		// a terminal. Stdout may still be redirected.
		canPrompt: term.IsTerminal(int(os.Stdin.Fd())) && term.IsTerminal(int(os.Stderr.Fd())),
		readSecret: func() (string, error) {
			secret, err := term.ReadPassword(int(os.Stdin.Fd()))
			return string(secret), err
		},
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

// IsStderrTTY reports whether messages are shown in a terminal, rather than
// captured by a script or CI job.
func (s *IOStreams) IsStderrTTY() bool {
	return s.stderrTTY
}

// CanPrompt reports whether the user can answer prompts, i.e. the CLI is
// running in a terminal rather than a script or CI job.
func (s *IOStreams) CanPrompt() bool {
	return s.canPrompt
}

// ReadSecret writes prompt to ErrOut and reads a line from the terminal
// without showing what is typed.
func (s *IOStreams) ReadSecret(prompt string) (string, error) {
	if !s.canPrompt || s.readSecret == nil {
		return "", errors.New("cannot prompt: not running in a terminal")
	}
	fmt.Fprint(s.ErrOut, prompt)
	secret, err := s.readSecret()
	// The newline the user typed isn't echoed either.
	fmt.Fprintln(s.ErrOut)
	return secret, err
}

// Confirm writes a yes/no question to ErrOut and reads the answer from In.
// Anything other than "y" or "yes" counts as no.
func (s *IOStreams) Confirm(question string) (bool, error) {
	if !s.canPrompt {
		return false, errors.New("cannot prompt: not running in a terminal")
	}
	fmt.Fprintf(s.ErrOut, "%s [y/N] ", question)
	answer, err := bufio.NewReader(s.In).ReadString('\n')
	if err != nil && !errors.Is(err, io.EOF) {
		return false, err
	}
	switch strings.ToLower(strings.TrimSpace(answer)) {
	case "y", "yes":
		return true, nil
	}
	return false, nil
}

// SetTerminalInput makes test streams behave like a terminal, answering
// secret prompts with readSecret.
func (s *IOStreams) SetTerminalInput(readSecret func() (string, error)) {
	s.canPrompt = true
	s.stderrTTY = true
	s.readSecret = readSecret
}
