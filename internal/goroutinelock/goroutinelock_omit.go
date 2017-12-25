// +build nogoroutinelock

package goroutinelock

type GoroutineLock struct{}

func New() GoroutineLock {
	return GoroutineLock{}
}

func (g GoroutineLock) Check() {}

func (g GoroutineLock) CheckNotOn() {}
