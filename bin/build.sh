#!/bin/bash

case $1 in
    "test")
        mkdir -p ./_test
        go build -o ./_test/clerk_test -ldflags "-X main.Version=$2" main.go
    ;;
esac
