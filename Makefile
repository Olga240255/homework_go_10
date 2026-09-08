SHELL := /bin/bash
GO ?= go
PACKAGES := ./...

.PHONY: help run build compile fmt fmt-check vet test test-store test-http test-integration test-race check-start ci clean

help:
	@echo "Домашняя работа: Notes API, 40 шагов"
	@echo "  make check-start      - проверить стартовые заглушки без функциональных тестов"
	@echo "  make test-store       - проверить модель и store"
	@echo "  make test-http        - проверить router и handlers"
	@echo "  make test             - запустить все тесты; до решения они красные"
	@echo "  make ci               - полная проверка готового решения"

run:
	$(GO) run ./cmd/api

build:
	@mkdir -p bin
	$(GO) build -buildvcs=false -o bin/notes-api ./cmd/api

compile:
	$(GO) test -run '^$$' $(PACKAGES)

fmt:
	gofmt -w $$(find . -name '*.go' -not -path './bin/*')

fmt-check:
	@files="$$(gofmt -l $$(find . -name '*.go' -not -path './bin/*'))"; \
	if [ -n "$$files" ]; then echo "Найдены неотформатированные Go-файлы:"; echo "$$files"; exit 1; fi

vet:
	$(GO) vet $(PACKAGES)

test:
	$(GO) test $(PACKAGES)

test-store:
	$(GO) test ./internal/note

test-http:
	$(GO) test ./internal/httpapi

test-integration:
	$(GO) test ./test/integration/...

test-race:
	$(GO) test -race $(PACKAGES)

check-start: compile fmt-check vet build

ci: compile fmt-check vet test test-race build

clean:
	rm -f bin/notes-api
	@rmdir bin 2>/dev/null || true
