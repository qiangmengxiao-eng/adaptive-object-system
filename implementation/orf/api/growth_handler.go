package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/repository"
)

type GrowthRequest struct {
	Product string `json:"product"`

	Keywords []string `json:"keywords"`

	Competitors []repository.Competitor `json:"competitors"`
}

// handleGrowthAnalysis analyzes growth strategy.
func (s *Server) handleGrowthAnalysis(
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

	if s.System.GrowthEngine == nil {

		http.Error(
			w,
			"growth engine unavailable",
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
			"/growth-analysis",
		)

	_ = object

	var request GrowthRequest

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
		s.System.GrowthEngine.Analyze(
			request.Product,
			request.Competitors,
			request.Keywords,
		)

	writeJSON(
		w,
		result,
	)
}
