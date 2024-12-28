_PHONY: build watch

default: watch

build:
	go build -ldflags "-s -w"

watch:
	@air
