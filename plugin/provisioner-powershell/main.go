package main

import (
	"github.com/orivej/packer/packer/plugin"
	"github.com/orivej/packer/provisioner/powershell"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(new(powershell.Provisioner))
	server.Serve()
}
