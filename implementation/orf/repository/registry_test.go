package repository

import (
	"testing"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/memory"
)

func TestRegistry(t *testing.T) {
	fs := memory.New()

	repo := New(fs)

	registry := NewRegistry(repo)

	err := registry.Register(
		"user",
		[]byte(`
name: user
type: entity
`),
		&ObjectMetadata{
			Version: 1,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	object, err := registry.Get("user")

	if err != nil {
		t.Fatal(err)
	}

	if object.Definition.Name != "user" {
		t.Fatal("invalid object name")
	}

	list, err := registry.List()

	if err != nil {
		t.Fatal(err)
	}

	if len(list) != 1 {
		t.Fatalf("objects=%d", len(list))
	}
}
