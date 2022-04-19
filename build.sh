#!/bin/bash

#go get code.nurture.farm/Core/Contracts
go mod tidy
go mod vendor

# Build docker image
docker build -t blood-bank-system-service:$1 .
