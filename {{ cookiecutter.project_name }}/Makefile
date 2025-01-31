VERSION:=$(shell git rev-list --count HEAD)-$(shell git rev-parse --short HEAD)
DATE?=$(shell date -u '+%Y-%m-%d-%H%M')

help:
	@echo 'Available commands:'
	@echo '* help            - Show this message'
	@echo '* lint            - Lint code'
	@echo '* test            - Run all tests'
	@echo '* unit            - Run unit tests'
	@echo '* integration     - Run integration tests'
	@echo '* cover           - Coverage report'
	@echo '* build           - Build binaries'
	@echo '* run             - Start the main file'
	@echo '* clean           - Clean built binaries'

build:
		@mkdir -p dist
		go build \
			-ldflags='-X "main.Version=$(VERSION)" -X "main.BuildTime=$(DATE)"' \
			-o dist/native/cookiecutter \
			./main.go
		@cp app.yaml ./dist/native/app.yaml
		GOOS=linux GOARCH=amd64 go build \
			-ldflags='-X "main.Version=$(VERSION)" -X "main.BuildTime=$(DATE)"' \
			-o dist/linux-amd64/cookiecutter \
			./main.go
		@cp app.yaml ./dist/linux-amd64/app.yaml
		@mkdir -p dist

clean:
		@rm -rf dist
		@rm -rf .cover

fmt:
	go fmt ./...

lint:
		@golangci-lint run

test: unit integration

run:
	go run -v main.go

unit:
	@echo 'Running unit tests...'
	@mkdir -p .cover
	go test -v -count=1 \
		-tags=!integration ./internal/... \
			-coverprofile .cover/cover.out

integration:
	@echo 'Running integration tests...'
	go test -v -count=1 \
		-tags=integration ./test/...

cover:
ifndef CI
	go tool cover -html .cover/cover.out
else
	go tool cover -html .cover/cover.out -o .cover/all.html
endif
