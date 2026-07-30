package reflection

type ReflectionHistory struct {
	Reflections []Reflection
}

func NewReflectionHistory() *ReflectionHistory {

	return &ReflectionHistory{

		Reflections: []Reflection{},
	}
}

func (h *ReflectionHistory) Add(
	reflection Reflection,
) {

	h.Reflections =
		append(
			h.Reflections,
			reflection,
		)
}
