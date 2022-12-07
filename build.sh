#!/bin/bash

mkdir -p ./_test

langs=(
    golang \
    modernjs \
    python \
    bash \
)

for lang in ${langs[@]}
do
    go build -o ./_test/$lang/clerk main.go
done
