GO_BUILD_ENV := GOOS=linux GOARCH=amd64
DOCKER_BUILD=$(shell pwd)/.docker_build
DOCKER_CMD=$(DOCKER_BUILD)/secretly-sender/cmd

$(DOCKER_CMD): clean
    mkdir -p $(DOCKER_BUILD)
    $(GO_BUILD_ENV) go build -v -o $(DOCKER_CMD) .

clean:
    rm -rf $(DOCKER_BUILD)

heroku: $(DOCKER_CMD)
    heroku container:push web
    heroku container:push worker

install:
    go install github.com/fajarnugroho/twitter-news-discord-bot/cmd/tndb

build:
    go build ./cmd/tndb