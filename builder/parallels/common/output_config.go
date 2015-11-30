package common

import (
	"fmt"
	"os"

	"github.com/orivej/packer/common"
	"github.com/orivej/packer/template/interpolate"
)

type OutputConfig struct {
	OutputDir string `mapstructure:"output_directory"`
}

func (c *OutputConfig) Prepare(ctx *interpolate.Context, pc *common.PackerConfig) []error {
	if c.OutputDir == "" {
		c.OutputDir = fmt.Sprintf("output-%s", pc.PackerBuildName)
	}

	var errs []error
	if !pc.PackerForce {
		if _, err := os.Stat(c.OutputDir); err == nil {
			errs = append(errs, fmt.Errorf(
				"Output directory '%s' already exists. It must not exist.", c.OutputDir))
		}
	}

	return errs
}
