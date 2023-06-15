build:
	@echo "Building Wasm"
	@tinygo build -o main.wasm -scheduler=none --no-debug -target=wasi main.go

buildwithgo:
	@echo "Building Wasm"
	@GOOS=wasip1 GOARCH=wasm gotip build -o main.wasm main.go