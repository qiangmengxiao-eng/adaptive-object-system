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

	if !fs.root.isDir {
		t.Fatal("root should be a directory")
	}
}

func TestExists(t *testing.T) {
	fs := New()

	tests := []struct {
		name string
		path string
		want bool
	}{
		{
			name: "root exists",
			path: "/",
			want: true,
		},
		{
			name: "missing file",
			path: "/README.md",
			want: false,
		},
		{
			name: "missing directory",
			path: "/docs",
			want: false,
		},
		{
			name: "relative path",
			path: "docs",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := fs.Exists(tt.path)
			if err != nil {
				t.Fatalf("Exists(%q) returned error: %v", tt.path, err)
			}

			if got != tt.want {
				t.Fatalf("Exists(%q) = %v, want %v", tt.path, got, tt.want)
			}
		})
	}
}

func TestStatRoot(t *testing.T) {
	fs := New()

	info, err := fs.Stat("/")
	if err != nil {
		t.Fatalf("Stat returned error: %v", err)
	}

	if info.Path != "/" {
		t.Fatalf("Path = %q, want %q", info.Path, "/")
	}

	if !info.IsDir {
		t.Fatal("root should be a directory")
	}

	if info.Size != 0 {
		t.Fatalf("Size = %d, want 0", info.Size)
	}
}

func TestStatMissing(t *testing.T) {
	fs := New()

	_, err := fs.Stat("/missing")
	if err == nil {
		t.Fatal("expected error for missing path")
	}
}

func TestAddNode(t *testing.T) {
	fs := New()

	if err := fs.addNode("/docs", true, nil); err != nil {
		t.Fatalf("addNode directory failed: %v", err)
	}

	if err := fs.addNode("/docs/api.md", false, []byte("hello")); err != nil {
		t.Fatalf("addNode file failed: %v", err)
	}

	exists, err := fs.Exists("/docs/api.md")
	if err != nil {
		t.Fatalf("Exists returned error: %v", err)
	}

	if !exists {
		t.Fatal("expected /docs/api.md to exist")
	}

	info, err := fs.Stat("/docs/api.md")
	if err != nil {
		t.Fatalf("Stat returned error: %v", err)
	}

	if info.IsDir {
		t.Fatal("api.md should be a file")
	}

	if info.Size != 5 {
		t.Fatalf("Size = %d, want 5", info.Size)
	}
}
