package main

import (
	"github.com/orivej/packer/packer/plugin"
	"github.com/orivej/packer/provisioner/windows-restart"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(new(restart.Provisioner))
	server.Serve()
}
