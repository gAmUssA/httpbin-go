#!/usr/bin/env bash

docker buildx create --name mybuilder --use
docker buildx build --platform linux/amd64,linux/arm64 -t gamussa/httpbin-go --push .