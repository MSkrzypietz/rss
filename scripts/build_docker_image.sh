#!/bin/bash

if [ $# -eq 0 ]; then
  echo "Error: No argument provided."
  echo "Usage: $0 <version>"
  exit 1
fi

versionRegex="^[0-9]+\.[0-9]+\.[0-9]+$"
if [[ ! $1 =~ $versionRegex ]]; then
  echo "Error: Argument is not a valid version number (expected format: X.Y.Z)"
  exit 1
fi

sudo docker buildx build --platform linux/amd64,linux/arm64 -t mskrzypietz/rss:$1 --push .