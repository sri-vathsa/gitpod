#!/bin/bash

DIR=$(cd $(dirname "${BASH_SOURCE}") && pwd -P)

go install github.com/golang/protobuf/protoc-gen-go@v1.4.3
go install google.golang.org/protobuf/runtime/protoimpl@v1.25.0
go install google.golang.org/protobuf/runtime/protoimpl@v1.25.0

protoc -I. -I.. --go_out=plugins=grpc:. imgbuilder.proto
mv github.com/gitpod-io/gitpod/image-builder/api/* go
rm -rf github.com

go install github.com/golang/mock/mockgen@v1.5.0

cd go
# source mode does not always work for gRPC: see https://github.com/golang/mock/pull/163
mockgen -package mock . ImageBuilderClient,ImageBuilder_BuildClient,ImageBuilder_LogsClient > mock/mock.go
cd ..

cd typescript
export PATH=$DIR/typescript/node_modules/.bin:$PATH
protoc --plugin=protoc-gen-grpc=`which grpc_tools_node_protoc_plugin` --js_out=import_style=commonjs,binary:src --grpc_out=src -I.. -I../../ ../*.proto
protoc --plugin=protoc-gen-ts=`which protoc-gen-ts` --ts_out=src -I /usr/lib/protoc/include -I .. -I../../ ../*.proto
cd src
node $DIR/../content-service-api/typescript/patch-grpc-js.ts
cd ../..