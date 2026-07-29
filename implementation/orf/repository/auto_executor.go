package repository

import (
	"time"

	"gopkg.in/yaml.v3"
)

// ExecutionRecord represents automatic execution result.
type ExecutionRecord struct {
	Object string `json:"object" yaml:"object"`

	Plan string `json:"plan" yaml:"plan"`

	Step string `json:"step" yaml:"step"`

	Status string `json:"status" yaml:"status"`

	Result string `json:"result" yaml:"result"`

	Version int `json:"version" yaml:"version"`

	CreatedAt time.Time `json:"created_at" yaml:"created_at"`
}

// MarshalYAML controls execution record YAML format.
func (e ExecutionRecord) MarshalYAML() (interface{}, error) {

	return struct {
		Object string `yaml:"object"`

		Plan string `yaml:"plan"`

		Step string `yaml:"step"`

		Status string `yaml:"status"`

		Result string `yaml:"result"`

		Version int `yaml:"version"`

		CreatedAt string `yaml:"created_at"`
	}{
		Object: e.Object,

		Plan: e.Plan,

		Step: e.Step,

		Status: e.Status,

		Result: e.Result,

		Version: e.Version,

		CreatedAt: e.CreatedAt.Format(
			time.RFC3339Nano,
		),
	}, nil
}

// UnmarshalYAML restores execution record from YAML.
func (e *ExecutionRecord) UnmarshalYAML(
	node *yaml.Node,
) error {

	var data struct {
		Object string `yaml:"object"`

		Plan string `yaml:"plan"`

		Step string `yaml:"step"`

		Status string `yaml:"status"`

		Result string `yaml:"result"`

		Version int `yaml:"version"`

		CreatedAt string `yaml:"created_at"`
	}

	if err :=
		node.Decode(
			&data,
		); err != nil {

		return err
	}

	t, err :=
		time.Parse(
			time.RFC3339Nano,
			data.CreatedAt,
		)

	if err != nil {

		return err
	}

	e.Object = data.Object

	e.Plan = data.Plan

	e.Step = data.Step

	e.Status = data.Status

	e.Result = data.Result

	e.Version = data.Version

	e.CreatedAt = t

	return nil
}

// AutoExecutor executes generated plans.
type AutoExecutor struct {
	Task *TaskEngine

	Store *ExecutionStore

	Records []ExecutionRecord
}

// NewAutoExecutor creates auto executor.
func NewAutoExecutor(
	task *TaskEngine,
	store *ExecutionStore,
) *AutoExecutor {

	return &AutoExecutor{

		Task: task,

		Store: store,

		Records: make(
			[]ExecutionRecord,
			0,
		),
	}
}

// Load restores execution records.
func (e *AutoExecutor) Load() {

	if e.Store == nil {

		return
	}

	records, err :=
		e.Store.Load()

	if err != nil {

		return
	}

	e.Records =
		records
}

// ExecutePlan executes generated plan.
func (e *AutoExecutor) ExecutePlan(
	object string,
	plan GeneratedPlan,
) []ExecutionRecord {

	records :=
		make(
			[]ExecutionRecord,
			0,
		)

	for _, step := range plan.Steps {

		task :=
			e.Task.Execute(
				object,
				plan.Name,
				step,
			)

		record :=
			ExecutionRecord{

				Object: object,

				Plan: plan.Name,

				Step: step,

				Status: task.Status,

				Result: task.Result,

				Version: 1,

				CreatedAt: time.Now(),
			}

		e.Records =
			append(
				e.Records,
				record,
			)

		if e.Store != nil {

			_ =
				e.Store.Save(
					e.Records,
				)
		}

		records =
			append(
				records,
				record,
			)
	}

	return records
}

// Get returns execution records.
func (e *AutoExecutor) Get(
	object string,
) []ExecutionRecord {

	result :=
		make(
			[]ExecutionRecord,
			0,
		)

	for _, record := range e.Records {

		if record.Object ==
			object {

			result =
				append(
					result,
					record,
				)
		}
	}

	return result
}
