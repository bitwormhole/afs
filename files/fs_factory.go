package files

import (
	"os"

	"github.com/bitwormhole/afs"
	"github.com/bitwormhole/afs/support"
)

// FS ...
func FS() afs.FS {
	o := os.Getenv("os")
	if o == "windows" {
		return support.GetWindowsFS()
	}
	return support.GetPosixFS()
}
