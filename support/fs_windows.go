package support

import (
	"os"
	"strings"

	"github.com/bitwormhole/afs"
)

// GetWindowsFS ...
func GetWindowsFS() afs.FS {
	platform := &myWindowsFS{}
	return NewFS(platform)
}

////////////////////////////////////////////////////////////////////////////////

type myWindowsFS struct {
	core CommonFileSystemCore
}

func (inst *myWindowsFS) _Impl() PlatformFileSystem {
	return inst
}

func (inst *myWindowsFS) NormalizePath(path string) (string, error) {
	sep := inst.Separator()
	elements := inst.core.PathToElements(path)
	elements, err := inst.core.NormalizePathElements(elements)
	if err != nil {
		return "", err
	}
	path = inst.core.ElementsToPath(elements, "", sep)
	if strings.HasSuffix(path, ":") {
		if !strings.Contains(path, sep) {
			return path + sep, nil
		}
	}
	return path, nil
}

func (inst *myWindowsFS) PathSeparator() string {
	return ";"
}

func (inst *myWindowsFS) Separator() string {
	return "\\"
}

func (inst *myWindowsFS) isRootExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func (inst *myWindowsFS) ListRoots() []string {
	const driveA rune = 'A'
	const driveZ rune = 'Z'
	list := make([]string, 0)
	for drive := driveA; drive <= driveZ; drive++ {
		path := string(drive) + ":\\"
		if inst.isRootExists(path) {
			list = append(list, path)
		}
	}
	return list
}

func (inst *myWindowsFS) GetCommonFileSystem() CommonFileSystem {
	return &inst.core
}
