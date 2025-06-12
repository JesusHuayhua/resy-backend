package internal

type StatusCode int

const (
	InProgress StatusCode = iota + 1
	Busy
	Halted
	Error
)
