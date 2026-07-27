package repository

import (
	"testing"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/memory"
)

func TestLoadObject(t *testing.T) {
	fs := memory.New()

	repo := New(fs)

	err := repo.CreateObject(
		"user",
		[]byte(`
name: user
type: entity
`),
	)

	if err != nil {
		t.Fatal(err)
	}

	err = repo.WriteObjectMetadata(
		"user",
		&ObjectMetadata{
			Version: 1,
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	object, err := repo.LoadObject("user")

	if err != nil {
		t.Fatal(err)
	}

	if object.Definition.Name != "user" {
		t.Fatal("invalid definition")
	}

	if object.Metadata.Version != 1 {
		t.Fatal("invalid metadata")
	}
}
