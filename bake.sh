#!/bin/bash

set -e

docker-compose build
docker-compose run app
docker build -f Dockerfile.bake -t nathanleclaire/hubfwd .
docker push nathanleclaire/hubfwd
