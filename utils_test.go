package deepreturn

import "testing"
import "fmt"

func TestWaitForErr_Error(t *testing.T) {
	ch := make(chan interface{})
	exp := fmt.Errorf("some error")
	go func() {
		ch <- exp
	}()
	err := WaitForErr(ch)
	if err != exp {
		t.Errorf("Expected error to be %v, but it was %v", exp, err)
	}
}

func TestWaitForErr_NonError(t *testing.T) {
	ch := make(chan interface{})
	exp := "some string (not an error)"
	go func() {
		ch <- exp
	}()
	err := WaitForErr(ch)
	if err != nil {
		t.Errorf("Expected error to be %v, but it was %v", nil, err)
	}
}
