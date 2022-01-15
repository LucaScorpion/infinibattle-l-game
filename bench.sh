#!/usr/bin/env bash
set -eu

go test ./internal/bot -bench=.
