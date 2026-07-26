package repository

import "path/filepath"

// DefinitionFileName is the standard file name of an object definition.
const DefinitionFileName = "definition.yaml"

// DefinitionPath returns the full path to the definition file under dir.
func DefinitionPath(dir string) string {
	return filepath.Join(dir, DefinitionFileName)
}
