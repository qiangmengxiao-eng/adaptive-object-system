package memory

// node represents a file or directory in the in-memory filesystem.
type node struct {
	name     string
	isDir    bool
	data     []byte
	children map[string]*node
}

// newDirectory creates a directory node.
func newDirectory(name string) *node {
	return &node{
		name:     name,
		isDir:    true,
		children: make(map[string]*node),
	}
}

// newFile creates a file node.
func newFile(name string, data []byte) *node {
	copied := append([]byte(nil), data...)

	return &node{
		name:  name,
		isDir: false,
		data:  copied,
	}
}
