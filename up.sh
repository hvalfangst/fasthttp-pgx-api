#!/bin/sh

# Exits immediately if a command exits with a non-zero status
set -e

# Run 'docker-compose up' for creating our DB container
docker-compose -f db/cars/docker-compose.yml up -d

# Build the Go application
go build -o golang_api src/main.go

# Run the application
./golang_api