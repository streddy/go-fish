package structs

import (
	"net/http"
)

// TackleBox is a struct that contains the initial traffic information
type TackleBox struct {
	Transport  *http.Transport
	Frequency  float64
	Routes     []Route
	MinLatency string
	MaxLatency string
    DropFreq   float64
    UseTrans   bool
}
