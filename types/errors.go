package types

import "fmt"

type RequestError struct {
	Err        error
	StatusCode int
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("err %v", r.Err)
}

type StandardError struct {
	Message string
}
