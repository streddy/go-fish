package gofish

import (
	"github.com/streddy/go-fish/structs"

	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"sync"
	"time"
)

// Freestyle hits routes and their dependencies in a random order under specified drops/latencies
func Freestyle(trafficParams structs.TackleBox, responseChannels []chan *structs.Response, done *bool) {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Send a request every 1/Frequency seconds (at fastest)
	ticker := time.NewTicker(time.Second / time.Duration(trafficParams.Frequency))
	for range ticker.C {
		// if caller has decided to stop sending traffic, break
		if done {
			break
		}

		// simulate packet drop
		dropPacket := determineDrop(trafficParams.DropFreq, random)
		if dropPacket {
			// pick a random route
			index := random.Intn(len(trafficParams.Routes))
			route := trafficParams.Routes[index]

			// construct BaitBox struct for this route
			requestInfo := structs.BaitBox{
				Route:      route,
				Transport:  trafficParams.Transport,
				MinLatency: trafficParams.MinLatency,
				MaxLatency: trafficParams.MaxLatency,
			}

            // prepare request-specific fields
            // dependent on whether we are using Transport
            if trafficParams.UseTrans {
                request := prepareBaitTransport(requestInfo)
            } else {

            }
			request := prepareBait(requestInfo)

			// send request
			requestStart := time.Now()
			response := castReel(request)

			// hit all of the route's dependencies
			for _, dependency := range route.MandatoryDependencies {
				requestInfo.Route = dependency
				request := prepareBait(requestInfo)
				castReel(request)
			}

			// fill remaining response fields and place response in channel
			response.Start = requestStart
			response.Duration = time.Since(requestStart).Nanoseconds()
			responseChannels[index] <- response
		}
	}
}

// Warmup is used to warmup a cache on a specific route. It does not handle responses
func Warmup(trafficParams structs.TackleBox, done *bool) {
	// Create request
	route := trafficParams.Routes[0]
	requestBodyReader := strings.NewReader(route.RequestBody)
	request, _ := http.NewRequest(route.Method, route.Url, requestBodyReader)
	
	// Create requests for route's dependencies
	var dependencyList []*http.Request
	for i, dependency := range route.MandatoryDependencies {
		depReqBodyReader := strings.NewReader(dependency.RequestBody)	
		depReq, _ := http.NewRequest(dependency.Method, dependency.Url, depReqBodyReader)
		dependencyList[i] := depReq
	}

	// Send a request every 1/Frequency seconds (at fastest)
	ticker := time.NewTicker(time.Second / time.Duration(tps.Frequency))
	for range ticker.C {
		// If caller has decided to stop sending traffic, break
		if done {
			break
		}

		httpResponse, err := trafficParams.Transport.RoundTrip(request)

		// hit all the described dependencies in routes.json
		for _, depReq := range dependencyList {
			httpResponse, err := trafficParams.Transport.RoundTrip(depReq)
		}
	}
}

// OneFish sends a single request and ignores the response. Useful for calibrating timers
func OneFish(trafficParams structs.TackleBox) {
	route := trafficParams.Routes[0]
	requestBodyReader := strings.NewReader(route.RequestBody)
	request, _ := http.NewRequest(route.Method, route.Url, requestBodyReader)
	request.Header.set(route.Headers)

	transport.RoundTrip(request)
}

// HELPER FUNCTIONS

// Determine whether a packet should be dropped
func determineDrop(dropFreq float64, random *Rand) bool {
	dropNum := 0
	// generate a random value between 0 and 1000
	for i := 0; i < 1000; i++ {
		dropNum += random.Intn(2)
	}

	return dropNum/1000 <= dropFreq
}

// Prepares a Bait struct that will be used in sending the request as a Transport object
func prepareBait(requestInfo structs.BaitBox) *structs.Bait {
	// construct request
	route = requestInfo.Route
	requestBodyReader := strings.NewReader(route.RequestBody)
	request, _ := http.NewRequest(route.Method, route.Url, requestBodyReader)

	// determine a random latency for this request
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	minLat, _ := time.ParseDuration(requestInfo.MinLatency)
	maxLat, _ := time.ParseDuration(requestInfo.MaxLatency)
	timeInterval := random.Int63n(time.Nanoseconds(maxLat) -
		time.Nanoseconds(minLat))
	+time.Nanoseconds(minLat)

	// split incoming header string by \n and build header pairs
	headerPairs := strings.Split(route.Headers, "\n")
	for i := range headerPairs {
		split := strings.SplitN(headerPairs[i], ":", 2)
		if len(split) == 2 {
			request.Header.Set(split[0], split[1])
		}
	}

	bait := &structs.Bait{
		Transport: requestInfo.Transport,
		Request:   request,
		Latency:   timeInterval,
	}

	return bait
}

// Sends a request and introduces traffic latency
func castReel(request *structs.Bait) *structs.Response {
	// introduce latency
	time.sleep(*request.Latency)

	// send request and handle response
	httpResponse, err := *request.Transport.RoundTrip(*request.Request)
	response := reelIn(httpResponse, err != nil)

	return response
}

// Parses the response and returns it to caller
func reelIn(httpResponse *http.Response, err bool) *structs.Response {
	response := &structs.Response{}

	if err {
		response.Error = true
	} else {
		if httpResponse.ContentLength < 0 { // -1 if the length is unknown
			content, err := ioutil.ReadAll(httpResponse.Body)
			if err == nil {
				response.Size = int64(len(content))
			}
		} else {
			response.Size = httpResponse.ContentLength
		}
		response.StatusCode = httpResponse.StatusCode
		defer httpResponse.Body.Close()
	}

	return response
}
