package memory

import "testing"

func TestNew(t *testing.T) {
	fs := New()

	if fs == nil {
		t.Fatal("New returned nil")
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
	}

	for _, tt := range tests {
		got, err := fs.Exists(tt.path)
		if err != nil {
			t.Fatalf("Exists(%q) returned error: %v", tt.path, err)
		}

		if got != tt.want {
			t.Fatalf("Exists(%q) = %v, want %v", tt.path, got, tt.want)
		}
	}
}
