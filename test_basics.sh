#!/usr/bin/env bash

set -e

go fmt ./basics/...

# avoids cached results 
go test -v -count=1 ./basics/...