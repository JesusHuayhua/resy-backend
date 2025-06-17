package internal

type Filter struct {
	Key   string `json:"key"`
	Value string `json:"value,omitempty"`
}

type StatusCode int

const (
	InProgress StatusCode = iota + 1
	Busy
	Halted
	Error
)
