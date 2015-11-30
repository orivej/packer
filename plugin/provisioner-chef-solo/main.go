package main

import (
	"github.com/orivej/packer/packer/plugin"
	"github.com/orivej/packer/provisioner/chef-solo"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterProvisioner(new(chefsolo.Provisioner))
	server.Serve()
}
