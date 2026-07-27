package repository

// MigrateObject migrates object schema version.
func (r *Repository) MigrateObject(
	name string,
	service *MigrationService,
	target int,
) error {

	definition, err := r.ReadObjectDefinition(name)

	if err != nil {
		return err
	}

	oldVersion := definition.Version

	err = service.Migrate(
		definition,
		target,
	)

	if err != nil {
		return err
	}

	err = r.WriteObjectDefinition(
		name,
		definition,
	)

	if err != nil {
		return err
	}

	_ = oldVersion

	return nil
}
