package memory

import (
	"path"
	"strings"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/filesystem"
)

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

// find locates a node by its absolute path.
func (m *MemoryFS) find(p string) *node {
	if m == nil || m.root == nil {
		return nil
	}

	p = path.Clean(p)

	if p == "/" {
		return m.root
	}

	if !strings.HasPrefix(p, "/") {
		return nil
	}

	parts := strings.Split(strings.TrimPrefix(p, "/"), "/")

	current := m.root

	for _, part := range parts {
		if part == "" {
			continue
		}

		if !current.isDir {
			return nil
		}

		child, ok := current.children[part]
		if !ok {
			return nil
		}

		current = child
	}

	return current
}

// Exists reports whether a path exists.
func (m *MemoryFS) Exists(path string) (bool, error) {
	return m.find(path) != nil, nil
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
