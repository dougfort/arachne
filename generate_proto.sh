#!/bin/bash

set -e
set -x

protoc -I arachne/ arachne/arachne.proto --go_out=plugins=grpc:arachne
