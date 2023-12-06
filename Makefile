# Makefile

all: build

build:
	mkdir -p ./dist
	go build -o ./dist

clean:
	rm -rf ./dist

clean_build: clean build

.PHONY: all build clean clean_build
