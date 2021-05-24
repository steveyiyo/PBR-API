#!/bin/bash

echo "Starting for build!"

echo "Linux amd64"
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o build/PBR-API_Linux