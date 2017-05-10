#! /bin/bash
# create the binary for the text client at $GOPATH/bin/textclient

set -e
set -x

pushd $GOPATH/src/github.com/dougfort/arachne/textclient
go install -race
popd
