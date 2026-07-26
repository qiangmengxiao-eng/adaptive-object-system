package filesystem

// RepositoryFS provides read-only access to an object repository.
type RepositoryFS interface {
	// Exists reports whether the specified repository path exists.
	Exists(path string) (bool, error)

	// ReadDir reads the entries in a repository directory.
	ReadDir(path string) ([]DirectoryEntry, error)

	// ReadFile reads the contents of a repository file.
	ReadFile(path string) ([]byte, error)

	// Stat returns metadata for a repository path.
	Stat(path string) (FileInfo, error)
}
