package common

import (
	"github.com/orivej/packer/template/interpolate"
)

// FloppyConfig is configuration related to created floppy disks and attaching
// them to a VirtualBox machine.
type FloppyConfig struct {
	FloppyFiles []string `mapstructure:"floppy_files"`
}

func (c *FloppyConfig) Prepare(ctx *interpolate.Context) []error {
	if c.FloppyFiles == nil {
		c.FloppyFiles = make([]string, 0)
	}

	return nil
}
