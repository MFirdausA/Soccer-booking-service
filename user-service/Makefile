GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

## Live reload:
watch-prepare: ## Install the tools required for the watch command
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh

watch: ## Run the service with hot reload
	bin/air

## Build:
build: ## Build the service
	go build -o user-service

## Docker:
docker-compose: ## Start the service in docker
	docker-compose up -d --build --force-recreate

docker-build: ## Build the Docker image with a specified tag
	@echo Building Docker image...
ifndef tag
	$(error Please specify the tag parameter, e.g., make docker-build tag=1.0.0)
endif
	docker build --platform linux/amd64 -t frdausa/user-service:$(tag) .
	@echo Docker image built with tag $(tag)

docker-push: ## Push the Docker image with a specified tag
	@echo Pushing Docker image...
ifndef tag
	$(error Please specify the tag parameter, e.g., make docker-push tag=1.0.0)
endif
	docker push frdausa/user-service:$(tag)
	@echo Docker image pushed with tag $(tag)