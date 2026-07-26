package filesystem

type RepositoryFS interface {
	Exists(path string) (bool, error)
	ReadDir(path string) ([]DirectoryEntry, error)
	ReadFile(path string) ([]byte, error)
	Stat(path string) (FileInfo, error)
}

type MutableRepositoryFS interface {
	RepositoryFS

	Mkdir(path string) error
	WriteFile(path string, data []byte) error
	Delete(path string) error
}
