# docker build
#   docker build -t purpleworks/fleet-ui .
#
# docker run
#   docker run --rm -p [port]:3000 -e ETCD_PEER=[your_etcd_peer_ip] -v [your_ssh_private_key_file_path]:/root/id_rsa purpleworks/fleet-ui
#   docker run --rm -p 3000:3000 -e ETCD_PEER=10.0.0.1 -v ~/.ssh/id_rsa:/root/id_rsa purpleworks/fleet-ui

FROM progrium/busybox
MAINTAINER app@purpleworks.co.kr

# install packages
RUN opkg-install openssh-client-utils

# add files
ADD run.sh /root/fleet-ui/run.sh
ADD tmp/fleet-ui /root/fleet-ui/fleet-ui
ADD tmp/fleetctl /usr/local/bin/fleetctl
ADD angular/dist /root/fleet-ui/public

# set workdir
WORKDIR /root/fleet-ui

# export port
EXPOSE 3000

# run!
CMD ./run.sh
