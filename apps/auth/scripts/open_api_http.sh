#!/bin/bash

readonly service = "$1"
readonly output_dir = "$2"
readonly package = "$3"

oapi-codegen --package="$package" -generate chi-server "$service".yaml > api.gen.go
oapi-codegen --package="$package" -generate types auth.yaml > api_types.gen.go