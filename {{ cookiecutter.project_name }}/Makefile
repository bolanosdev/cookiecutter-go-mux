VERSION:=$(shell git rev-list --count HEAD)-$(shell git rev-parse --short HEAD)
DATE?=$(shell date -u '+%Y-%m-%d-%H%M')

postgres:
	docker run -d --rm --name postgres-cookie \
  -p 5434:5432 \
  -e POSTGRES_USER=cookie_user \
  -e POSTGRES_PASSWORD=cookie_pwd \
	-e POSTGRES_DB=cookie_db \
  -d postgres

migrate_up:
	migrate -path db/migrations -database postgres://cookie_user:cookie_pwd@localhost:5434/cookie_db?sslmode=disable -verbose up
	
migrate_down:
	migrate -path db/migrations -database postgres://cookie_user:cookie_pwd@localhost:5434/cookie_db?sslmode=disable -verbose down

create_db:
	docker exec -it postgres-cookie createdb --username=cookie_user --owner=cookie_user cookie_db

drop_db:
	docker exec -it postgres-cookie dropdb --username=cookie_user cookie_db

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
