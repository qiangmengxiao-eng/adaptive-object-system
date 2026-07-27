package repository

import "testing"

func TestObjectDefinitionValidate(t *testing.T) {
	definition := &ObjectDefinition{
		Name: "user",
		Type: "entity",
	}

	if err := definition.Validate(); err != nil {
		t.Fatalf("Validate returned error: %v", err)
	}
}

func TestObjectDefinitionValidateMissingName(t *testing.T) {
	definition := &ObjectDefinition{
		Type: "entity",
	}

	if err := definition.Validate(); err == nil {
		t.Fatal("expected validation error")
	}
}

func TestObjectDefinitionValidateNil(t *testing.T) {
	var definition *ObjectDefinition

	if err := definition.Validate(); err == nil {
		t.Fatal("expected validation error")
	}
}
