.DEFAULT_GOAL := help

.PHONY: help lint build test check

help: ## Elérhető parancsok
	@grep -hE '^[a-z-]+:.*## ' $(MAKEFILE_LIST) | awk -F':.*## ' '{printf "  make %-8s %s\n", $$1, $$2}'

lint: ## gofmt + go vet
	@out=$$(gofmt -l .); if [ -n "$$out" ]; then echo "Formázatlan fájlok:"; echo "$$out"; exit 1; fi
	go vet ./...

build: ## Minden csomag lefordul
	go build ./...

test: ## Tesztek
	go test ./... -count=1

check: lint build test ## Amit a CI futtat
