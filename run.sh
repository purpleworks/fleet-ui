#!/bin/sh

export ETCD_PEER=${ETCD_PEER:-"172.17.42.1"}

if [ -s /root/id_rsa ]
then
  eval `ssh-agent -s` && ssh-add /root/id_rsa
fi

/gopath/bin/fleet-ui
