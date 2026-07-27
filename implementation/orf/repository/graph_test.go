package repository

import "testing"

func TestObjectGraph(t *testing.T) {
	graph := NewObjectGraph()

	err := graph.AddRelation(ObjectRelation{
		From: "user",
		To:   "order",
		Type: "owns",
	})

	if err != nil {
		t.Fatal(err)
	}

	relations := graph.Relations()

	if len(relations) != 1 {
		t.Fatalf("relations=%d", len(relations))
	}

	if relations[0].To != "order" {
		t.Fatal("invalid relation")
	}
}

func TestFindRelations(t *testing.T) {
	graph := NewObjectGraph()

	graph.AddRelation(ObjectRelation{
		From: "user",
		To:   "order",
		Type: "owns",
	})

	graph.AddRelation(ObjectRelation{
		From: "user",
		To:   "profile",
		Type: "has",
	})

	result := graph.FindRelations("user")

	if len(result) != 2 {
		t.Fatalf("relations=%d", len(result))
	}
}
