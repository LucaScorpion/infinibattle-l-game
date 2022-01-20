#!/usr/bin/env bash
set -eu

echo "= Running all tests..."
go test ./...

echo
echo "= Running benchmarks..."
go test ./internal/bot -bench=.
