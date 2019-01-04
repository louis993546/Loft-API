workflow "Docker: Build + Tag + Publish" {
  on = "push"
  resolves = ["Deploy to Zeit", "Deploy to Heroku (develop)"]
}

action "`develop` only" {
  uses = "actions/bin/filter@b2bea07"
  args = "branch develop"
}

action "Build image" {
  uses = "actions/docker/cli@76ff57a"
  needs = ["`develop` only"]
  args = "build --rm ."
}

action "Docker Tag this image" {
  uses = "actions/docker/tag@76ff57a"
  needs = ["Build image"]
}

action "Login to Docker Hub Registry" {
  uses = "actions/docker/login@76ff57a"
  needs = ["Docker Tag this image"]
}

action "Publish image to Docker Hub" {
  uses = "actions/docker/cli@76ff57a"
  needs = ["Login to Docker Hub Registry"]
}

action "Deploy to Heroku (develop)" {
  uses = "actions/heroku@6db8f1c"
  needs = ["Publish image to Docker Hub"]
}

action "Deploy to Zeit" {
  uses = "actions/zeit-now@9fe84d5"
  needs = ["Publish image to Docker Hub"]
}
