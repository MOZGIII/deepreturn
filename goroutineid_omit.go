// +build nogoroutinechecks

package deepreturn

type goroutineID struct{}

func getGoroutineID() goroutineID {
	return goroutineID(struct{}{})
}

func checkGoroutineID(other goroutineID, expectMatch bool) {}
