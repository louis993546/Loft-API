# Loft API

## Ground rules

- Specification in OpenAPI 3.0.2
- Structure follows JSONAPI 1.0 STABLE

## Structure

### Goal (v1.0.0)

1. Database
   - PostgreSQL
2. Auth service
   - TBC
3. Router
   - `main.go`
   - Powered by mux
   - has basically no business logic
   - It should make scaling a lot easier
4. {each endpoint}

### Currently

1. Database
   - Either a docker image, or Heroku pgSQL
2. Service
   - Router + Auth + every endpoint = `main.go`

## Build and/or run

### Just GO

This project is using go module from 1.11

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