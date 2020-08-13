package parser

import "github.com/ranakan19/parser/pkg/testingutil/filesystem"

// GetFs returns the filesystem object
func (d *DevfileCtx) GetFs() filesystem.Filesystem {
	return d.Fs
}
