#! /bin/bash
# run the binary for the bot client

set -e
set -x

$GOPATH/bin/botclient -max-turns=10 \
	|& tee /tmp/botclient.log
