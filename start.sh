#!/bin/bash

eval `ssh-agent -s`
ssh-add /ssh/id_rsa
/gopath/bin/fleet-ui