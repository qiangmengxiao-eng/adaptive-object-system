package repository

import (
	"testing"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/memory"
)

func TestObjectSystem(t *testing.T) {
	fs := memory.New()

	repo := New(fs)

	system := NewObjectSystem(repo)

	if system.Registry == nil {
		t.Fatal("registry missing")
	}

	if system.Graph == nil {
		t.Fatal("graph missing")
	}

	if system.Behavior == nil {
		t.Fatal("behavior missing")
	}
}
