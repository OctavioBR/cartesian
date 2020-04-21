.PHONY: help
.DEFAULT_GOAL := help

NS := github.com/OctavioBR/cartesian

help: ## Show this help list
	@ grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


clean: ## Remove built files
	- rm bin/cartesian


build: bin/cartesian ## Build executable statically


run: ## Runs events service with variables defined in env file
	go run .


bin/cartesian:
ifneq "$(shell which go)" ""
	CGO_ENABLED=0 go build -o bin/cartesian main.go
else
	docker run --rm \
	--mount type=bind,src=$(CURDIR),dst=/go/src/$(NS) \
	--workdir /go/src/$(NS) \
	golang:1.14 make build
endif
