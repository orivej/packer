package common

import (
	"github.com/orivej/packer/template/interpolate"
)

type PrlctlConfig struct {
	Prlctl [][]string `mapstructure:"prlctl"`
}

func (c *PrlctlConfig) Prepare(ctx *interpolate.Context) []error {
	if c.Prlctl == nil {
		c.Prlctl = make([][]string, 0)
	}

	return nil
}
