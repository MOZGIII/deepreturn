// +build !nogoroutinechecks

package deepreturn

import (
	"fmt"

	"github.com/huandu/goroutine"
)

type goroutineID int64

func getGoroutineID() goroutineID {
	return goroutineID(goroutine.GoroutineId())
}

func checkGoroutineID(other goroutineID, expectMatch bool) {
	currentID := getGoroutineID()
	match := other == currentID
	if expectMatch {
		if !match {
			panic(fmt.Errorf("execution: goroutine ID of the current routine does not match the expected ID"))
		}
	} else {
		if match {
			panic(fmt.Errorf("execution: goroutine ID of the current routine does match the expected ID, which in not allowed"))
		}
	}
}
