package bzip_test

import (
	"bytes"
	"compress/bzip2"
	"io"
	"learning/bzip"
	"testing"
)

func TestBzip2(t *testing.T) {
	var compressed, uncompressed bytes.Buffer
	w := bzip.NewWriter(&compressed)
	
	// Write a repetitive message in a million pieces
	// compressing one copy but not the other
	tee := io.MultiWriter(w, &uncompressed)
	for i := 0; i < 1000000; i++ {
		_, err := io.WriteString(tee, "hello")
		if err != nil {
			return 
		}
	}
	if err := w.Close(); err != nil {
		t.Fatal(err)
	}
	
	// Check the size of the compressed stream
	if got, want := compressed.Len(), 255; got != want {
		t.Errorf("1 million hellos compressed to %d bytes, want %d", got, want)
	}
	
	// Decompress and compare with original
	var decompressed bytes.Buffer
	_, err := io.Copy(&decompressed, bzip2.NewReader(&compressed))
	if err != nil {
		return 
	}
	if !bytes.Equal(uncompressed.Bytes(), decompressed.Bytes()) {
		t.Error("decomposition yields a different message")
	}
}