#!/bin/bash
set -euo pipefail

# Build the Docker image
docker build -t flock-docker-test .

# Run the Docker image 10 times
for i in {1..10}
do
  (
    docker run --rm --name="flock-docker-test-$i" flock-docker-test \
    | sed "s/^/[flock-docker-test-${i}] /"
  ) &
done

# Wait for all background jobs to finish
wait
