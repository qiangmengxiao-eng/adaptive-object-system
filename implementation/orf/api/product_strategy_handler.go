package api

import (
	"encoding/json"
	"net/http"
)

type ProductStrategyRequest struct {
	Product string `json:"product"`

	MarketScore float64 `json:"market_score"`

	Competition string `json:"competition"`

	Margin float64 `json:"margin"`
}

func (s *Server) handleProductStrategy(
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

	if s.System.ProductStrategy == nil {

		http.Error(
			w,
			"product strategy unavailable",
			http.StatusInternalServerError,
		)

		return
	}

	var request ProductStrategyRequest

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
		s.System.ProductStrategy.Generate(
			request.Product,
			request.MarketScore,
			request.Competition,
			request.Margin,
		)

	writeJSON(
		w,
		result,
	)
}
