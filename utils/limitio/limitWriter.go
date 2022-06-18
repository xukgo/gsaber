package limitio

import (
	"github.com/juju/ratelimit"
	"io"
	"sync/atomic"
)

type LimitWriter struct {
	writer     io.Writer
	bucket     *ratelimit.Bucket
	totalCount int64
}

func NewLimitWriter(writer io.Writer, bucket *ratelimit.Bucket) *LimitWriter {
	return &LimitWriter{
		writer:     writer,
		bucket:     bucket,
		totalCount: 0,
	}
}

func (w *LimitWriter) Write(p []byte) (n int, err error) {
	if w.bucket != nil {
		count := int64(len(p))
		w.bucket.Wait(count)
	}
	n, err = w.writer.Write(p)
	atomic.AddInt64(&w.totalCount, int64(n))
	return
}

func (w *LimitWriter) GetTotalWriteCount() int64 {
	n := atomic.LoadInt64(&w.totalCount)
	return n
}
