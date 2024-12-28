_PHONY: build

default: build

build:
	@GOOS=js GOARCH=wasm go build -ldflags "-s -w" -o game.wasm

dev: build
	@air
