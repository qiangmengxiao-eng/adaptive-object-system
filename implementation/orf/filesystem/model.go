package filesystem

// DirectoryEntry describes an entry in a repository directory.
type DirectoryEntry struct {
	Name string
	Path string

	IsDirectory bool
}

// FileInfo describes metadata for a repository path.
type FileInfo struct {
	Path string

	Size int64

	IsDir bool
}
