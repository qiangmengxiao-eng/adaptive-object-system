package repository

import "testing"

func TestBehaviorEngine(t *testing.T) {
	engine := NewBehaviorEngine()

	err := engine.Register(ObjectBehavior{
		Name:   "validate",
		Action: "validate",
	})

	if err != nil {
		t.Fatal(err)
	}

	behavior, ok := engine.Get("validate")

	if !ok {
		t.Fatal("behavior not found")
	}

	if behavior.Action != "validate" {
		t.Fatal("invalid behavior")
	}
}
