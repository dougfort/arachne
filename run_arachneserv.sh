#! /bin/bash
# run the binary for arachneserv

set -euxo pipefail

export ARACHNE_ADDRESS=":10000"

$GOPATH/bin/arachneserv
