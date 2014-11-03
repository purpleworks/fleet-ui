fleet-ui
========

Web based UI for [fleet](https://github.com/coreos/fleet)

![fleet-ui machine list](images/screenshot.png "fleet-ui machine list")

![fleet-ui unit detail](images/screenshot2.png "fleet-ui unit detail")

![fleet-ui new unit](images/screenshot3.png "fleet-ui new unit")

## Getting started

(1) run docker container

- -e ETCD_PEER=`your_etcd_peer_ip`
- -p `port`:3000
- -v `your_ssh_private_key_file_path`:/root/id_rsa

```sh
docker run --rm -p 3000:3000 -e ETCD_PEER=10.0.0.1 -v ~/.ssh/id_rsa:/root/id_rsa purpleworks/fleet-ui
```

(2) enjoy!


## Usage

### docker build

```
$ cd angular
$ npm install
$ bower install
$ npm build
$ cd ..
$ docker build -t purpleworks/fleet-ui .
```

### run in local dev environment
run api server
```
$ go install
$ fleet-ui -etcd-peer=[your_etcd_peer_ip]
```
run web server

```
$ cd angular
$ npm install
$ bower install
$ grunt serve
```

## Prerequire

### go 설치

go가 설치되어 있지 않다면, 아래 페이지를 참고하여 go 환경을 설정한다
https://golang.org/doc/install
https://golang.org/doc/code.html#Organization

정상적으로 설치한 후 아래와 같이 go 환경을 확인한다.

```
$ go version
go version go1.3 darwin/amd64
```

```
$ echo $GOPATH
/Users/your_name/workspace/go
```

### workspace 구성

workspace($GOPATH)에 자신의 github 계정으로 fork한 fleet-ui 소스를 내려받는다.

```
$ mkdir $GOPATH/src/github.com
$ cd $GOPATH/src/github.com
$ git clone git@github.com:your_name/fleet-ui.git
```

최종적으로 아래와 같은 구조가 되도록 한다
```
bin/
pkg/
src/
    github.com/
        your_name/
          fleet-ui/
              .git/
              Dockerfile
              README.md
              app.go
              start.sh
```

## LINK

- [fleet-ui API list](https://github.com/purpleworks/fleet-ui/wiki)
- [coreos/fleet](https://github.com/coreos/fleet)
- [coreos-clustering document](https://coreos.com/using-coreos/clustering/)
- [fleet unit file](https://coreos.com/docs/launching-containers/launching/fleet-unit-files/)

## Contributing

1. Fork it ( https://github.com/purpleworks/fleet-ui/fork )
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create a new Pull Request
