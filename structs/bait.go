package structs

import "net/http"

type Bait struct {
    Request http.Request
    MinLatency, MaxLatency float64
}
