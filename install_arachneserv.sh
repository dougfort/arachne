#! /bin/bash
# create the binary for arachne server at $GOPATH/bin/arachneserv

set -e
set -x

pushd $GOPATH/src/github.com/dougfort/arachne/arachneserv
go install -race
popd
