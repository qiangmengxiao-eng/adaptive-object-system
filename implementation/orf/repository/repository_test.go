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
func TestReadObjectDefinition(t *testing.T) {
	fs := memory.New()

	repo := New(fs)

	definition := []byte(`
name: user
type: entity
`)

	if err := repo.CreateObject("user", definition); err != nil {
		t.Fatal(err)
	}

	object, err := repo.ReadObjectDefinition("user")
	if err != nil {
		t.Fatalf("ReadObjectDefinition returned error: %v", err)
	}

	if object.Name != "user" {
		t.Fatalf("Name = %q, want %q", object.Name, "user")
	}

	if object.Type != "entity" {
		t.Fatalf("Type = %q, want %q", object.Type, "entity")
	}
}
