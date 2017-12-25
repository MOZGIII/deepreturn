// +build !nogoroutinelock

package goroutinelock

import (
	"github.com/MOZGIII/deepreturn/internal/http2goroutineid"
)

// GoroutineLock represents an goroutine lock.
type GoroutineLock goroutineID

// New returns a new goroutine lock set to check against
// the goroutine that called New.
func New() GoroutineLock {
	return GoroutineLock(getGoroutineID())
}

// Check panics if goroutine that called Check is not
// the same that called New.
func (g GoroutineLock) Check() {
	if getGoroutineID() != goroutineID(g) {
		panic("running on the wrong goroutine")
	}
}

// CheckNotOn panics if goroutine that called CheckNotOn
// is the same that called New.
func (g GoroutineLock) CheckNotOn() {
	if getGoroutineID() == goroutineID(g) {
		panic("running on the wrong goroutine")
	}
}

type goroutineID uint64

func getGoroutineID() goroutineID {
	return goroutineID(http2goroutineid.CurGoroutineID())
}
