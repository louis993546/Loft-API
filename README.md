# Loft API

## Ground rules

- Specification in OpenAPI 3.0.2
- Structure follows JSONAPI 1.0 STABLE

## Build and/or run

### Just GO

1. Make sure Go is setup correctly
2. `go get github.com/gorilla/mux`
3. build and run in the way you want
   - `go install` + `loft`
   - `go run main.go`
   - `go build` + `$GOPATH/bin/loft`

### Docker

1. Make sure Docker is up and running
2. `docker build --rm -f "Dockerfile" -t api-loft:latest .`
3. `docker run -it --rm --name my-running-app -p 8080:8080 api-loft:latest`