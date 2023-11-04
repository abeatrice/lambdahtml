.PHONY: help
help: ## This help
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

null:
	help

# CLI Arguments
ARGS = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`
EVENT = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

CMD :=
GO := docker run -it \
	-v `pwd`:/app/ \
	-v ~/.aws:/root/.aws \
	--user $(id -u):$(id -g) \
	--env GOOS=linux \
	--env GOARCH=amd64 \
	-w /app/ golang:1.19.0

.PHONY: build
build: ## Build the app
	go build -o main main.go \
		&& zip -o main.zip main

.PHONY: invoke
invoke: ## invoke local function: make invoke EVENT="test_events/get-root.json"
	go build -o main main.go \
	&& cat $(EVENT) | docker -- run --rm -v `pwd`:/var/task:ro,delegated \
		--env-file .env \
		-i -e DOCKER_LAMBDA_USE_STDIN=1 lambci/lambda:go1.x main

.PHONY: run
run: ## run local server
	docker-compose up --build
