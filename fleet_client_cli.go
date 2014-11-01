package main

import (
	"github.com/juju/errgo"
	execPkg "os/exec"
)

const (
	FLEETCTL        = "fleetctl"
	ENDPOINT_OPTION = "--endpoint"
	ENDPOINT_VALUE  = "http://172.17.42.1:4001"
)

type ClientCLI struct {
	etcdPeer string
}

func NewClientCLI() FleetClient {
	return NewClientCLIWithPeer(ENDPOINT_VALUE)
}

func NewClientCLIWithPeer(etcdPeer string) FleetClient {
	return &ClientCLI{
		etcdPeer: etcdPeer,
	}
}

func (this *ClientCLI) Submit(name, filePath string) error {
	cmd := execPkg.Command(FLEETCTL, ENDPOINT_OPTION, this.etcdPeer, "submit", filePath)
	_, err := exec(cmd)

	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}

func (this *ClientCLI) Start(name string) error {
	cmd := execPkg.Command(FLEETCTL, ENDPOINT_OPTION, this.etcdPeer, "start", "--no-block=true", name)
	_, err := exec(cmd)

	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}

func (this *ClientCLI) Stop(name string) error {
	cmd := execPkg.Command(FLEETCTL, ENDPOINT_OPTION, this.etcdPeer, "stop", "--no-block=true", name)
	_, err := exec(cmd)

	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}

func (this *ClientCLI) Load(name string) error {
	cmd := execPkg.Command(FLEETCTL, ENDPOINT_OPTION, this.etcdPeer, "load", "--no-block=true", name)
	_, err := exec(cmd)

	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}

func (this *ClientCLI) Destroy(name string) error {
	cmd := execPkg.Command(FLEETCTL, ENDPOINT_OPTION, this.etcdPeer, "destroy", name)
	_, err := exec(cmd)

	if err != nil {
		return errgo.Mask(err)
	}

	return nil
}
