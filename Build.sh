#!/bin/bash
VERSION="latest"
if [ $1 != "" ]; then
    VERSION=$1
fi

echo "Building efsctl test docker image ..."
echo "VERSION is $VERSION"

# Build the app using version number specified as the first script parameter
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o efsctl main.go
docker build -t dsalin/efsctl:$VERSION .
