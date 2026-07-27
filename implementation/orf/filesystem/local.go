package filesystem

import (
	"os"
	"path/filepath"
)

// LocalFS is a filesystem backed by local disk.
type LocalFS struct {
	root string
}

// NewLocal creates a local filesystem.
func NewLocal(root string) *LocalFS {
	return &LocalFS{
		root: root,
	}
}

func (l *LocalFS) resolve(path string) string {
	return filepath.Join(l.root, filepath.Clean(path))
}

// Exists checks whether path exists.
func (l *LocalFS) Exists(path string) (bool, error) {
	_, err := os.Stat(l.resolve(path))

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

// Mkdir creates directory.
func (l *LocalFS) Mkdir(path string) error {
	return os.MkdirAll(l.resolve(path), 0755)
}

// WriteFile writes data.
func (l *LocalFS) WriteFile(path string, data []byte) error {
	full := l.resolve(path)

	if err := os.MkdirAll(filepath.Dir(full), 0755); err != nil {
		return err
	}

	return os.WriteFile(full, data, 0644)
}

// ReadFile reads data.
func (l *LocalFS) ReadFile(path string) ([]byte, error) {
	return os.ReadFile(l.resolve(path))
}

// Delete removes path.
func (l *LocalFS) Delete(path string) error {
	return os.RemoveAll(l.resolve(path))
}

// ReadDir reads directory.
func (l *LocalFS) ReadDir(path string) ([]DirectoryEntry, error) {

	entries, err := os.ReadDir(l.resolve(path))

	if err != nil {
		return nil, err
	}

	result := make([]DirectoryEntry, 0, len(entries))

	for _, entry := range entries {

		result = append(result, DirectoryEntry{
			Name:        entry.Name(),
			Path:        filepath.Join(path, entry.Name()),
			IsDirectory: entry.IsDir(),
		})
	}

	return result, nil
}

// Stat returns file info.
func (l *LocalFS) Stat(path string) (FileInfo, error) {

	info, err := os.Stat(l.resolve(path))

	if err != nil {
		return FileInfo{}, err
	}

	return FileInfo{
		Path:  path,
		Size:  info.Size(),
		IsDir: info.IsDir(),
	}, nil
}
