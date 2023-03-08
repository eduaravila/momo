#!/bin/bash
# This sets the script to exit immediately if any command in the script fails with a non-zero exit code.
set -e

readonly service = "$1"
readonly package = "$3"
readonly output_dir = "$3"


oapi-codegen  -generate chi-server -package "$package" -o "$output_dir/openapi_api.gen.go" api/openapi/"$service".yaml

oapi-codegen -generate types -package "$package" -o "$output_dir/openapi_types.gen.go" "api/openapi/$service".yaml