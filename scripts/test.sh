#!/usr/bin/env bash

echo "==> Running 'gotest' ..."
go test -covermode=count -coverprofile=coverage.out -v ./akamai/...
