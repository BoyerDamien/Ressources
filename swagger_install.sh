#! /bin/bash


go get -u github.com/go-swagger/go-swagger
dir=$(mktemp -d) 
git clone https://github.com/go-swagger/go-swagger "$dir" 
cd "$dir"
go install ./cmd/swagger
