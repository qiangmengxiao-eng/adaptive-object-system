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
func TestListObjects(t *testing.T) {
	fs := memory.New()

	repo := New(fs)

	if err := repo.CreateObject("user", []byte("name: user")); err != nil {
		t.Fatal(err)
	}

	if err := repo.CreateObject("product", []byte("name: product")); err != nil {
		t.Fatal(err)
	}

	objects, err := repo.ListObjects()
	if err != nil {
		t.Fatalf("ListObjects returned error: %v", err)
	}

	expected := []string{
		"product",
		"user",
	}

	if len(objects) != len(expected) {
		t.Fatalf("objects = %v, want %v", objects, expected)
	}

	for i := range expected {
		if objects[i] != expected[i] {
			t.Fatalf("objects = %v, want %v", objects, expected)
		}
	}
}
