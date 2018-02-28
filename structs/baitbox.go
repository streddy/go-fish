package structs

import {
	"net/http"
)

// BaitBox is a struct used as an intermediary in creating a request (Bait struct)
type BaitBox struct {
	Route 		Route
	Transport 	*http.Transport
	MinLatency	string
	MaxLatency 	string
}
