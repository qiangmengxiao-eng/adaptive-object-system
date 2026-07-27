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
