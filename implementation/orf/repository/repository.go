package repository

import "github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/filesystem"

// Repository represents an object repository backed by a filesystem.
type Repository struct {
	fs filesystem.MutableRepositoryFS
}

// New creates a repository backed by the given filesystem.
func New(fs filesystem.MutableRepositoryFS) *Repository {
	return &Repository{
		fs: fs,
	}
}

// FS returns the underlying filesystem.
func (r *Repository) FS() filesystem.MutableRepositoryFS {
	return r.fs
}

// EnsureObjectsDirectory creates the objects directory if missing.
func (r *Repository) EnsureObjectsDirectory() error {
	exists, err := r.fs.Exists("/objects")
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	return r.fs.Mkdir("/objects")
}
