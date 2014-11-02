# docker build
#   docker build -t purpleworks/fleet-ui .
#
# docker run
#   docker run --rm -p [port]:3000 -e ETCD_PEER=[your_etcd_peer_ip] -v [your_ssh_private_key_file_path]:/root/id_rsa purpleworks/fleet-ui
#   docker run --rm -p 3000:3000 -e ETCD_PEER=10.0.0.1 -v ~/.ssh/id_rsa:/root/id_rsa purpleworks/fleet-ui

FROM ubuntu:14.04
MAINTAINER app@purpleworks.co.kr

ENV DEBIAN_FRONTEND noninteractive 

# update ubuntu latest
RUN \
  apt-get -qq update && \
  apt-get -qq -y dist-upgrade

# install essential packages
RUN \
  apt-get -qq -y install build-essential software-properties-common python-software-properties git curl

# install fleetctl
RUN curl -s -L https://github.com/coreos/fleet/releases/download/v0.8.3/fleet-v0.8.3-linux-amd64.tar.gz | tar xz -C /tmp && \
  mv /tmp/fleet-v0.8.3-linux-amd64/fleetctl /usr/local/bin

## install golang 1.3.1
RUN mkdir /goroot && curl -s -L https://storage.googleapis.com/golang/go1.3.1.linux-amd64.tar.gz | tar xz -C /goroot --strip-components=1
RUN mkdir /gopath
ENV GOROOT /goroot
ENV GOPATH /gopath
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin

# install go package
RUN \
  go get -u github.com/mattn/go-sqlite3 && \
  go get -u github.com/codegangsta/negroni && \
  go get -u github.com/gorilla/mux && \
  go get -u github.com/gorilla/websocket && \
  go get -u gopkg.in/unrolled/render.v1

RUN \
  go get -u github.com/coreos/fleet/schema && \
  go get -u github.com/juju/errgo

# add source
WORKDIR /gopath/src/github.com/jaehue/fleet-ui
ADD . /gopath/src/github.com/jaehue/fleet-ui
RUN go install
RUN cp -r angular/dist ./public/

# Add VOLUME
VOLUME  ["/root/id_rsa"]

# export port
EXPOSE 3000

# run!
CMD ./run.sh
