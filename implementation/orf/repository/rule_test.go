package repository

import "testing"

func TestRuleEngine(t *testing.T) {

	engine :=
		NewRuleEngine()

	err :=
		engine.Register(
			ObjectRule{
				Name:   "initialize",
				Event:  "object.created",
				Action: "initialize",
			},
		)

	if err != nil {
		t.Fatal(err)
	}

	result :=
		engine.Match(
			ObjectEvent{
				Type:    "object.created",
				Object:  "user",
				Version: 1,
			},
		)

	if len(result) != 1 {

		t.Fatal(
			"rule not matched",
		)
	}
}
