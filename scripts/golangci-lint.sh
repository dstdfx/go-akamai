#!/usr/bin/env bash

echo "==> Running 'golangci-lint' ..."
golangci-lint run ./...
