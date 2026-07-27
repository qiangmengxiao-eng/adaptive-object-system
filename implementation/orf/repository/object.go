package repository

import (
	"fmt"
	"path"
)

// CreateObject creates a new object with its definition file.
func (r *Repository) CreateObject(name string, definition []byte) error {
	if r == nil || r.fs == nil {
		return fmt.Errorf("repository filesystem is nil")
	}

	objectPath := path.Join("/objects", name)

	if err := r.fs.Mkdir("/objects"); err != nil {
		// allow existing objects directory
		exists, existsErr := r.fs.Exists("/objects")
		if existsErr != nil || !exists {
			return err
		}
	}

	if err := r.fs.Mkdir(objectPath); err != nil {
		return err
	}

	definitionPath := DefinitionPath(objectPath)

	return r.fs.WriteFile(definitionPath, definition)
}

// ReadObject reads an object's definition file.
func (r *Repository) ReadObject(name string) ([]byte, error) {
	if r == nil || r.fs == nil {
		return nil, fmt.Errorf("repository filesystem is nil")
	}

	objectPath := path.Join("/objects", name)

	definitionPath := DefinitionPath(objectPath)

	return r.fs.ReadFile(definitionPath)
}
