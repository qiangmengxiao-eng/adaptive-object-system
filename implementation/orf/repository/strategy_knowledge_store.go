package repository

import (
	"gopkg.in/yaml.v3"
)

const StrategyKnowledgeFile = "/strategy_knowledge.yaml"

type StrategyKnowledgeDocument struct {
	Knowledge []StrategyKnowledge `yaml:"knowledge"`
}

type StrategyKnowledgeStore struct {
	fs MutableGraphFS
}

func NewStrategyKnowledgeStore(
	fs MutableGraphFS,
) *StrategyKnowledgeStore {

	return &StrategyKnowledgeStore{
		fs: fs,
	}
}

func (s *StrategyKnowledgeStore) Load() (
	[]StrategyKnowledge,
	error,
) {

	data, err :=
		s.fs.ReadFile(
			StrategyKnowledgeFile,
		)

	if err != nil {

		return []StrategyKnowledge{}, nil
	}

	var doc StrategyKnowledgeDocument

	err =
		yaml.Unmarshal(
			data,
			&doc,
		)

	return doc.Knowledge, err
}

func (s *StrategyKnowledgeStore) Append(
	item StrategyKnowledge,
) error {

	list, _ :=
		s.Load()

	list =
		append(
			list,
			item,
		)

	data, err :=
		yaml.Marshal(
			StrategyKnowledgeDocument{
				Knowledge: list,
			},
		)

	if err != nil {

		return err
	}

	return s.fs.WriteFile(
		StrategyKnowledgeFile,
		data,
	)
}
