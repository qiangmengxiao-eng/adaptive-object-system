package repository

import (
	"fmt"
	"path"
	"sort"

	"gopkg.in/yaml.v3"
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

// DeleteObject removes an object and its definition file.
func (r *Repository) DeleteObject(name string) error {
	if r == nil || r.fs == nil {
		return fmt.Errorf("repository filesystem is nil")
	}

	objectPath := path.Join("/objects", name)

	definitionPath := DefinitionPath(objectPath)

	// Remove definition file first.
	if err := r.fs.Delete(definitionPath); err != nil {
		return err
	}

	// Remove now-empty object directory.
	return r.fs.Delete(objectPath)
}

// ListObjects returns all object names.
func (r *Repository) ListObjects() ([]string, error) {
	if r == nil || r.fs == nil {
		return nil, fmt.Errorf("repository filesystem is nil")
	}

	entries, err := r.fs.ReadDir("/objects")
	if err != nil {
		return nil, err
	}

	objects := make([]string, 0)

	for _, entry := range entries {
		if entry.IsDirectory {
			objects = append(objects, entry.Name)
		}
	}

	sort.Strings(objects)

	return objects, nil
}

// ExistsObject reports whether an object exists.
func (r *Repository) ExistsObject(name string) (bool, error) {
	if r == nil || r.fs == nil {
		return false, fmt.Errorf("repository filesystem is nil")
	}

	objectPath := path.Join("/objects", name)

	return r.fs.Exists(objectPath)
}

// ReadObjectDefinition reads and parses an object's definition.
func (r *Repository) ReadObjectDefinition(name string) (*ObjectDefinition, error) {
	data, err := r.ReadObject(name)
	if err != nil {
		return nil, err
	}

	var definition ObjectDefinition

	if err := yaml.Unmarshal(data, &definition); err != nil {
		return nil, err
	}

	return &definition, nil
}
