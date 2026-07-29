package repository

// Storage defines repository file operations.
type Storage interface {
	ReadFile(
		path string,
	) ([]byte, error)

	WriteFile(
		path string,
		data []byte,
	) error

	Delete(
		path string,
	) error

	Exists(
		path string,
	) (bool, error)
}
