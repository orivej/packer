package main

import (
	"github.com/orivej/packer/packer/plugin"
	"github.com/orivej/packer/post-processor/docker-tag"
)

func main() {
	server, err := plugin.Server()
	if err != nil {
		panic(err)
	}
	server.RegisterPostProcessor(new(dockertag.PostProcessor))
	server.Serve()
}
