.PHONY: all format lint clean

GOCMD=go
LDFLAGS="-s -w ${LDFLAGS_OPT}"
FUNCTIONS=migrate findAll create findById delete update

all: ## Build every lambda
	make ${FUNCTIONS}

.PHONY: $(FUNCTIONS)
$(FUNCTIONS): ## Create zip file
	GOOS=linux GOARCH=amd64 ${GOCMD} build -tags lambda.norpc -o bin/$@/bootstrap cmd/$@/main.go 
	zip -j bin/$@/function.zip bin/$@/bootstrap

format: ## Format code
	${GOCMD} fmt ./...

lint: ## Run linter
	golangci-lint run

clean: ## Cleanup build dir
	rm -r bin/

local:
	sam local start-api

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
