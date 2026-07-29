package repository

import "testing"

func TestRuntimeEngine(t *testing.T) {

	engine :=
		NewRuntimeEngine(
			nil,
		)

	object :=
		engine.Start(
			ObjectDefinition{
				Name: "user",

				Type: "object",

				ID: "user",
			},
		)

	if object.Name != "user" {

		t.Fatal(
			"runtime object failed",
		)
	}

	found, ok :=
		engine.Get(
			"user",
		)

	if !ok {

		t.Fatal(
			"object missing",
		)
	}

	if found.State.Status != "created" {

		t.Fatal(
			"invalid state",
		)
	}
}
