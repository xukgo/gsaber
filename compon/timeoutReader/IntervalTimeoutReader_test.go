package timeoutReader

import (
	"fmt"
	"io"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_IntervalTimeoutReader_01(t *testing.T) {
	pr, pw := io.Pipe()
	wg := new(sync.WaitGroup)

	tmReader := NewIntervalTimeoutReader(pr, time.Second, time.Millisecond*50)
	defer tmReader.CancelDetect()

	wg.Add(1)
	go tmReader.StartDetect(func() {
		wg.Done()
	})

	wg.Add(1)
	go func() {
		defer pw.Close()
		pw.Write([]byte("hello1"))
		pw.Write([]byte("hello2"))
		time.Sleep(time.Millisecond * 1200)
		pw.Write([]byte("hello3"))
		wg.Done()
	}()

	buf := make([]byte, 256)
	n, err := tmReader.Read(buf)
	assert.True(t, err == nil)
	fmt.Printf("read %d bytes: %s\n", n, string(buf[:n]))

	n, err = tmReader.Read(buf)
	assert.True(t, err == nil)
	fmt.Printf("read %d bytes: %s\n", n, string(buf[:n]))

	n, err = tmReader.Read(buf)
	assert.True(t, err != nil)
	assert.True(t, strings.Contains(err.Error(), "timeout"))
	fmt.Printf("read error: %s\n", err.Error())

	wg.Wait()
}
