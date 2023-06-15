package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleRequestFn = handleRequest
}

// handleRequest serves a static response from the Dapr sidecar.
func handleRequest(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	resp.Headers().Set("Content-Type", "text/plain")
	resp.Body().WriteString("Hello, this is wasm middleware, you are requesting: " + req.GetURI())
	return false, 0 // do not execute next middleware
}
