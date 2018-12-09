# Loft API

## Ground rules

- Specification in OpenAPI 3.0.2
- Structure follows JSONAPI 1.0 STABLE

## Build and/or run

### Just GO

This project is using go module from 1.10

1. Make sure Go is setup correctly
2. build and run in the way you want
   - `go install` + `loft`
   - `go run main.go`
   - `go build` + `loft`

### Docker

1. Make sure Docker is up and running
2. `docker build --rm -f "Dockerfile" -t api-loft:latest .`
   - Copy necessary files into the container
   - `go build` will trigger modules to `go get` dependencies specified in `go.mod`
3. `docker run -it --rm -p 8080:8080 api-loft:latest`
   - Execute `CMD` specified in the `Dockerfile`