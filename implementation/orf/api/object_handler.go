package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/repository"
)

// CreateObjectRequest represents object creation request.
type CreateObjectRequest struct {
	Name string `json:"name"`
}

// handleObjects handles /objects.
func (s *Server) handleObjects(
	w http.ResponseWriter,
	r *http.Request,
) {

	switch r.Method {

	case http.MethodPost:

		s.createObject(
			w,
			r,
		)

	default:

		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)
	}
}

// createObject creates object.
func (s *Server) createObject(
	w http.ResponseWriter,
	r *http.Request,
) {

	var request CreateObjectRequest

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

	if request.Name == "" {

		http.Error(
			w,
			"name required",
			http.StatusBadRequest,
		)

		return
	}

	definition :=
		[]byte(
			"name: " + request.Name + "\n" +
				"type: object\n" +
				"version: 1\n",
		)

	err =
		s.System.Registry.Register(
			request.Name,
			definition,
			&repository.ObjectMetadata{
				Version: 1,
			},
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	objectDefinition, err :=
		s.System.Repository.ReadObjectDefinition(
			request.Name,
		)

	if err == nil {

		s.System.Runtime.Start(
			*objectDefinition,
		)

		event :=
			repository.NewObjectEvent(
				"object.created",
				request.Name,
				"create",
				"",
			)

		_ =
			s.System.EventBus.Publish(
				event,
			)

		_ =
			s.System.Runtime.AddEvent(
				request.Name,
				event,
			)
	}

	writeJSON(
		w,
		map[string]string{

			"name": request.Name,

			"status": "created",
		},
	)
}

// handleObject handles /objects/{name}.
func (s *Server) handleObject(
	w http.ResponseWriter,
	r *http.Request,
) {

	// Phase 7 Amazon Autonomous Agent

	if strings.HasSuffix(
		r.URL.Path,
		"/amazon-agent",
	) {

		s.handleAmazonAgent(
			w,
			r,
		)

		return
	}

	// Phase 5 Amazon Listing
	if strings.HasSuffix(
		r.URL.Path,
		"/listing",
	) {

		s.handleObjectListing(
			w,
			r,
		)

		return
	}

	if strings.HasSuffix(
		r.URL.Path,
		"/listing/quality",
	) {

		s.handleListingQuality(
			w,
			r,
		)

		return
	}

	if strings.HasSuffix(
		r.URL.Path,
		"/listing/optimize",
	) {

		s.handleListingOptimize(
			w,
			r,
		)

		return
	}

	name :=
		strings.TrimPrefix(
			r.URL.Path,
			"/objects/",
		)

	if strings.HasSuffix(
		name,
		"/listing",
	) {

		s.handleObjectListing(
			w,
			r,
		)

		return
	}

	if strings.HasSuffix(
		name,
		"/events",
	) {

		s.handleObjectEvents(
			w,
			r,
		)

		return
	}

	if strings.HasSuffix(
		name,
		"/experience",
	) {

		s.handleObjectExperience(
			w,
			r,
		)

		return
	}

	if strings.HasSuffix(
		name,
		"/reflection",
	) {

		s.handleObjectReflection(
			w,
			r,
		)

		return
	}

	// Phase 3 Knowledge API
	if strings.HasSuffix(
		name,
		"/knowledge",
	) {

		s.handleObjectKnowledge(
			w,
			r,
		)

		return
	}

	if strings.HasSuffix(
		name,
		"/decision",
	) {

		s.handleObjectDecision(
			w,
			r,
		)

		return
	}

	if strings.HasSuffix(
		name,
		"/adaptations",
	) {

		s.handleObjectAdaptations(
			w,
			r,
		)

		return
	}

	if strings.HasSuffix(
		name,
		"/goal",
	) {

		s.handleObjectGoal(
			w,
			r,
		)

		return
	}

	if strings.HasSuffix(
		name,
		"/intent",
	) {

		s.handleObjectIntent(
			w,
			r,
		)

		return
	}

	if strings.HasSuffix(
		name,
		"/plan",
	) {

		s.handleObjectPlan(
			w,
			r,
		)

		return
	}

	if strings.HasSuffix(
		name,
		"/execute",
	) {

		s.handleObjectTask(
			w,
			r,
		)

		return
	}

	if strings.HasSuffix(
		name,
		"/tasks",
	) {

		s.handleObjectTask(
			w,
			r,
		)

		return
	}

	if strings.HasSuffix(
		name,
		"/generated-plans",
	) {

		s.handleObjectGeneratedPlans(
			w,
			r,
		)

		return
	}

	if strings.HasSuffix(
		name,
		"/executions",
	) {

		s.handleObjectExecutions(
			w,
			r,
		)

		return
	}

	if strings.HasSuffix(
		name,
		"/collaborate",
	) {

		s.handleObjectCollaborate(
			w,
			r,
		)

		return
	}

	if strings.HasSuffix(
		name,
		"/lifecycle",
	) {

		s.handleObjectLifecycle(
			w,
			r,
		)

		return
	}

	if strings.HasSuffix(
		name,
		"/optimization",
	) {

		s.handleObjectOptimization(
			w,
			r,
		)

		return
	}

	if name == "" {

		http.Error(
			w,
			"object name required",
			http.StatusBadRequest,
		)

		return
	}

	view, err :=
		s.System.ObjectViewService.Inspect(
			name,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusNotFound,
		)

		return
	}

	writeJSON(
		w,
		view,
	)
}
