package deepreturn

// WaitForAny blocks execution until some value or nil is returned.
func WaitForAny(valch <-chan interface{}) interface{} { return <-valch }

// WaitForErr blocks execution until an error (or nil) is returned.
func WaitForErr(valch <-chan interface{}) error {
	val := WaitForAny(valch)
	err, ok := val.(error)
	if ok {
		return err
	}
	return nil
}
