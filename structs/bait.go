package structs

import (
    "net/http"
    "time"
)

// Bait is a struct containing a request and its traffic behavior
type Bait struct {
    Transport   *http.Transport
    Request     *http.Request
    Latency     int64
}
