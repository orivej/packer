package main

import (
	"github.com/orivej/packer/packer/plugin"
	"github.com/orivej/packer/provisioner/puppet-masterless"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(new(puppetmasterless.Provisioner))
	server.Serve()
}
