package deepreturn

import (
	"runtime"
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
	goroutineID goroutineID
	valch       chan interface{}
}

func (t *terminator) terminate(val interface{}) {
	checkGoroutineID(t.goroutineID, true)
	t.valch <- val
	runtime.Goexit()
	*(*int)(nil) = 0 // not reached
}

// Start starts an execution.
func Start(routine func(terminate TerminatorFn)) <-chan interface{} {
	valch := make(chan interface{}, 1)
	go func() {
		gid := getGoroutineID()
		ter := &terminator{gid, valch}
		defer close(valch)
		routine(TerminatorFn(ter.terminate))
	}()
	return valch
}
