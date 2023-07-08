package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

// a custom error type (audio file processing)

type errKind int

const (
	_ errKind = iota
	noHeader
	cantReadHeader
	invalidHeaderType
	invalidChunkLength
	invalidChunkType
	invalidDataLength
)

type WaveError struct {
	kind  errKind
	value int
	err   error
}

func (e WaveError) Error() string {
	switch e.kind {
	case noHeader:
		return "no header (file too short?)"
	case cantReadHeader:
		return fmt.Sprintf("can't read header[%d]: %s", e.value, e.err.Error())
	case invalidHeaderType:
		return "invalid header type"
	case invalidChunkLength:
		return fmt.Sprintf("invalid chunk length: %d", e.value)
	default:
		return "unhandled error" // ???
	}
}

// with returns an error with a particular value (e.g. header type)
func (e WaveError) with(val int) WaveError {
	e1 := e
	e1.value = val
	return e1
}

// from returns an error with a particular location and underlying error (e.g. from standard library)
func (e WaveError) from(pos int, err error) WaveError {
	e1 := e
	e1.value = pos
	e1.err = err
	return e1
}

// Prototypical errors

var (
	HeaderMissing      = WaveError{kind: noHeader}
	HeaderReadFailed   = WaveError{kind: cantReadHeader}
	InvalidHeaderType  = WaveError{kind: invalidHeaderType}
	InvalidChunkLength = WaveError{kind: invalidChunkLength}
	InvalidChunkType   = WaveError{kind: invalidChunkType}
	InvalidDataLength  = WaveError{kind: invalidDataLength}
)

// Example of how these are used
type Header struct {
	TotalLength uint32
	riff        os.File
}

const headerSize = 256

func DecodeHeader(b []byte) (*Header, []byte, error) {
	var err error
	var pos int

	header := Header{TotalLength: uint32(len(b))}
	buf := bytes.NewReader(b)

	if len(b) < headerSize {
		return &header, nil, HeaderMissing
	}

	if err = binary.Read(buf, binary.BigEndian, &header.riff); err != nil {
		return &header, nil, HeaderReadFailed.from(pos, err)
	}

	//...//
	// Note: RIFF is Resource Interchange File Format
	// A RIFF file starts out with a file header followed by a sequence of data chunks.
	return nil, nil, nil
}

// Wrapping and unwrapping errors

type HAL9009 struct {
	err    error
	victim string
}

func (h HAL9009) OpenPodBayDoors() error {
	// Logic ... //
	if h.err != nil {
		return fmt.Errorf("sorry %s, I can't: %w", h.victim, h.err)
	}
	// More logic ... //
	return nil

}
func (e WaveError) Unwrap() error {
	return e.err
}

func (we *WaveError) Is(err error) bool {
	e, ok := err.(WaveError) // reflection

	if !ok {
		return false
	}

	return e.err == we.err
}

// error.As looks at the chain of errors and tries to extract the error being looked for

// panic and recover behaves like exception handling

func f() {
	panic("oh, no!")
}

func main() {
	defer func() {
		if p := recover(); p != nil { // recover captures the panic created by f()
			fmt.Println("recover:", p)
		}
	}()

	f()
}
