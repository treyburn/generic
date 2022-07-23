SHELL := /bin/zsh

.PHONY: test
test:
	go test -race ./...