#! /bin/bash
# run the binary for the text client

set -e
set -x

export ARACHNE_ADDRESS=":10000"

$GOPATH/bin/textclient
