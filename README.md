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

### Docker

| Action             | bash command                                                                                   |
|--------------------|------------------------------------------------------------------------------------------------|
| Build docker image | `docker build --rm -t api-loft:latest .`                                                       |
| Run bash in image  | `docker run -it api-loft:latest bash`                                                          |
| Run `CMD` in image | `docker run -it --rm -p 8080:8080 api-loft:latest` (TODO: probably missing a couple variables) |

### Heroku

Currently there are 1 pipeline + 4 heroku apps

| App                   | Purpose                                                        | Auto deploy |
|-----------------------|----------------------------------------------------------------|-------------|
| `loft-api-production` | Production                                                     | `master`    |
| `loft-api-staging`    | For the next release                                           | N/A         |
| `loft-api-develop`    | For the develop branch                                         | `develop`   |
| `loft-api-testbed`    | A 100% random branch that can just be push for testing purpose | N/A         |

And if you want to manually deploy something, you can do it in the heroku dashboard of the app, or do it in git

1. Add your design heroku app into your local git repo (`heroku git:remote -a <heroku app name>`)
2. `git push heroku <local branch name>:master`

Make sure if you do this, you don't mess up the 4 apps and deploy something to the wrong app.

## CI/CD

### Travis CI

In Progress

### CircleCI

TBD

### GitLab CI

TBD