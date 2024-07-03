#!/usr/bin/env sh


if ! which docker >/dev/null 2>&1 ; then
  echo "Docker needs to be installed to build the sample"
  exit 1
fi

echo "Check for Docker daemon..."
if ! docker info >/dev/null 2>&1 ; then
  echo "Docker daemon needs to be running to build the sample"
  exit 2
fi

echo "Check for Docker daemon..."
if [ "$(basename "$(pwd)")" != "samples" ] ; then
  echo "This script must be run from the './sample' subdirectory off the readme-merge repository root."
  exit 3
fi

# Build the Docker image from the parent directory
docker build -t readme-merge:sample ../

# Run the Docker container with three parameters and remove it after execution
docker run --rm readme-merge:sample \
  -v "$(pwd)/..:/app/repo" \
  ./repo/md/_index.md ./repo/. no_commit

# Show
less ./repo/README.md

