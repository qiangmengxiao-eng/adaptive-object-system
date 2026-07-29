package repository

import "gopkg.in/yaml.v3"

const KnowledgeFileName = "knowledge.yaml"

type KnowledgeStore struct {
	fs MutableGraphFS
}

type knowledgeDocument struct {
	Knowledge []Knowledge `yaml:"knowledge"`
}

func NewKnowledgeStore(
	fs MutableGraphFS,
) *KnowledgeStore {

	return &KnowledgeStore{

		fs: fs,
	}
}

// Load reads knowledge records.
func (s *KnowledgeStore) Load() (
	[]Knowledge,
	error,
) {

	data, err :=
		s.fs.ReadFile(
			KnowledgeFilePath(),
		)

	if err != nil {

		return []Knowledge{}, nil
	}

	var doc knowledgeDocument

	err =
		yaml.Unmarshal(
			data,
			&doc,
		)

	if err != nil {

		return nil, err
	}

	return doc.Knowledge, nil
}

// Append adds new knowledge.
func (s *KnowledgeStore) Append(
	item Knowledge,
) error {

	list, err :=
		s.Load()

	if err != nil {

		return err
	}

	list =
		append(
			list,
			item,
		)

	return s.save(
		list,
	)
}

// Update updates existing knowledge.
//
// Match rule:
// object + strategy
func (s *KnowledgeStore) Update(
	item Knowledge,
) error {

	list, err :=
		s.Load()

	if err != nil {

		return err
	}

	for index, current := range list {

		if current.Object ==
			item.Object &&
			current.Strategy ==
				item.Strategy {

			list[index] =
				item

			return s.save(
				list,
			)
		}
	}

	// If not found,
	// create new knowledge.
	list =
		append(
			list,
			item,
		)

	return s.save(
		list,
	)
}

// save persists knowledge.
func (s *KnowledgeStore) save(
	list []Knowledge,
) error {

	data, err :=
		yaml.Marshal(
			knowledgeDocument{
				Knowledge: list,
			},
		)

	if err != nil {

		return err
	}

	return s.fs.WriteFile(
		KnowledgeFilePath(),
		data,
	)
}

// KnowledgeFilePath returns storage path.
func KnowledgeFilePath() string {

	return "/knowledge.yaml"
}
