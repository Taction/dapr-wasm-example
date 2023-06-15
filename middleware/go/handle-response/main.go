package main

import (
	"fmt"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleResponseFn = handleResponse
}

// handleResponse can get the request and response information
func handleResponse(reqCtx uint32, req api.Request, resp api.Response, isError bool) {
	//handler.Host.Log(api.LogLevelInfo, "handleResponse from wasm side")
	handler.Host.Log(api.LogLevelInfo, fmt.Sprintf("HandleResponse uri: %s, response status code: %d", req.GetURI(), resp.GetStatusCode()))
	return
}
