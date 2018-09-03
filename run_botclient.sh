#! /bin/bash
# run the binary for the bot client

set -euxo pipefail

export ARACHNE_ADDRESS=":10000"

$GOPATH/bin/botclient -max-turns=100 \
	|& tee /tmp/botclient.log
