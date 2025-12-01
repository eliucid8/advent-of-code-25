.PHONY: all build clean

all: build

build:
	go build ./...

clean:
	rm -f build/day*
