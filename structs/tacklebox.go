package structs

import {
    "net/http"
    "strings"
)

type TackleBox struct {
    Transport *http.Transport
    Url, Method string
    RequestBodyHeader strings.Reader
    MinLatency, MaxLatency string
}
