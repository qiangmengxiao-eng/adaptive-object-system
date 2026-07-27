package repository

// RelationQuery defines relation search conditions.
type RelationQuery struct {
	From string
	Type string
	To   string
}

// QueryRelations searches relations by conditions.
func (g *ObjectGraph) QueryRelations(query RelationQuery) []ObjectRelation {
	result := make([]ObjectRelation, 0)

	for _, relation := range g.relations {
		if query.From != "" && relation.From != query.From {
			continue
		}

		if query.Type != "" && relation.Type != query.Type {
			continue
		}

		if query.To != "" && relation.To != query.To {
			continue
		}

		result = append(result, relation)
	}

	return result
}
