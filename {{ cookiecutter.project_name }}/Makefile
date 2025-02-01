VERSION:=$(shell git rev-list --count HEAD)-$(shell git rev-parse --short HEAD)
DATE?=$(shell date -u '+%Y-%m-%d-%H%M')

POSTGRES_USER={{ cookiecutter.postgres_user }}
POSTGRES_PASSWORD={{ cookiecutter.postgres_password }}
POSTGRES_PORT={{ cookiecutter.postgres_port }}
POSTGRES_HOSTNAME={{ cookiecutter.postgres_hostname }}
POSTGRES_DB={{ cookiecutter.project_name }}_db
POSTGRES_CONNECTION_STRING=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@$(POSTGRES_HOSTNAME):$(POSTGRES_PORT)/$(POSTGRES_DB)?sslmode=disable

postgres:
	docker run -d --rm --name postgres  \
  -p $(POSTGRES_PORT):5432 \
  -e POSTGRES_USER=$(POSTGRES_USER) \
  -e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
	-d postgres:12-alpine

jaeger: 
	docker run -d --rm --name jaeger \
		-p 5778:5778 \
		-p 16686:16686 \
		-p 4317:4317 \
		-p 4318:4318 \
		-p 14250:14250 \
		-p 14268:14268 \
		-p 9411:9411 \
		jaegertracing/jaeger:2.0.0 \
		--set receivers.otlp.protocols.http.endpoint=0.0.0.0:4318 \
			--set receivers.otlp.protocols.grpc.endpoint=0.0.0.0:4317

create_db:
	docker exec -it postgres createdb --username=$(POSTGRES_USER) --owner=$(POSTGRES_USER) $(POSTGRES_DB)

drop_db:
	docker exec -it postgres dropdb --username=$(POSTGRES_USER) $(POSTGRES_DB)

migrate_up:
	migrate -path db/migrations -database $(POSTGRES_CONNECTION_STRING) -verbose up
	
migrate_down:
	migrate -path db/migrations -database $(POSTGRES_CONNECTION_STRING) -verbose down

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
	@mkdir -p .cover
	go test -v -count=1 -tags=integration ./test/...


cover:
ifndef CI
	go tool cover -html .cover/cover.out
else
	go tool cover -html .cover/cover.out -o .cover/all.html
endif
