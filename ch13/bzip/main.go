package main

import (
	"bytes"
	"io"
)

func main() {
	var _ io.Writer = new(bytes.Buffer)
}
