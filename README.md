# Golang FastHTTP API with PGX



## Requirements

* x86-64
* Linux/Unix
* [Golang](https://go.dev/)
* [Docker](https://www.docker.com/products/docker-desktop/)

## Startup

The script "up" creates our DB container, compiles and executes our binary:
```
1. docker-compose -f db/cars/docker-compose.yml up -d
2. go build -o golang_api src/main.go
3. ./golang_api
```

## Shutdown

The script "down" removes our DB container:
```
1. docker-compose -f db/cars/docker-compose.yml down
```