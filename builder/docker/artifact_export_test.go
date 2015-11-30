package docker

import (
	"github.com/orivej/packer/packer"
	"testing"
)

func TestExportArtifact_impl(t *testing.T) {
	var _ packer.Artifact = new(ExportArtifact)
}
