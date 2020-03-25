#!/usr/bin/env bash
set -x

# Circle checks out the code inside the GOPATH.
# This throws off godog.
# Hence copying it outside the GOPATH
# so that Go modules work correctly.
cp -r /go/src/github.com/kevgo/tikibase/ ~
cd ~/tikibase
GO111MODULE=on godog --format=progress
