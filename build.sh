#!/usr/bin/env bash
set -eu

go build -o luca.bin ./main.go
zip luca.zip luca.bin
rm luca.bin
