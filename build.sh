#!/usr/bin/env bash

echo build start
cd api
go clean
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
docker build -t module-name .

cd ..

cd internal-api
go clean
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
docker build -t internal-module-name .

echo build end