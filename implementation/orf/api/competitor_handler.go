package api

import (
	"encoding/json"
	"net/http"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/repository"
)

type CompetitorAnalysisRequest struct {
	Product string `json:"product"`

	Competitors []repository.Competitor `json:"competitors"`
}

func (s *Server) handleCompetitorAnalysis(
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

	if s.System.CompetitorEngine == nil {

		http.Error(
			w,
			"competitor engine unavailable",
			http.StatusInternalServerError,
		)

		return
	}

	var request CompetitorAnalysisRequest

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
		s.System.CompetitorEngine.Analyze(
			request.Product,
			request.Competitors,
		)

	writeJSON(
		w,
		result,
	)
}
