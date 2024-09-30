#!/bin/bash

set -e
cp /app/docker/services/env/$SERVICE_NAME.env /app/.env
go build -o $SERVICE_NAME -buildvcs="false" /app/cmd/$SERVICE_NAME > /dev/stdout
/app/$SERVICE_NAME