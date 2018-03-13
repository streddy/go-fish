package structs

import (
	"time"
)

// Response is a struct that stores information about an http response to a request
type Response struct {
	Duration   int64
	Error      bool
	Size       int64
	Start      time.Time
	StatusCode int
}
