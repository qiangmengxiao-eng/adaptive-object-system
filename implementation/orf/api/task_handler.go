package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/repository"
)

// CreateTaskRequest represents task execution request.
type CreateTaskRequest struct {
	Plan string `json:"plan"`

	Action string `json:"action"`
}

// handleObjectTask handles task execution API.
func (s *Server) handleObjectTask(
	w http.ResponseWriter,
	r *http.Request,
) {

	name :=
		strings.TrimPrefix(
			r.URL.Path,
			"/objects/",
		)

	name =
		strings.TrimSuffix(
			name,
			"/execute",
		)

	switch r.Method {

	case http.MethodPost:

		var request CreateTaskRequest

		err :=
			json.NewDecoder(
				r.Body,
			).Decode(
				&request,
			)

		if err != nil {

			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)

			return
		}

		task :=
			s.System.Task.Execute(
				name,
				request.Plan,
				request.Action,
			)

		event :=
			repository.NewObjectEvent(
				"task.completed",
				name,
				request.Action,
				request.Plan,
			)

		_ =
			s.System.EventBus.Publish(
				event,
			)

		_ =
			s.System.Runtime.AddEvent(
				name,
				event,
			)

		observation :=
			repository.NewObservation(
				name,
				"task.completed",
				request.Action,
				"success",
			)

		_ =
			s.System.ObservationStore.Append(
				observation,
			)

		// Phase 3:
		// Learn and promote knowledge.
		if s.System.Knowledge != nil {

			_, _ =
				s.System.Knowledge.Learn(
					name,
					request.Plan,
				)

			_ =
				s.System.Knowledge.Promote(
					name,
					request.Plan,
				)
		}

		if s.System.Decision != nil {

			decision, err :=
				s.System.Decision.Decide(
					name,
				)

			if err == nil &&
				decision != nil {

				if s.System.Planner != nil &&
					s.System.AutoExecutor != nil {

					plan :=
						s.System.Planner.Generate(
							name,
							*decision,
						)

					_ =
						s.System.AutoExecutor.ExecutePlan(
							name,
							plan,
						)
				}

				if s.System.Adaptation != nil {

					_,
						_ =
						s.System.Adaptation.Adapt(
							name,
							decision,
						)
				}
			}
		}

		writeJSON(
			w,
			task,
		)

	case http.MethodGet:

		name =
			strings.TrimSuffix(
				strings.TrimPrefix(
					r.URL.Path,
					"/objects/",
				),
				"/tasks",
			)

		tasks :=
			s.System.Task.Get(
				name,
			)

		writeJSON(
			w,
			tasks,
		)

	default:

		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
	}
}
