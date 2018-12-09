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

1. Make sure you have go 1.11.2 or newer
2. `go build`
3. `PORT={port of your choice} ./loft`
4. Open browser, and it should be up and running

### Docker

1. Make sure Docker is up and running
2. `docker build --rm -t api-loft:latest .`
   - Copy necessary files into the container
   - `go build` will trigger modules to `go get` dependencies specified in `go.mod`
3. `docker run -it --rm -p 8080:8080 -e PORT=8080 api-loft:latest`
   - Execute `CMD` specified in the `Dockerfile`

## Database

At some point this section will be integrated into the build section, but right now these 2 parts are not connected.

1. [Install psql to your machine](https://www.postgresql.org/download/)
2. Setup user?
3. Init schema

## Publish image

1. `docker login` if necessary
2. build the docker image (see [above](#docker))
3. `docker tag api-loft:{version name} louis993546/loft:{version name}`
4. `docker push louis993546/loft:{version name}`
   This will take a while

For more info, check [the official documentation](https://docs.docker.com/docker-cloud/builds/push-images/)