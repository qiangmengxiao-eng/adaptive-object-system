package api

import (
	"encoding/json"
	"net/http"
)

type KnowledgeRecommendRequest struct {
	Category string `json:"category"`
}

// handleKnowledgeRecommend returns learned strategies.
func (s *Server) handleKnowledgeRecommend(
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

	if s.System.StrategyRecommender == nil {

		http.Error(
			w,
			"knowledge recommender unavailable",
			http.StatusInternalServerError,
		)

		return
	}

	var request KnowledgeRecommendRequest

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

	result :=
		s.System.StrategyRecommender.Recommend(
			request.Category,
		)

	writeJSON(
		w,
		result,
	)
}
