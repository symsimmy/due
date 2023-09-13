GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

.PHONY: init
# init env
init:
	go install github.com/dobyte/due@latest
	go install github.com/dobyte/due/network/ws@latest
	go install github.com/dobyte/due/registry/etcd@latest
	go install github.com/dobyte/due/locate/redis@latest
	go install github.com/dobyte/due/transport/grpc@latest