package repository

import (
	"gopkg.in/yaml.v3"
)

const AuditFileName = "audit.yaml"

// AuditStore persists audit records.
type AuditStore struct {
	fs Storage
}

// NewAuditStore creates audit store.
func NewAuditStore(
	fs Storage,
) *AuditStore {

	return &AuditStore{

		fs: fs,
	}
}

// Append adds audit record.
func (s *AuditStore) Append(
	record AuditRecord,
) error {

	records, err :=
		s.List()

	if err != nil {

		records =
			make(
				[]AuditRecord,
				0,
			)
	}

	records =
		append(
			records,
			record,
		)

	data, err :=
		yaml.Marshal(
			records,
		)

	if err != nil {

		return err
	}

	return s.fs.WriteFile(
		"/"+AuditFileName,
		data,
	)
}

// List returns audit records.
func (s *AuditStore) List() (
	[]AuditRecord,
	error,
) {

	data, err :=
		s.fs.ReadFile(
			"/" + AuditFileName,
		)

	if err != nil {

		return []AuditRecord{}, nil
	}

	var records []AuditRecord

	err =
		yaml.Unmarshal(
			data,
			&records,
		)

	if err != nil {

		return nil, err
	}

	return records, nil
}
