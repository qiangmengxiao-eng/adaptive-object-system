package repository

import "testing"

func TestObjectDefinitionDefaultVersion(t *testing.T) {
	definition := ObjectDefinition{
		Name: "user",
	}

	err := definition.Prepare()
	if err != nil {
		t.Fatal(err)
	}

	if definition.Version != 1 {
		t.Fatalf(
			"version=%d want 1",
			definition.Version,
		)
	}
}

func TestObjectDefinitionVersionValidation(t *testing.T) {
	definition := ObjectDefinition{
		Name:    "user",
		Version: 0,
	}

	err := definition.Prepare()
	if err != nil {
		t.Fatal(err)
	}

	if definition.Version != 1 {
		t.Fatalf(
			"version=%d want 1",
			definition.Version,
		)
	}
}
