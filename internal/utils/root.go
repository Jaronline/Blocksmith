package utils

// ReturnsError Returns whether a function threw an error.
// Useful for chaining when handling a lot of errors.
func ReturnsError[D any](dest *D, err *error, f func() (D, error)) bool {
	*dest, *err = f()
	return *err == nil
}
