#!/bin/bash

#echo "Make sure current directory is code's root folder(main.go in the root folder)"

#export GOPRIVATE=code.nurture.farm
go mod tidy
go mod vendor
go fmt
echo "Downloaded dependencies for go modules"

CONFIG_DIR="${CONFIG_DIR:-$PWD/config}"
echo "CONFIG_DIR = $CONFIG_DIR, config is read from this directory if you want to modify this path pass variable and run it again, CONFIG_DIR=<...> ./run_local.sh"
echo "We read db_config.json(db connection params), other files from $CONFIG_DIR, make sure you have right values in it"

echo "Building the code"
cd zerotouch/golang/ || exit
go build -o main . || exit

echo "Build successful! Starting on port 5000, make sure this port is not already used by another process"
echo "**************************** Application log from here onwards **************************"
DB_USERNAME=root DB_PASSWORD=12345678 CONFIG_DIR=$CONFIG_DIR ./main
