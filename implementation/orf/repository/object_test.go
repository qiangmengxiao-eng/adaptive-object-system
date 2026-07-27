package repository

import (
	"testing"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/memory"
)

func TestCreateObject(t *testing.T) {
	fs := memory.New()

	repo := New(fs)

	err := repo.CreateObject("user", []byte("name: user"))
	if err != nil {
		t.Fatalf("CreateObject returned error: %v", err)
	}

	data, err := fs.ReadFile("/objects/user/definition.yaml")
	if err != nil {
		t.Fatalf("definition file missing: %v", err)
	}

	if string(data) != "name: user" {
		t.Fatalf("data = %q, want %q", string(data), "name: user")
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

func TestDeleteObject(t *testing.T) {
	fs := memory.New()

	repo := New(fs)

	err := repo.CreateObject("user", []byte("name: user"))
	if err != nil {
		t.Fatal(err)
	}

	err = repo.DeleteObject("user")
	if err != nil {
		t.Fatalf("DeleteObject returned error: %v", err)
	}

	exists, err := fs.Exists("/objects/user")
	if err != nil {
		t.Fatal(err)
	}

	if exists {
		t.Fatal("object should be deleted")
	}

	exists, err = fs.Exists("/objects/user/definition.yaml")
	if err != nil {
		t.Fatal(err)
	}

	if exists {
		t.Fatal("definition file should be deleted")
	}
}

func TestDeleteMissingObject(t *testing.T) {
	fs := memory.New()

	repo := New(fs)

	err := repo.DeleteObject("missing")
	if err == nil {
		t.Fatal("expected error")
	}
}
