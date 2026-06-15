#!/usr/bin/env bash
VERSION=1.1.2
echo Build project version: ${VERSION} 

docker build -t monitor-service:${VERSION} .
