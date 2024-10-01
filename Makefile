.PHONY: start

VERSION := v0.1.0

ifneq ("$(wildcard .env)","")
	include .env
endif

GOLANG_BINARY := /usr/local/go/bin/go

AUTH_SERVICE_NAME := auth
BACKEND_SERVICE_NAME := backend
GATEWAY_ADMIN_SERVICE_NAME := gateway_admin
GATEWAY_USER_SERVICE_NAME := gateway_user
GATEWAY_FIAT_COMPANY_SERVICE_NAME := gateway_fiat_company

LDFLAGS :=
DOCKER=$(which docker)

build: \
	build-auth \
	build-backend \
	build-gateway-admin \
	build-gateway-user \
	build-gateway-fiat-company
	
build-backend:
	go build -o binaries/$(BACKEND_SERVICE_NAME)/$(BACKEND_SERVICE_NAME) -ldflags="$(LDFLAGS)" ./cmd/$(BACKEND_SERVICE_NAME)/

seedup-backend:
	cat ./database/backend/seeds/seeds.up.sql | psql postgres://$(PSQL_BACKEND_USER):$(PSQL_BACKEND_PASS)@$(PSQL_BACKEND_HOST):$(PSQL_BACKEND_PORT)/$(PSQL_BACKEND_DB)

seeddown-backend:
	cat ./database/backend/seeds/seeds.down.sql | psql postgres://$(PSQL_BACKEND_USER):$(PSQL_BACKEND_PASS)@$(PSQL_BACKEND_HOST):$(PSQL_BACKEND_PORT)/$(PSQL_BACKEND_DB)

drop-backend:
	cat ./database/backend/seeds/drop.sql | psql postgres://$(PSQL_BACKEND_USER):$(PSQL_BACKEND_PASS)@$(PSQL_BACKEND_HOST):$(PSQL_BACKEND_PORT)/$(PSQL_BACKEND_DB)

drop-all: \
	drop-auth \
	drop-backend \
	drop-logger

seeddown-all: \
	seeddown-backend \
	seeddown-auth
seedup-all: \
	seedup-auth \
	seedup-backend

mup-auth:
	migrate -database postgres://$(PSQL_AUTH_USER):$(PSQL_AUTH_PASS)@$(PSQL_AUTH_HOST):$(PSQL_AUTH_PORT)/$(PSQL_AUTH_DB)?sslmode=disable -path ./database/auth/migrations up

mdown-auth:
	migrate -database postgres://$(PSQL_AUTH_USER):$(PSQL_AUTH_PASS)@$(PSQL_AUTH_HOST):$(PSQL_AUTH_PORT)/$(PSQL_AUTH_DB)?sslmode=disable -path ./database/auth/migrations down -all

mup-backend:
	migrate -database postgres://$(PSQL_BACKEND_USER):$(PSQL_BACKEND_PASS)@$(PSQL_BACKEND_HOST):$(PSQL_BACKEND_PORT)/$(PSQL_BACKEND_DB)?sslmode=disable -path ./database/backend/migrations up

mdown-backend:
	migrate -database postgres://$(PSQL_BACKEND_USER):$(PSQL_BACKEND_PASS)@$(PSQL_BACKEND_HOST):$(PSQL_BACKEND_PORT)/$(PSQL_BACKEND_DB)?sslmode=disable -path ./database/backend/migrations down -all

setenv:
	set -o allexport; . ./.env; set +o allexport

test:
	go clean -testcache
	go test ./...

gen:
	go generate ./...

vuln:
	govulncheck ./...

install-deps:
	go install github.com/golang/mock/mockgen@latest
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	go install golang.org/x/vuln/cmd/govulncheck@latest
