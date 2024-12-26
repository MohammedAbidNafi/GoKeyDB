.PHONY: build install

build:
	go build -o gokeydb

install: build
	sudo mv gokeydb /usr/local/bin/
