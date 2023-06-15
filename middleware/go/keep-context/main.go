package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.Host.EnableFeatures(api.FeatureBufferResponse)
	handler.HandleRequestFn = handleRequest
	handler.HandleResponseFn = handleResponse
}

var globalContext = make(map[uint32]string)
var contextCounter uint32

func handleResponse(reqCtx uint32, req api.Request, resp api.Response, isError bool) {
	handler.Host.Log(api.LogLevelInfo, "handleResponse get uri from context "+globalContext[reqCtx])

	// Serve response
	resp.Headers().Set("Content-Type", "text/plain")
	resp.SetStatusCode(200)
	resp.Body().WriteString("Hello ")
	resp.Body().WriteString(globalContext[reqCtx])
	return
}

// handleRequest serves a static response from the Dapr sidecar.
func handleRequest(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	contextCounter++

	globalContext[contextCounter] = req.GetURI()
	handler.Host.Log(api.LogLevelInfo, "handleRequest uri: "+req.GetURI())
	return true, contextCounter // continue to execute the next middleware
}
