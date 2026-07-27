package repository

import (
	"testing"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/memory"
)

func TestQueryObjects(t *testing.T) {
	fs := memory.New()

	repo := New(fs)

	err := repo.CreateObject(
		"user",
		[]byte(`
name: user
type: object
`),
	)

	if err != nil {
		t.Fatal(err)
	}

	result, err := repo.QueryObjects(ObjectQuery{
		Name: "user",
	})

	if err != nil {
		t.Fatal(err)
	}

	if len(result) != 1 {
		t.Fatalf("got %d objects", len(result))
	}

	if result[0] != "user" {
		t.Fatalf("got %s", result[0])
	}
}
