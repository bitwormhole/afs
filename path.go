package afs

// Path 是表示绝对路径的接口
type Path interface {
	GetFS() FS

	GetParent() Path

	GetChild(name string) Path

	GetName() string

	String() string

	Length() int64

	Mkdir() error

	Mkdirs() error

	Delete() error

	List() []string

	Exists() bool

	IsFile() bool

	IsDirectory() bool
}
