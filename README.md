fleet-ui
========

Web based UI for [fleet](https://github.com/coreos/fleet)

![fleet-ui machine list](images/screenshot.png "fleet-ui machine list")

![fleet-ui unit detail](images/screenshot2.png "fleet-ui unit detail")

![fleet-ui new unit](images/screenshot3.png "fleet-ui new unit")

## Getting started

(1) run docker container

- -e ETCD_PEER=`your_etcd_peer_ip:peer_port`
- -p `port`:3000
- -v `your_ssh_private_key_file_path`:/root/id_rsa

```sh
docker run --rm -p 3000:3000 -e ETCD_PEER=10.0.0.1:4001 -v ~/.ssh/id_rsa:/root/id_rsa purpleworks/fleet-ui
```

(2) enjoy!

## Prerequire

- backend
  - [go](http://golang.org/doc/install)
  - [fleetctl](https://github.com/coreos/fleet/releases)
- frontend
  - [nodejs](https://nodejs.org/)
  - [yeoman](http://yeoman.io/)
  - [grunt](http://gruntjs.com/)
  - [bower](http://bower.io/)
  - [ruby](https://www.ruby-lang.org/en/downloads/)
  - [compass](http://compass-style.org/install/)

## Organize your workspace

clone your forked github repository to workspace($GOPATH)

```
$ mkdir $GOPATH/src/github.com
$ cd $GOPATH/src/github.com
$ git clone git@github.com:your_name/fleet-ui.git
```

here's an example:
```
bin/
pkg/
src/
    github.com/
        your_name/
          fleet-ui/
              .git/
              .dockerignore
              .gitignore
              CHANGELOG.md
              Dockerfile
              README.md
              angular
              app.go
              (...)
```

## Development

### backend (api server)

```
$ go install
$ fleet-ui -etcd-peer=[your_etcd_peer_ip]
```

### frontend (web based ui)

```
$ cd angular
$ npm install
$ bower install
$ grunt server
```

## Build All

```
$ ./build.sh
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

## License
MIT (see [LICENSE](LICENSE) file)
