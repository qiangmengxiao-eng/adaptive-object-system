package repository

import (
	"testing"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/memory"
)

func TestRepositoryUsesMutableFS(t *testing.T) {
	fs := memory.New()

	repo := New(fs)

	if repo.FS() != fs {
		t.Fatal("repository should keep the provided filesystem")
	}

	if err := fs.Mkdir("/objects"); err != nil {
		t.Fatal(err)
	}

	exists, err := fs.Exists("/objects")
	if err != nil {
		t.Fatal(err)
	}

	if !exists {
		t.Fatal("expected /objects to exist")
	}
}
