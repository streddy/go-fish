package gofish

import (
    "github.com/streddy/go-fish/structs"

    "math/rand"
    "net/http"
    "fmt"
    "strings"
    "time"
)

func PrepareBait(requestInfo structs.TackleBox) *structs.Bait {
    // construct request
    request, _ := http.NewRequest(requestInfo.Method, requestInfo.Url,
                                  requestInfo.RequestBodyReader)
    
    // determine a random latency for this request
    minLat, _ := time.ParseDuration(requestInfo.MinLatency)
    maxLat, _ := time.ParseDuration(requestInfo.MaxLatency)
    timeInterval := rand.Int63n(time.Nanoseconds(maxLat) - 
                                time.Nanoseconds(minLat)) 
                    + time.Nanoseconds(minLat)

    bait := &structs.Bait{
        Transport:  requestInfo.Transport,
        Request:    request,
        Latency:    timeInterval,
    }

    return bait
}

func GoFish(request *structs.Bait) *http.Response {
    time.sleep(*request.Latency)
    httpResponse, err := *request.Transport.RoundTrip(*request.Request)
    return httpResponse
} 