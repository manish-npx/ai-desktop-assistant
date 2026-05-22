.PHONY: help install dev build test clean

help:
	@echo "AI Desktop Assistant - Build Commands"
	@echo ""
	@echo "Available targets:"
	@echo "  install        Install dependencies (Go + Node)"
	@echo "  dev            Run development mode (hot reload)"
	@echo "  build          Build production binary"
	@echo "  build-win      Build Windows binary"
	@echo "  build-mac      Build macOS binary"
	@echo "  build-linux    Build Linux binary"
	@echo "  test           Run all tests"
	@echo "  test-go        Run Go tests"
	@echo "  test-react     Run React tests"
	@echo "  migrate        Run database migrations"
	@echo "  clean          Clean build artifacts"

install:
	go mod download
	go mod verify
	cd frontend && npm install && cd ..

dev:
	wails dev

build:
	wails build

build-win:
	wails build -platform windows/amd64

build-mac:
	wails build -platform darwin/universal

build-linux:
	wails build -platform linux/amd64

test:
	go test -v ./...
	cd frontend && npm test -- --run && cd ..

test-go:
	go test -v ./...

test-react:
	cd frontend && npm test -- --run && cd ..

migrate:
	@echo "Running database migrations..."
	@sh scripts/migrate.sh

clean:
	rm -rf build/
	rm -rf frontend/dist/
	go clean

fmt:
	go fmt ./...
	cd frontend && npm run format && cd ..

lint:
	go vet ./...
	cd frontend && npm run lint && cd ..
