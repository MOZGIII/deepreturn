package deepreturn

import (
	"testing"
)

func Test_DeferredStatementsAreCalledWhenTerminated(t *testing.T) {
	deferCheckCh := make(chan string, 10)
	deferWorked := "defer worked"

	valch := Start(func(terminate TerminatorFn) {
		defer func() { deferCheckCh <- deferWorked }()
		terminate(nil) // cause termination
		t.Error("should not be here")
	})
	val := <-valch
	if val != nil {
		t.Errorf("Expected val to be %v, but it was %v", nil, val)
	}

	select {
	case val := <-deferCheckCh:
		if val != "defer worked" {
			t.Errorf("Go unexpected results form defer check chan: %q, expected %s", val, deferWorked)
		}
	default:
		t.Errorf("Defer statement didn't work (can't read from defer check channel)")
	}
}

func Test_DeferredStatementsAreCalledWhenTerminatedDeep(t *testing.T) {
	deferCheckCh := make(chan string, 10)
	deferWorked := "defer worked"

	valch := Start(func(terminate TerminatorFn) {
		defer func() { deferCheckCh <- deferWorked }()
		f := func() {
			terminate(nil) // cause termination
			t.Error("should not be here (subfunction)")
		}
		f() // call f to cause termination in a function down the stack
		t.Error("should not be here")
	})
	val := <-valch
	if val != nil {
		t.Errorf("Expected val to be %v, but it was %v", nil, val)
	}

	select {
	case val := <-deferCheckCh:
		if val != "defer worked" {
			t.Errorf("Go unexpected results form defer check chan: %q, expected %s", val, deferWorked)
		}
	default:
		t.Errorf("Defer statement didn't work (can't read from defer check channel)")
	}
}

func Test_DeferredStatementsAreCalledWhenReturns(t *testing.T) {
	deferCheckCh := make(chan string, 10)
	deferWorked := "defer worked"

	valch := Start(func(terminate TerminatorFn) {
		defer func() { deferCheckCh <- deferWorked }()
		// function just returns
	})
	val := <-valch
	if val != nil {
		t.Errorf("Expected val to be %v, but it was %v", nil, val)
	}

	select {
	case val := <-deferCheckCh:
		if val != "defer worked" {
			t.Errorf("Go unexpected results form defer check chan: %q, expected %s", val, deferWorked)
		}
	default:
		t.Errorf("Defer statement didn't work (can't read from defer check channel)")
	}
}

func Test_TerminatingFromOtherGoroutinesNotAllowed(t *testing.T) {
	checkCh := make(chan string, 10)

	valch := Start(func(terminate TerminatorFn) {
		checkCh <- "first (just started)"

		go func() {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("The code did not panic")
				}
			}()

			checkCh <- "second (in goroutine)"

			terminate(nil) // terminate from different goroutine

			checkCh <- "third (never get here)"
		}()
	})
	val := <-valch
	if val != nil {
		t.Errorf("Expected val to be %v, but it was %v", nil, val)
	}

	exp := "first (just started)"
	if check := <-checkCh; check != exp {
		t.Errorf("Expected %q but got %q", exp, check)
	}

	exp = "second (in goroutine)"
	if check := <-checkCh; check != exp {
		t.Errorf("Expected %q but got %q", exp, check)
	}

	select {
	case check := <-checkCh:
		t.Errorf("Expected channel to be closed but got %q", check)
	default: // ok
	}
}

func Test_TerminatePassesValue(t *testing.T) {
	exp := "hello"
	valch := Start(func(terminate TerminatorFn) {
		terminate(exp) // cause termination passing a value
	})
	val := <-valch
	if val != exp {
		t.Errorf("Expected val to be %v, but it was %v", exp, val)
	}
}

func Test_Empty(t *testing.T) {
	valch := Start(func(terminate TerminatorFn) {})
	val := <-valch
	if val != nil {
		t.Errorf("Expected val to be %v, but it was %v", nil, val)
	}
}
