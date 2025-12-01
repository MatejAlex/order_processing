#!/usr/bin/env zsh

GOOS=linux GOARCH=amd64 go build .

docker build . -t matejalex/orderprocessor:latest

docker run -d --name orderprocessor -p 8080:8080 matejalex/orderprocessor:latest