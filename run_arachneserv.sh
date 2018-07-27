#! /bin/bash
# run the binary for arachneserv

set -e
set -x

export ARACHNE_ADDRESS=":10000"

$GOPATH/bin/arachneserv &
