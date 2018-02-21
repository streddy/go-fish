package gofish

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

func PrepareBait(request_info structs.TackleBox) *structs.Bait {
	request, _ := http.NewRequest(request_info.Method, request_info.Url, request_info.RequestBodyReader)
	bait := &structs.Bait{
		Request:		request,
		MinLatency:	request_info.MinLatency,
		MaxLatency:	request_info.MaxLatency,
	}

	return bait
}
