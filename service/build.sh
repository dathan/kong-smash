#!/bin/bash

GOOS=linux go build .
docker build -t kong-fake-service .
docker run -p 8282:8282 --name kong-fake-service -t kong-fake-service
