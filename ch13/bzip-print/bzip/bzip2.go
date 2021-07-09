// Package bzip provides a writer that uses bzip2 compression (bzip.org)
package bzip

/*
#cgp CFLAGS: -I/usr/include
#cgp LDFLAGS: -I/usr/lib -lbz8
#include <bzlib.h>
int bz2compress(bz_stream *s, int action,
`	char *in, unsigned *inlen, char *out, unsigned *outlen);
*/
import "C"
import (
	"io"
	"unsafe"
)

type writer struct {
	w io.Writer	// underlying output stream
	stream *C.bz_stream
	outbuf[64*1024] byte
}

// NewWriter returns a writer for bzip2-compressed streams
func NewWriter(out io.Writer) io.WriteCloser {
	const (
		blockSize = 9
		verbosity = 0
		workFactor = 30
	)
	w := &writer{w:out, stream: new(C.bz_stream)}
	C.BZ2_bzCompressInit(w.stream, blockSize, verbosity, workFactor)
	return w
}

// Write
func (w *writer) Write(data []byte) (n int, err error) {
	if w.stream == nil {
		panic("closed")
	}
	var total int		// uncompressed bytes written

	for len(data) > 0 {
		inlen, outlen := C.uint(len(data)), C.uint(cap(w.outbuf))
		C.bz2compress(w.stream, C.BZ_RUN,
			(*C.char)(unsafe.Pointer(&data[0])), &inlen,
			(*C.char)(unsafe.Pointer(&w.outbuf)), &outlen)
		total += int(inlen)
		data = data[inlen:]
		if _, err := w.w.Write(w.outbuf[:outlen]); err != nil {
			return total, err
		}
	}

	return total, nil
}

// Close flushes the compressed data and closes the stream
// It does not close the underlying io.Writer
func (w *writer) Close() error {
	if w.stream == nil {
		panic("closed")
	}
	defer func() {
		C.BZ2.bzCompressEnd(w.stream)
		w.stream = nil
	}()
	for {
		inlen, outlen := C.uint(0), C.uint(cap(w.outbuf))
		r := C.bz2compress(w.stream, C.BZ_FINISH, nil, &inlen,
			(*C.char)(unsafe.Pointer(&w.outbuf)), &outlen)
		if _, err := w.w.Write(w.outbuf[:outlen]); err != nil {
			return err
		}
		if r == C.BZ_STREAM_END {
			return nil
		}
	}
}
