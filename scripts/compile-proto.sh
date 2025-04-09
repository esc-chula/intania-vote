#!/bin/bash

# Generate Go gRPC code
protoc \
  --go_opt=module=github.com/esc-chula/intania-vote \
  --go_out=./ \
  --go-grpc_opt=module=github.com/esc-chula/intania-vote \
  --go-grpc_out=./ \
  --proto_path=./ \
  ./proto/**/*.proto

# Generate TypeScript gRPC code
protoc \
  --plugin=protoc-gen-ts_proto=./node_modules/.bin/protoc-gen-ts_proto \
  --ts_proto_out=./libs/grpc-ts \
  --ts_proto_opt=outputServices=grpc-js,env=node,esModuleInterop=true \
  --ts_proto_opt=exportCommonSymbols=false \
  --proto_path=./ \
  ./proto/**/*.proto

echo "Protocol Buffers generation completed!"
