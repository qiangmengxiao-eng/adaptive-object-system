package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

type SellerAnalysisRequest struct {
	Product string `json:"product"`

	Cost float64 `json:"cost"`

	Price float64 `json:"price"`

	Features []string `json:"features"`

	Keywords []string `json:"keywords"`
}

// handleSellerAnalysis analyzes Amazon product opportunity.
func (s *Server) handleSellerAnalysis(
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

	if s.System.SellerIntelligence == nil {

		http.Error(
			w,
			"seller intelligence unavailable",
			http.StatusInternalServerError,
		)

		return
	}

	object :=
		strings.TrimPrefix(
			r.URL.Path,
			"/objects/",
		)

	object =
		strings.TrimSuffix(
			object,
			"/seller-analysis",
		)

	var request SellerAnalysisRequest

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

	result, err :=
		s.System.SellerIntelligence.Analyze(
			object,
			request.Product,
			request.Cost,
			request.Price,
			request.Features,
			request.Keywords,
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
		result,
	)
}
