package io_manager

// IOManager is an interface that defines the methods for reading lines and writing result data.
// ReadLines reads lines from a source and returns them as a slice of strings and an error if any.
// WriteResult writes the result data to a destination. The data parameter can be of any type. It returns an error if any.
type IOManager interface {
	ReadLines() ([]string, error)
	WriteResult(data any) error
}
