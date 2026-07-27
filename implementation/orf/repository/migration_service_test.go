package repository

import "testing"

func TestMigrationService(t *testing.T) {

	engine := NewMigrationEngine()

	err := engine.Register(Migration{
		FromVersion: 1,
		ToVersion:   2,
	})

	if err != nil {
		t.Fatal(err)
	}

	service := NewMigrationService(engine)

	definition := &ObjectDefinition{
		Name:    "user",
		Version: 1,
	}

	err = service.Migrate(
		definition,
		2,
	)

	if err != nil {
		t.Fatal(err)
	}

	if definition.Version != 2 {
		t.Fatalf(
			"version=%d want 2",
			definition.Version,
		)
	}
}
