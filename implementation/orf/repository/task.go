package repository

import "time"

// Task represents executable object task.
type Task struct {
	Object string `json:"object"`

	Plan string `json:"plan"`

	Action string `json:"action"`

	Status string `json:"status"`

	Result string `json:"result"`

	Version int `json:"version"`

	CreatedAt time.Time `json:"created_at"`
}

// TaskEngine manages object tasks.
type TaskEngine struct {
	tasks []Task
}

// NewTaskEngine creates task engine.
func NewTaskEngine() *TaskEngine {

	return &TaskEngine{

		tasks: make(
			[]Task,
			0,
		),
	}
}

// Execute creates and executes task.
func (e *TaskEngine) Execute(
	object string,
	plan string,
	action string,
) Task {

	task :=
		Task{

			Object: object,

			Plan: plan,

			Action: action,

			Status: "completed",

			Result: "success",

			Version: 1,

			CreatedAt: time.Now(),
		}

	e.tasks =
		append(
			e.tasks,
			task,
		)

	return task
}

// Get returns object tasks.
func (e *TaskEngine) Get(
	object string,
) []Task {

	result :=
		make(
			[]Task,
			0,
		)

	for _, task := range e.tasks {

		if task.Object ==
			object {

			result =
				append(
					result,
					task,
				)
		}
	}

	return result
}
