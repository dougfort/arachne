#! /bin/bash
# run the binary for the text client

set -euxo pipefail

export ARACHNE_ADDRESS=":10000"

$GOPATH/bin/textclient
