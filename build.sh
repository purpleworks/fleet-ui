#!/bin/sh

go install
cp $GOPATH/bin/fleet-ui tmp/
curl -s -L https://github.com/coreos/fleet/releases/download/v0.9.0/fleet-v0.9.0-linux-amd64.tar.gz | tar xz -C tmp/
docker build -t purpleworks/fleet-ui:$1 .
