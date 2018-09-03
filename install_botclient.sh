#! /bin/bash
# create the binary for the text client at $GOPATH/bin/botclient

set -euxo pipefail

pushd $GOPATH/src/github.com/dougfort/arachne/cmd/botclient
go install -race
popd
