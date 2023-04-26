#!/bin/bash

app="$1" # app name
base_env_file="$2"
env_file="$3" # .test.env

base_env_file=${base_env_file:-".env.development"}

cd "./apps/$app"

# Read variables from env files and pass them to the 'env' command
godotenv -f "../../$base_env_file","../../$env_file" go test -v -count=1 -p=8 -parallel=8 -race ./...
