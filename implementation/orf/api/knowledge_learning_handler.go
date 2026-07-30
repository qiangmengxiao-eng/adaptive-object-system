package api

import (
	"encoding/json"
	"net/http"
)

type KnowledgeLearnRequest struct {
	Category string `json:"category"`

	Strategy string `json:"strategy"`

	SuccessRate float64 `json:"success_rate"`

	Confidence float64 `json:"confidence"`
}

func (s *Server) handleKnowledgeLearn(
	w http.ResponseWriter,
	r *http.Request,
) {

	if r.Method != http.MethodPost {

		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)

		return
	}

	if s.System.KnowledgeLearning == nil {

		http.Error(
			w,
			"knowledge learning unavailable",
			http.StatusInternalServerError,
		)

		return
	}

	var request KnowledgeLearnRequest

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

	err =
		s.System.KnowledgeLearning.Learn(
			"amazon-agent",
			request.Category,
			request.Strategy,
			request.SuccessRate,
			request.Confidence,
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

			"status": "learned",
		},
	)
}
