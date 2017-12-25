package deepreturn

import (
	"runtime"

	"github.com/MOZGIII/deepreturn/internal/goroutinelock"
)

// Terminator provides means to terminate the execution.
type Terminator interface {
	Terminate(interface{})
}

// TerminatorFn is a func wrapper for Terminator interface.
type TerminatorFn func(interface{})

var _ Terminator = TerminatorFn(nil)

// Terminate will terminate the execution of a goroutine.
func (f TerminatorFn) Terminate(val interface{}) { f(val) }

type terminator struct {
	goroutineLock goroutinelock.GoroutineLock
	valch         chan interface{}
}

func (t *terminator) terminate(val interface{}) {
	t.goroutineLock.Check()
	t.valch <- val
	runtime.Goexit()
	*(*int)(nil) = 0 // not reached
}

// Start starts an execution.
func Start(routine func(terminate TerminatorFn)) <-chan interface{} {
	valch := make(chan interface{}, 1)
	go func() {
		glock := goroutinelock.New()
		ter := &terminator{glock, valch}
		defer close(valch)
		routine(TerminatorFn(ter.terminate))
	}()
	return valch
}
