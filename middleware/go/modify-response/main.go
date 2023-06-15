package main

import (
	"bytes"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.Host.EnableFeatures(api.FeatureBufferResponse) // change response must enable buffer response feature
	handler.HandleResponseFn = handleResponse
}

// handleResponse can modify the response
func handleResponse(reqCtx uint32, req api.Request, resp api.Response, isError bool) {
	// get the original response body
	buf := &bytes.Buffer{}
	resp.Body().WriteTo(buf)
	resp.Headers().Set("Content-Type", "text/plain")
	// modify the response body, adding a prefix "hello"
	resp.Body().WriteString("hello ")
	resp.Body().Write(buf.Bytes())
}
