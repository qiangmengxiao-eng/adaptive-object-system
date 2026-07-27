package repository

import "testing"

func TestObjectDefinitionPrepare(t *testing.T) {
	definition := &ObjectDefinition{
		Name: " user ",
	}

	if err := definition.Prepare(); err != nil {
		t.Fatalf("Prepare returned error: %v", err)
	}

	if definition.Name != "user" {
		t.Fatalf("Name = %q, want %q", definition.Name, "user")
	}

	if definition.Type != "object" {
		t.Fatalf("Type = %q, want %q", definition.Type, "object")
	}

	if definition.ID != "user" {
		t.Fatalf("ID = %q, want %q", definition.ID, "user")
	}
}

func TestObjectDefinitionPrepareMissingName(t *testing.T) {
	definition := &ObjectDefinition{
		Type: "entity",
	}

	if err := definition.Prepare(); err == nil {
		t.Fatal("expected validation error")
	}
}

func TestObjectDefinitionPrepareNil(t *testing.T) {
	var definition *ObjectDefinition

	if err := definition.Prepare(); err == nil {
		t.Fatal("expected validation error")
	}
}
