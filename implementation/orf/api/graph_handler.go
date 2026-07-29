package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/repository"
)

// CreateRelationRequest represents relation creation request.
type CreateRelationRequest struct {
	From string `json:"from"`

	To string `json:"to"`

	Type string `json:"type"`
}

// handleGraph handles /graph.
func (s *Server) handleGraph(
	w http.ResponseWriter,
	r *http.Request,
) {

	switch r.Method {

	case http.MethodPost:

		s.createRelation(
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

// createRelation creates object relation.
func (s *Server) createRelation(
	w http.ResponseWriter,
	r *http.Request,
) {

	var request CreateRelationRequest

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

	relation :=
		repository.ObjectRelation{

			From: request.From,

			To: request.To,

			Type: request.Type,
		}

	err =
		s.System.GraphService.AddRelation(
			relation,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	writeJSON(
		w,
		map[string]string{

			"status": "relation added",
		},
	)
}

// handleGraphQuery handles /graph/{object}.
func (s *Server) handleGraphQuery(
	w http.ResponseWriter,
	r *http.Request,
) {

	object :=
		strings.TrimPrefix(
			r.URL.Path,
			"/graph/",
		)

	if object == "" {

		http.Error(
			w,
			"object required",
			http.StatusBadRequest,
		)

		return
	}

	relations, err :=
		s.System.GraphService.QueryRelations(
			object,
		)

	if err != nil {

		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)

		return
	}

	writeJSON(
		w,
		relations,
	)
}
