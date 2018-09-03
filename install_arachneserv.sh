#! /bin/bash
# create the binary for arachne server at $GOPATH/bin/arachneserv

set -euxo pipefail

pushd $GOPATH/src/github.com/dougfort/arachne/cmd/arachneserv
go install -race
popd
