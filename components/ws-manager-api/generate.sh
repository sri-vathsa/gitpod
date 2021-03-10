#!/bin/bash

set -e

go get github.com/golang/protobuf/protoc-gen-go@v1.4.3
go get google.golang.org/protobuf/reflect/protoreflect@v1.25.0
go get google.golang.org/protobuf/runtime/protoimpl@v1.25.0

protoc -I. -I.. --go_out=plugins=grpc:. core.proto
mv github.com/gitpod-io/gitpod/ws-manager/api/* go
rm -rf github.com

go get github.com/golang/mock/mockgen@latest
cd go
# source mode does not always work for gRPC: see https://github.com/golang/mock/pull/163
mockgen -package mock github.com/gitpod-io/gitpod/ws-manager/api WorkspaceManagerClient,WorkspaceManager_SubscribeClient > mock/mock.go
cd ..

# generate typescript client
cd typescript
rm src/core_*.ts src/core_*.js
export PATH=$(yarn bin):$PATH
protoc --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` --js_out=import_style=commonjs,binary:src --grpc_out=src -I.. -I../../ ../core.proto
protoc --plugin=protoc-gen-ts=`which protoc-gen-ts` --ts_out=src -I /usr/lib/protoc/include -I .. -I../../ ../core.proto
cd src
node ../../../content-service-api/typescript/patch-grpc-js.ts
cd ../..