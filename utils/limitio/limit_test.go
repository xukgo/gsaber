package limitio

import (
	"fmt"
	"github.com/juju/ratelimit"
	"testing"
	"time"
)

func TestLimitRate(t *testing.T) {
	var bucket = ratelimit.NewBucketWithQuantum(time.Millisecond*100, 100, 100)
	startAt := time.Now()
	bucket.Wait(100)
	fmt.Printf("elapse time ms:%d\n", time.Since(startAt))
	startAt = time.Now()
	bucket.Wait(200)
	fmt.Printf("elapse time ms:%d\n", time.Since(startAt))
	startAt = time.Now()
	bucket.Wait(300)
	fmt.Printf("elapse time ms:%d\n", time.Since(startAt))
	startAt = time.Now()
	bucket.Wait(150)
	fmt.Printf("elapse time ms:%d\n", time.Since(startAt))
}
