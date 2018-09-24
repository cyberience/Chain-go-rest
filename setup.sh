#!/usr/bin/env bash

# Set the GoPath if not already, prevent unlimited growth
if [[ $PATH != *"go"* ]]; then
    export GOPATH=~/go
    export PATH=$PATH:$GOPATH/bin
fi

# Install the Loibraries for Go to build
go get -v -u github.com/gorilla/mux
go get -v -u github.com/AENCO-Global/Chain-go-sdk