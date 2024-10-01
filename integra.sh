#!/bin/bash
#cp docker.env .env
go build -o ./cmd/tests/integra/integra ./cmd/tests/integra
./cmd/tests/integra/integra "$@"