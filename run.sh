#!/bin/sh

eval `ssh-agent -s` && ssh-add /root/id_rsa && /gopath/bin/fleet-ui
