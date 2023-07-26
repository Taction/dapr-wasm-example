package main

import (
	"encoding/json"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler"
	"github.com/http-wasm/http-wasm-guest-tinygo/handler/api"
)

type Config struct {
	Environment string `json:"environment"`
}

func main() {
	// config could be in any format depending on how the wasm code parses it
	configBytes := handler.Host.GetConfig()
	config := Config{}
	json.Unmarshal(configBytes, &config)
	handler.Host.Log(api.LogLevelInfo, "Config environment: "+config.Environment)
}
