package monotime

import (
	"sync/atomic"
)

type IdleTimer struct {
	last uint64
}

func (t *IdleTimer) Reset() {
	atomic.StoreUint64(&t.last, Now())
}

func (t *IdleTimer) NanosecSince() uint64 {
	return Now() - atomic.LoadUint64(&t.last)
}
