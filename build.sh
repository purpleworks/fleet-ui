#!/bin/sh

FLEET_VERSION=0.9.2
DOCKER_IMAGE_VERSION=${1:-"latest"}

# echo
echo "FLEET VERSION - "$FLEET_VERSION
echo "BUILD DOCKER IMAGE VERSION - "$DOCKER_IMAGE_VERSION

# build angular
cd angular
npm install
bower install
grunt build
cd ..

# build go app
go install
cp $GOPATH/bin/fleet-ui tmp/
curl -s -L https://github.com/coreos/fleet/releases/download/v${FLEET_VERSION}/fleet-v${FLEET_VERSION}-linux-amd64.tar.gz | \
  tar xz fleet-v${FLEET_VERSION}-linux-amd64/fleetctl -O > tmp/fleetctl
chmod +x tmp/fleetctl
docker build -t purpleworks/fleet-ui:$DOCKER_IMAGE_VERSION .
