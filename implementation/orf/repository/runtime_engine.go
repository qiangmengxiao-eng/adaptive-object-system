package repository

// RuntimeEngine manages runtime objects.
type RuntimeEngine struct {
	objects map[string]*RuntimeObject

	store *RuntimeStore
}

// NewRuntimeEngine creates runtime engine.
func NewRuntimeEngine(
	store *RuntimeStore,
) *RuntimeEngine {

	engine :=
		&RuntimeEngine{

			objects: make(
				map[string]*RuntimeObject,
			),

			store: store,
		}

	engine.LoadAll()

	return engine
}

// Start creates runtime object.
func (e *RuntimeEngine) Start(
	definition ObjectDefinition,
) *RuntimeObject {

	object :=
		NewRuntimeObject(
			definition,
		)

	e.objects[object.Name] =
		object

	_ =
		e.Save(
			object,
		)

	return object
}

// Get returns runtime object.
func (e *RuntimeEngine) Get(
	name string,
) (
	*RuntimeObject,
	bool,
) {

	object, ok :=
		e.objects[name]

	return object, ok
}

// AddEvent adds event.
func (e *RuntimeEngine) AddEvent(
	name string,
	event ObjectEvent,
) error {

	object, ok :=
		e.Get(
			name,
		)

	if !ok {

		return nil
	}

	object.AddEvent(
		event,
	)

	return e.Save(
		object,
	)
}

// Save persists all runtime objects.
func (e *RuntimeEngine) Save(
	object *RuntimeObject,
) error {

	if object != nil {

		e.objects[object.Name] =
			object
	}

	if e.store == nil {

		return nil
	}

	objects :=
		make(
			[]*RuntimeObject,
			0,
			len(e.objects),
		)

	for _, item := range e.objects {

		objects =
			append(
				objects,
				item,
			)
	}

	return e.store.SaveAll(
		objects,
	)
}

// List returns runtime objects.
func (e *RuntimeEngine) List() []*RuntimeObject {

	result :=
		make(
			[]*RuntimeObject,
			0,
			len(e.objects),
		)

	for _, object := range e.objects {

		result =
			append(
				result,
				object,
			)
	}

	return result
}

// LoadAll restores runtime objects.
func (e *RuntimeEngine) LoadAll() {

	if e.store == nil {

		return
	}

	objects, err :=
		e.store.LoadAll()

	if err != nil {

		return
	}

	for _, object := range objects {

		if object == nil {

			continue
		}

		e.objects[object.Name] =
			object
	}
}
