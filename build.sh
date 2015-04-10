#!/bin/bash

echo "Building statically compiled binary..."

# Shoutout to @kelseyhightower for making this way more accessible.
CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' .
