package repository

import "testing"

func TestQueryRelations(t *testing.T) {
	graph := NewObjectGraph()

	err := graph.AddRelation(ObjectRelation{
		From: "user",
		Type: "owns",
		To:   "order",
	})

	if err != nil {
		t.Fatal(err)
	}

	result := graph.QueryRelations(RelationQuery{
		From: "user",
	})

	if len(result) != 1 {
		t.Fatalf("got %d relations", len(result))
	}

	if result[0].To != "order" {
		t.Fatalf("unexpected target: %s", result[0].To)
	}
}
