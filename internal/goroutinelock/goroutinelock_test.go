// +build !nogoroutinelock

package goroutinelock

import "testing"

func TestGoroutineLock_Check_good(t *testing.T) {
	lock := New()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Goroutine lock check didn't pass, expected it to pass: %s", r)
		}
	}()

	lock.Check()
}

func TestGoroutineLock_Check_bad(t *testing.T) {
	lockch := make(chan GoroutineLock)
	go func() {
		lockch <- New()
		close(lockch)
	}()
	lock := <-lockch

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Code didn't panic for failed goroutine lock check")
		}
	}()

	lock.Check()
}

func TestGoroutineLock_CheckNotOn_good(t *testing.T) {
	lockch := make(chan GoroutineLock)
	go func() {
		lockch <- New()
		close(lockch)
	}()
	lock := <-lockch

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Goroutine lock check-not-on didn't pass, expected it to pass: %s", r)
		}
	}()

	lock.CheckNotOn()
}

func TestGoroutineLock_CheckNotOn_bad(t *testing.T) {
	lock := New()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Code didn't panic for failed goroutine lock check-not-on")
		}
	}()

	lock.CheckNotOn()
}
