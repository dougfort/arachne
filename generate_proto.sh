#!/bin/bash

set -euxo pipefail

protoc -I arachne/ arachne/arachne.proto --go_out=plugins=grpc:arachne
