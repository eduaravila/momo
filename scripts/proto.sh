#!/bin/bash

readonly app="$1"
readonly service="$2"

protoc \
    "--proto_path=apps/$app/api/protobuf" "$service.proto" \
    "--go_out=packages/genproto/$service" --go_opt=paths=source_relative \
    --go-grpc_opt=require_unimplemented_servers=false \
    "--go-grpc_out=packages/genproto/$service" --go-grpc_opt=paths=source_relative \

    
