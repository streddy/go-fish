package structs

import "strings"

type TackleBox struct {
	Url, Method string
	MinLatency, MaxLatency float64
	RequestBodyHeader strings.Reader
}