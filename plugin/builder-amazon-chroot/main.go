package main

import (
	"github.com/orivej/packer/builder/amazon/chroot"
	"github.com/orivej/packer/packer/plugin"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterBuilder(new(chroot.Builder))
	server.Serve()
}
