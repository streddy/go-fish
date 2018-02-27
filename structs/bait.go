package structs

import (
    "net/http"
    "time"
)

type Bait struct {
    Transport *http.Transport
    Request http.Request
    Latency int64
}
