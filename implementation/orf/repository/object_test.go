package repository

import (
	"testing"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/memory"
)

func TestCreateObject(t *testing.T) {
	fs := memory.New()

	repo := New(fs)

	definition := []byte("name: user")

	if err := repo.CreateObject("user", definition); err != nil {
		t.Fatalf("CreateObject returned error: %v", err)
	}

	data, err := fs.ReadFile("/objects/user/definition.yaml")
	if err != nil {
		t.Fatalf("ReadFile returned error: %v", err)
	}

	if string(data) != "name: user" {
		t.Fatalf("content = %q, want %q", string(data), "name: user")
	}
}

func TestCreateObjectCreatesDirectory(t *testing.T) {
	fs := memory.New()

	repo := New(fs)

	if err := repo.CreateObject("user", []byte("name: user")); err != nil {
		t.Fatal(err)
	}

	exists, err := fs.Exists("/objects/user")
	if err != nil {
		t.Fatal(err)
	}

	if !exists {
		t.Fatal("expected object directory to exist")
	}
}

func TestCreateObjectDuplicate(t *testing.T) {
	fs := memory.New()

	repo := New(fs)

	if err := repo.CreateObject("user", []byte("first")); err != nil {
		t.Fatal(err)
	}

	if err := repo.CreateObject("user", []byte("second")); err == nil {
		t.Fatal("expected duplicate object creation to fail")
	}
}
