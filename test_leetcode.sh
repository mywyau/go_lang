#!/usr/bin/env bash


set -e

go fmt ./leetcode/...

# avoids cached results 
go test -v -count=1 ./leetcode/... 