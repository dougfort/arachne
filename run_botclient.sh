#! /bin/bash
# run the binary for the bot client

set -e
set -x

$GOPATH/bin/botclient |& tee /tmp/botclient.log
