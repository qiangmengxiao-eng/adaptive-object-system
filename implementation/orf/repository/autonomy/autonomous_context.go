package autonomy

type AutonomousContext struct {
	Object string `json:"object"`

	Goal string `json:"goal"`

	State map[string]interface{} `json:"state"`

	Input interface{} `json:"input"`
}

func NewAutonomousContext(
	object string,
	goal string,
) *AutonomousContext {

	return &AutonomousContext{

		Object: object,

		Goal: goal,

		State: map[string]interface{}{},
	}
}
