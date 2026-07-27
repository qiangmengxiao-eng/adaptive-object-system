package repository

import (
	"testing"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/memory"
)

func TestNew(t *testing.T) {
	fs := memory.New()

	repo := New(fs)

	if repo == nil {
		t.Fatal("New returned nil")
	}
}

func TestFS(t *testing.T) {
	fs := memory.New()

	repo := New(fs)

	if repo.FS() != fs {
		t.Fatal("FS returned different filesystem")
	}
}

func TestRepositoryUsesMutableFS(t *testing.T) {
	fs := memory.New()

	repo := New(fs)

	err := repo.FS().Mkdir("/objects")
	if err != nil {
		t.Fatalf("Mkdir returned error: %v", err)
	}

	exists, err := fs.Exists("/objects")
	if err != nil {
		t.Fatal(err)
	}

	if !exists {
		t.Fatal("expected /objects to exist")
	}
}
func TestReadObject(t *testing.T) {
	fs := memory.New()

	repo := New(fs)

	err := repo.CreateObject("user", []byte("name: user"))
	if err != nil {
		t.Fatal(err)
	}

	data, err := repo.ReadObject("user")
	if err != nil {
		t.Fatalf("ReadObject returned error: %v", err)
	}

	if string(data) != "name: user" {
		t.Fatalf("data = %q, want %q", string(data), "name: user")
	}
}

func TestReadMissingObject(t *testing.T) {
	fs := memory.New()

	repo := New(fs)

	_, err := repo.ReadObject("missing")
	if err == nil {
		t.Fatal("expected error")
	}
}
