#!/bin/bash

set -e

DIR=$(cd $(dirname "${BASH_SOURCE}") && pwd -P)


go get github.com/golang/protobuf/protoc-gen-go@v1.4.3

protoc -I. -I.. --go_out=plugins=grpc:. imagespec.proto
protoc -I. -I.. --go_out=plugins=grpc:. provider.proto

mv github.com/gitpod-io/gitpod/registry-facade/api/* go
rm -rf github.com
