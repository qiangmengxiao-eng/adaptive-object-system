package repository

import "fmt"

// MigrationService provides object migration.
type MigrationService struct {
	engine *MigrationEngine
}

// NewMigrationService creates migration service.
func NewMigrationService(
	engine *MigrationEngine,
) *MigrationService {
	return &MigrationService{
		engine: engine,
	}
}

// Migrate migrates an object definition version.
func (s *MigrationService) Migrate(
	definition *ObjectDefinition,
	targetVersion int,
) error {

	if definition == nil {
		return fmt.Errorf("object definition is nil")
	}

	currentVersion := definition.Version

	if currentVersion == targetVersion {
		return nil
	}

	version, err := s.engine.Apply(
		currentVersion,
		targetVersion,
	)

	if err != nil {
		return err
	}

	definition.Version = version

	return nil
}
