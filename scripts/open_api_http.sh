#!/bin/bash
# This sets the script to exit immediately if any command in the script fails with a non-zero exit code.
set -e

readonly app="$1"
readonly service="$2"
readonly package="$3"
readonly output_dir="$4"

oapi-codegen \
    -generate chi-server \
    -package "$package" \
    -o "apps/$app/$output_dir/openapi_${service}_api.gen.go" \
    "apps/$app/api/openapi/$service.yaml"

oapi-codegen \
    -generate types \
    -package "$package" \
    -o "apps/$app/$output_dir/openapi_${service}_types.gen.go" \
    "apps/$app/api/openapi/$service.yaml"

oapi-codegen \
    -generate client \
    -package "$service" \
    -o "packages/client/${service}/openapi_client.gen.go" \
    "apps/$app/api/openapi/$service.yaml"

oapi-codegen \
    -generate types \
    -package "$service" \
    -o "packages/client/${service}/openapi_types.gen.go" \
    "apps/$app/api/openapi/$service.yaml"