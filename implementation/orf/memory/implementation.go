package memory

import (
	"fmt"
	"path"
	"sort"
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
		root: newDirectory("/"),
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

// addNode creates a file or directory under an existing parent directory.
func (m *MemoryFS) addNode(p string, isDir bool, data []byte) error {
	if m == nil || m.root == nil {
		return fmt.Errorf("filesystem is nil")
	}

	p = path.Clean(p)

	if p == "/" {
		return fmt.Errorf("cannot replace root")
	}

	parentPath := path.Dir(p)
	name := path.Base(p)

	parent := m.find(parentPath)
	if parent == nil {
		return fmt.Errorf("parent does not exist: %s", parentPath)
	}

	if !parent.isDir {
		return fmt.Errorf("parent is not a directory: %s", parentPath)
	}

	if _, exists := parent.children[name]; exists {
		return fmt.Errorf("path already exists: %s", p)
	}

	if isDir {
		parent.children[name] = newDirectory(name)
	} else {
		parent.children[name] = newFile(name, data)
	}

	return nil
}

// Mkdir creates a directory.
func (m *MemoryFS) Mkdir(p string) error {
	if m == nil || m.root == nil {
		return fmt.Errorf("filesystem is nil")
	}

	p = path.Clean(p)

	if p == "/" {
		return fmt.Errorf("cannot create root directory")
	}

	parentPath := path.Dir(p)
	name := path.Base(p)

	parent := m.find(parentPath)
	if parent == nil {
		return fmt.Errorf("parent does not exist: %s", parentPath)
	}

	if !parent.isDir {
		return fmt.Errorf("parent is not a directory: %s", parentPath)
	}

	if _, exists := parent.children[name]; exists {
		return fmt.Errorf("path already exists: %s", p)
	}

	parent.children[name] = newDirectory(name)

	return nil
}

// WriteFile writes data to a file.
// If the file exists, its contents are replaced.
func (m *MemoryFS) WriteFile(p string, data []byte) error {
	if m == nil || m.root == nil {
		return fmt.Errorf("filesystem is nil")
	}

	p = path.Clean(p)

	if p == "/" {
		return fmt.Errorf("cannot write root directory")
	}

	parentPath := path.Dir(p)
	name := path.Base(p)

	parent := m.find(parentPath)
	if parent == nil {
		return fmt.Errorf("parent does not exist: %s", parentPath)
	}

	if !parent.isDir {
		return fmt.Errorf("parent is not a directory: %s", parentPath)
	}

	if existing, ok := parent.children[name]; ok {
		if existing.isDir {
			return fmt.Errorf("path is a directory: %s", p)
		}

		existing.data = append([]byte(nil), data...)
		return nil
	}

	parent.children[name] = newFile(name, data)

	return nil
}

// Delete removes a file or an empty directory.
func (m *MemoryFS) Delete(p string) error {
	if m == nil || m.root == nil {
		return fmt.Errorf("filesystem is nil")
	}

	p = path.Clean(p)

	if p == "/" {
		return fmt.Errorf("cannot delete root directory")
	}

	parentPath := path.Dir(p)
	name := path.Base(p)

	parent := m.find(parentPath)
	if parent == nil {
		return fmt.Errorf("parent does not exist: %s", parentPath)
	}

	if !parent.isDir {
		return fmt.Errorf("parent is not a directory: %s", parentPath)
	}

	target, exists := parent.children[name]
	if !exists {
		return fmt.Errorf("path does not exist: %s", p)
	}

	if target.isDir && len(target.children) > 0 {
		return fmt.Errorf("directory is not empty: %s", p)
	}

	delete(parent.children, name)

	return nil
}

// Exists reports whether a path exists.
func (m *MemoryFS) Exists(path string) (bool, error) {
	return m.find(path) != nil, nil
}

// ReadDir returns the directory entries under path.
func (m *MemoryFS) ReadDir(p string) ([]filesystem.DirectoryEntry, error) {
	n := m.find(p)
	if n == nil {
		return nil, fmt.Errorf("path does not exist: %s", p)
	}

	if !n.isDir {
		return nil, fmt.Errorf("path is not a directory: %s", p)
	}

	entries := make([]filesystem.DirectoryEntry, 0, len(n.children))

	for _, child := range n.children {
		entries = append(entries, filesystem.DirectoryEntry{
			Name:        child.name,
			Path:        path.Join(p, child.name),
			IsDirectory: child.isDir,
		})
	}

	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name < entries[j].Name
	})

	return entries, nil
}

// ReadFile reads a file.
func (m *MemoryFS) ReadFile(p string) ([]byte, error) {
	n := m.find(p)
	if n == nil {
		return nil, fmt.Errorf("path does not exist: %s", p)
	}

	if n.isDir {
		return nil, fmt.Errorf("path is a directory: %s", p)
	}

	// Return a copy so callers cannot mutate the internal state.
	data := append([]byte(nil), n.data...)

	return data, nil
}

// Stat returns file information.
func (m *MemoryFS) Stat(path string) (filesystem.FileInfo, error) {
	n := m.find(path)
	if n == nil {
		return filesystem.FileInfo{}, fmt.Errorf("path does not exist: %s", path)
	}

	size := int64(len(n.data))
	if n.isDir {
		size = 0
	}

	return filesystem.FileInfo{
		Path:  path,
		Size:  size,
		IsDir: n.isDir,
	}, nil
}
