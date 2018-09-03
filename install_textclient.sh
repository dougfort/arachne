#! /bin/bash
# create the binary for the text client at $GOPATH/bin/textclient

set -euxo pipefail

pushd $GOPATH/src/github.com/dougfort/arachne/cmd/textclient
go install -race
popd
