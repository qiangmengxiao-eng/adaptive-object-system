package repository

// ObjectGraph stores relations between objects.
type ObjectGraph struct {
	relations []ObjectRelation
}

// NewObjectGraph creates an object graph.
func NewObjectGraph() *ObjectGraph {
	return &ObjectGraph{
		relations: make([]ObjectRelation, 0),
	}
}

// AddRelation adds a relation.
func (g *ObjectGraph) AddRelation(relation ObjectRelation) error {
	if err := relation.Validate(); err != nil {
		return err
	}

	g.relations = append(g.relations, relation)

	return nil
}

// Relations returns all relations.
func (g *ObjectGraph) Relations() []ObjectRelation {
	return g.relations
}

// FindRelations returns relations from an object.
func (g *ObjectGraph) FindRelations(from string) []ObjectRelation {
	result := make([]ObjectRelation, 0)

	for _, relation := range g.relations {
		if relation.From == from {
			result = append(result, relation)
		}
	}

	return result
}
