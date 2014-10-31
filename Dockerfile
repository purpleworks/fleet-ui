# docker build
#   docker build -t purpleworks/fleet-ui .
#
# docker run
#   docker run --rm -it -p 3000:3000 -v ~/.ssh/id_rsa:/ssh/id_rsa purpleworks/fleet-ui

FROM dockerfile/ubuntu
MAINTAINER jaehue@jang.io

# update ubuntu latest
RUN \
  apt-get -qq update && \
  apt-get -qq -y dist-upgrade

# install fleetctl
RUN cd /root &&  \
  wget -q https://github.com/coreos/fleet/releases/download/v0.8.3/fleet-v0.8.3-linux-amd64.tar.gz && \
  tar -xzf fleet-v0.8.3-linux-amd64.tar.gz && \
  mv /root/fleet-v0.8.3-linux-amd64/fleetctl /usr/local/bin

## install golang 1.3.1
RUN mkdir /goroot && curl https://storage.googleapis.com/golang/go1.3.1.linux-amd64.tar.gz | tar xvzf - -C /goroot --strip-components=1
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

# set caching point
ENV REFRESHED_AT 2014-10-31 4

# go get fleet-client-go
RUN go get -u github.com/jaehue/fleet-client-go

# add source
WORKDIR /gopath/src/github.com/jaehue/fleet-ui
ADD . /gopath/src/github.com/jaehue/fleet-ui
RUN go install

# Add VOLUME
VOLUME  ["/ssh/id_rsa"]

# export port
EXPOSE 3000

CMD eval `ssh-agent -s` && ssh-add /ssh/id_rsa && /gopath/bin/fleet-ui
