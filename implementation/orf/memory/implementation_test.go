package memory

import (
	"bytes"
	"testing"
)

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
		{"root exists", "/", true},
		{"missing file", "/README.md", false},
		{"missing directory", "/docs", false},
		{"relative path", "docs", false},
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

func TestReadDir(t *testing.T) {
	fs := New()

	if err := fs.addNode("/README.md", false, nil); err != nil {
		t.Fatal(err)
	}

	if err := fs.addNode("/docs", true, nil); err != nil {
		t.Fatal(err)
	}

	if err := fs.addNode("/LICENSE", false, nil); err != nil {
		t.Fatal(err)
	}

	entries, err := fs.ReadDir("/")
	if err != nil {
		t.Fatalf("ReadDir returned error: %v", err)
	}

	if len(entries) != 3 {
		t.Fatalf("len(entries) = %d, want 3", len(entries))
	}

	want := []struct {
		name        string
		path        string
		isDirectory bool
	}{
		{"LICENSE", "/LICENSE", false},
		{"README.md", "/README.md", false},
		{"docs", "/docs", true},
	}

	for i := range want {
		if entries[i].Name != want[i].name {
			t.Fatalf("entries[%d].Name = %q, want %q",
				i,
				entries[i].Name,
				want[i].name,
			)
		}

		if entries[i].Path != want[i].path {
			t.Fatalf("entries[%d].Path = %q, want %q",
				i,
				entries[i].Path,
				want[i].path,
			)
		}

		if entries[i].IsDirectory != want[i].isDirectory {
			t.Fatalf("entries[%d].IsDirectory = %v, want %v",
				i,
				entries[i].IsDirectory,
				want[i].isDirectory,
			)
		}
	}
}

func TestReadDirEmpty(t *testing.T) {
	fs := New()

	if err := fs.addNode("/docs", true, nil); err != nil {
		t.Fatal(err)
	}

	entries, err := fs.ReadDir("/docs")
	if err != nil {
		t.Fatalf("ReadDir returned error: %v", err)
	}

	if len(entries) != 0 {
		t.Fatalf("len(entries) = %d, want 0", len(entries))
	}
}

func TestReadDirMissing(t *testing.T) {
	fs := New()

	_, err := fs.ReadDir("/missing")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestReadDirFile(t *testing.T) {
	fs := New()

	if err := fs.addNode("/README.md", false, nil); err != nil {
		t.Fatal(err)
	}

	_, err := fs.ReadDir("/README.md")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestReadFile(t *testing.T) {
	fs := New()

	content := []byte("hello world")

	if err := fs.addNode("/hello.txt", false, content); err != nil {
		t.Fatal(err)
	}

	data, err := fs.ReadFile("/hello.txt")
	if err != nil {
		t.Fatalf("ReadFile returned error: %v", err)
	}

	if !bytes.Equal(data, content) {
		t.Fatalf("ReadFile = %q, want %q", data, content)
	}
}

func TestReadFileMissing(t *testing.T) {
	fs := New()

	_, err := fs.ReadFile("/missing.txt")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestReadFileDirectory(t *testing.T) {
	fs := New()

	if err := fs.addNode("/docs", true, nil); err != nil {
		t.Fatal(err)
	}

	_, err := fs.ReadFile("/docs")
	if err == nil {
		t.Fatal("expected error")
	}
}

func TestReadFileReturnsCopy(t *testing.T) {
	fs := New()

	if err := fs.addNode("/hello.txt", false, []byte("hello")); err != nil {
		t.Fatal(err)
	}

	data, err := fs.ReadFile("/hello.txt")
	if err != nil {
		t.Fatal(err)
	}

	data[0] = 'X'

	again, err := fs.ReadFile("/hello.txt")
	if err != nil {
		t.Fatal(err)
	}

	if string(again) != "hello" {
		t.Fatalf("expected stored data to remain %q, got %q", "hello", string(again))
	}
}

func TestReadFileEmpty(t *testing.T) {
	fs := New()

	if err := fs.addNode("/empty.txt", false, []byte{}); err != nil {
		t.Fatal(err)
	}

	data, err := fs.ReadFile("/empty.txt")
	if err != nil {
		t.Fatal(err)
	}

	if len(data) != 0 {
		t.Fatalf("len(data) = %d, want 0", len(data))
	}
}
