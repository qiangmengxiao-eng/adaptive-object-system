package memory

import "github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/filesystem"

// MemoryFS is an in-memory implementation of filesystem.RepositoryFS.
type MemoryFS struct {
	root *node
}

// New creates a new empty in-memory filesystem.
func New() *MemoryFS {
	return &MemoryFS{
		root: &node{
			name:     "/",
			isDir:    true,
			children: make(map[string]*node),
		},
	}
}

// Ensure MemoryFS implements filesystem.RepositoryFS.
var _ filesystem.RepositoryFS = (*MemoryFS)(nil)

// Exists reports whether a path exists.
func (m *MemoryFS) Exists(path string) (bool, error) {
	return false, nil
}

// ReadDir returns the directory entries under path.
func (m *MemoryFS) ReadDir(path string) ([]filesystem.DirectoryEntry, error) {
	return nil, nil
}

// ReadFile reads a file.
func (m *MemoryFS) ReadFile(path string) ([]byte, error) {
	return nil, nil
}

// Stat returns file information.
func (m *MemoryFS) Stat(path string) (filesystem.FileInfo, error) {
	return filesystem.FileInfo{}, nil
}
