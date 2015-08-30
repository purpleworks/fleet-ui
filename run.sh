#!/bin/sh

export ETCD_PEER=${ETCD_PEER:-http://$(/sbin/ip route|awk '/default/ { print $3 }'):4001}

if [ -f /root/id_rsa ]; then
  eval `ssh-agent -s` && ssh-add /root/id_rsa
fi

/root/fleet-ui/fleet-ui

