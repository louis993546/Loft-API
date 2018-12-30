# Loft GraphQL API

- Powered by gqlgen
- PostgreSQL database
- Hosted on Heroku
- Package in Docker image

## Run locally

### From source code

1. Clone the repository to your $GOPATH
2. (`dep ensure` I think?)
3. `go run server/server.go`
   - Make sure you/your system provides all the necessary environment variables, including:
     - `PORT`
     - `DATABASE_URL`
     - (optional) `DOMAIN`

### From docker image @ Docker Hub

Once the git process has been setup correctly, publishing image will be automated. Right now Docker Hub only has a very old version of the image

## Deploy

Right now it's still through manual `git push heroku master`, and it will be changed to use webhook in heroku ASAP

### Build docker image

`docker build --rm -t api-loft:latest .`

### Run bash in image

`docker run -it api-loft:latest bash`

### Run `CMD` in image

`docker run -it --rm -p 8080:8080 -e PORT=8080 api-loft:latest` (TODO: probably missing a couple variables)