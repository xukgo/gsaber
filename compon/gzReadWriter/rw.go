package gzReadWriter

import (
	"compress/gzip"
	"fmt"
	"io"
)

type ReadWriter struct {
	reader   io.Reader
	gzWriter *gzip.Writer
}

func NewGzReadWriter(reader io.Reader) *ReadWriter {
	return &ReadWriter{reader: reader}
}

func (this *ReadWriter) SetWriter(writer io.Writer) error {
	gzWriter, err := gzip.NewWriterLevel(writer, gzip.DefaultCompression)
	if err != nil {
		return err
	}
	this.gzWriter = gzWriter
	return nil
}

func (this *ReadWriter) Close() error {
	if this.gzWriter != nil {
		this.gzWriter.Flush()
		return this.gzWriter.Close()
	}
	return nil
}

func (this *ReadWriter) ReadWrite(p []byte) (n int, err error) {
	if this.gzWriter == nil {
		return 0, fmt.Errorf("gzWriter is nil")
	}
	n, err = this.reader.Read(p)
	if n > 0 {
		n, err = this.gzWriter.Write(p[:n])
		return n, err
	}
	if err == io.EOF {
		return n, io.EOF
	}
	if err != nil {
		return 0, err
	}
	return 0, nil
}
