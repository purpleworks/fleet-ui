#!/bin/sh

go install
cp $GOPATH/bin/fleet-ui tmp/
curl -s -L https://github.com/coreos/fleet/releases/download/v0.8.3/fleet-v0.8.3-linux-amd64.tar.gz | tar xz -C tmp/
docker build -t purpleworks/fleet-ui:$1 .
