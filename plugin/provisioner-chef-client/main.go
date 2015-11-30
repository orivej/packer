package main

import (
	"github.com/orivej/packer/packer/plugin"
	"github.com/orivej/packer/provisioner/chef-client"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(new(chefclient.Provisioner))
	server.Serve()
}
