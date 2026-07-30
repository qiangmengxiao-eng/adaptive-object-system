package api

import (
	"encoding/json"
	"net/http"
)

type MarketAnalysisRequest struct {
	Keyword string `json:"keyword"`

	SearchVolume int `json:"search_volume"`

	Competition int `json:"competition"`
}

// handleMarketAnalysis analyzes market opportunity.
func (s *Server) handleMarketAnalysis(
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

	if s.System.ProductOpportunity == nil {

		http.Error(
			w,
			"market engine unavailable",
			http.StatusInternalServerError,
		)

		return
	}

	var request MarketAnalysisRequest

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
		s.System.ProductOpportunity.Discover(
			request.Keyword,
			request.SearchVolume,
			request.Competition,
		)

	writeJSON(
		w,
		result,
	)
}
