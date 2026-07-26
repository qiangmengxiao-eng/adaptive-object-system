package memory

// node represents a file or directory in memory.
type node struct {
	name string

	isDir bool

	data []byte

	children map[string]*node
}
