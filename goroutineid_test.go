// +build !nogoroutinechecks

package deepreturn

import "testing"

func Test_checkGoroutineID_expectMatch_good(t *testing.T) {
	curr := getGoroutineID()

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Goroutine IDs for same goroutine didn't match, expected it to match: %s", r)
		}
	}()

	checkGoroutineID(curr, true)
}

func Test_checkGoroutineID_expectMatch_bad(t *testing.T) {
	gidch := make(chan goroutineID)
	go func() {
		gidch <- getGoroutineID()
		close(gidch)
	}()
	other := <-gidch

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Code didn't panic for goroutines with different IDs")
		}
	}()

	checkGoroutineID(other, true)
}

func Test_checkGoroutineID_expectMisMatch_good(t *testing.T) {
	gidch := make(chan goroutineID)
	go func() {
		gidch <- getGoroutineID()
		close(gidch)
	}()
	other := <-gidch

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Goroutine IDs for other goroutine did match, expected it to not match: %s", r)
		}
	}()

	checkGoroutineID(other, false)
}

func Test_checkGoroutineID_expectMisMatch_bad(t *testing.T) {
	curr := getGoroutineID()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Code didn't panic for goroutines with same IDs")
		}
	}()

	checkGoroutineID(curr, false)
}
