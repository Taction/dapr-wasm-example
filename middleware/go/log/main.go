package main

import (
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

func main() {
	handler.HandleRequestFn = handleRequest
}

func handleRequest(req api.Request, resp api.Response) (next bool, reqCtx uint32) {
	handler.Host.Log(api.LogLevelInfo, "Wasm log handleRequest uri: "+req.GetURI())
	return true, 0
}
