package api

import (
	"encoding/json"
	"net/http"
	"strings"
)

type AmazonAgentRequest struct {
	Product string `json:"product"`

	Features []string `json:"features"`

	Keywords []string `json:"keywords"`
}

// handleAmazonAgent executes autonomous Amazon workflow.
func (s *Server) handleAmazonAgent(
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

	if s.System.AmazonAgent == nil {

		http.Error(
			w,
			"amazon agent unavailable",
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
			"/amazon-agent",
		)

	var request AmazonAgentRequest

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
		s.System.AmazonAgent.Run(
			object,
			request.Product,
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
