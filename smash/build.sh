#!/bin/bash

GOOS=linux go build .
docker build -t kong-smash .
docker rm kong-smash
docker run --net="host" --name kong-smash -t kong-smash /app/smash -concurrency 8
