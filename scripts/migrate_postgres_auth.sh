#!/bin/bash
set -e
#!/bin/bash

readonly REPOSITORY_NAME="$1"
readonly DOCKER_FILE="$2"
readonly IMAGE_NAME="${REPOSITORY_NAME}:latest"
readonly SHOULD_BUILD="$3"
readonly DATABASE_URL="$4"

if [[ "$SHOULD_BUILD" == "true" ]]; then
  # Build the image
  docker build -t "${IMAGE_NAME}" -f "${DOCKER_FILE}" packages/postgres
fi

# Run the container
docker run --rm \
  --env DATABASE_URL=${DATABASE_URL} \
  --network postgres \
  "${IMAGE_NAME}" \
  sh -c "goose -dir ./migrations postgres 'postgresql://postgres:postgres@postgres:5432/postgres' up"
