package memory

import "testing"

func TestNew(t *testing.T) {
	fs := New()

	if fs == nil {
		t.Fatal("New returned nil")
	}

	if fs.root == nil {
		t.Fatal("root is nil")
	}
}
