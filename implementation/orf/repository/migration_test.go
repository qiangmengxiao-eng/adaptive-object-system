package repository

import "testing"

func TestMigrationEngine(t *testing.T) {

	engine := NewMigrationEngine()

	err := engine.Register(Migration{
		FromVersion: 1,
		ToVersion:   2,
	})

	if err != nil {
		t.Fatal(err)
	}

	version, err := engine.Apply(1, 2)

	if err != nil {
		t.Fatal(err)
	}

	if version != 2 {
		t.Fatalf(
			"version=%d want 2",
			version,
		)
	}
}

func TestMigrationNotFound(t *testing.T) {

	engine := NewMigrationEngine()

	_, err := engine.Apply(1, 3)

	if err == nil {
		t.Fatal(
			"expected migration error",
		)
	}
}
