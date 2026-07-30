package planning

type PlanHistory struct {
	Plans []Plan
}

func NewPlanHistory() *PlanHistory {

	return &PlanHistory{

		Plans: []Plan{},
	}
}

func (h *PlanHistory) Add(
	plan Plan,
) {

	h.Plans =
		append(
			h.Plans,
			plan,
		)
}

func (h *PlanHistory) Latest() *Plan {

	if len(h.Plans) == 0 {

		return nil
	}

	return &h.Plans[len(h.Plans)-1]
}

func (h *PlanHistory) Count() int {

	return len(
		h.Plans,
	)
}
