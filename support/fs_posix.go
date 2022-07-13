package support

import (
	"github.com/bitwormhole/afs"
)

// GetPosixFS ...
func GetPosixFS() afs.FS {
	platform := &myPosixFS{}
	return NewFS(platform)
}

// core
type myPosixFS struct {
	core CommonFileSystemCore
}

func (inst *myPosixFS) _Impl() PlatformFileSystem {
	return inst
}

func (inst *myPosixFS) GetCommonFileSystem() CommonFileSystem {
	return &inst.core
}

func (inst *myPosixFS) NormalizePath(path string) (string, error) {
	sep := inst.Separator()
	elements := inst.core.PathToElements(path)
	elements, err := inst.core.NormalizePathElements(elements)
	if err != nil {
		return "", err
	}
	path = inst.core.ElementsToPath(elements, sep, sep)
	return path, nil
}

func (inst *myPosixFS) PathSeparator() string {
	return ":"
}

func (inst *myPosixFS) Separator() string {
	return "/"
}

func (inst *myPosixFS) ListRoots() []string {
	return []string{"/"}
}
