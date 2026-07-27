package repository

import "fmt"

// Migration describes a schema migration.
type Migration struct {
	FromVersion int `yaml:"from_version"`

	ToVersion int `yaml:"to_version"`
}

// Validate checks migration validity.
func (m *Migration) Validate() error {
	if m == nil {
		return fmt.Errorf("migration is nil")
	}

	if m.FromVersion < 1 {
		return fmt.Errorf("invalid source version")
	}

	if m.ToVersion <= m.FromVersion {
		return fmt.Errorf("target version must be greater than source version")
	}

	return nil
}

// MigrationEngine manages schema migrations.
type MigrationEngine struct {
	migrations []Migration
}

// NewMigrationEngine creates a migration engine.
func NewMigrationEngine() *MigrationEngine {
	return &MigrationEngine{
		migrations: make([]Migration, 0),
	}
}

// Register adds a migration.
func (e *MigrationEngine) Register(migration Migration) error {
	if err := migration.Validate(); err != nil {
		return err
	}

	e.migrations = append(
		e.migrations,
		migration,
	)

	return nil
}

// Find finds a migration path.
func (e *MigrationEngine) Find(
	from int,
	to int,
) (*Migration, bool) {

	for _, migration := range e.migrations {
		if migration.FromVersion == from &&
			migration.ToVersion == to {

			result := migration
			return &result, true
		}
	}

	return nil, false
}

// Apply applies a migration.
func (e *MigrationEngine) Apply(
	version int,
	target int,
) (int, error) {

	migration, ok := e.Find(version, target)

	if !ok {
		return version, fmt.Errorf(
			"migration not found: %d -> %d",
			version,
			target,
		)
	}

	return migration.ToVersion, nil
}
