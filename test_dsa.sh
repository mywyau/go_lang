#!/usr/bin/env bash


set -e

go fmt ./dsa

# avoids cached results 
go test -v -count=1 ./dsa 