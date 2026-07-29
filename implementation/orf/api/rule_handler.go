package api

import (
	"encoding/json"
	"net/http"

	"github.com/qiangmengxiao-eng/adaptive-object-system/implementation/orf/repository"
)

// HandleRules handles lifecycle rules.
func (s *Server) HandleRules(
	w http.ResponseWriter,
	r *http.Request,
) {

	switch r.Method {

	case http.MethodGet:

		rules :=
			s.System.Lifecycle.List()

		writeJSON(
			w,
			rules,
		)

	case http.MethodPost:

		var rule repository.LifecycleRule

		err :=
			json.NewDecoder(
				r.Body,
			).Decode(
				&rule,
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
			s.System.Lifecycle.Register(
				rule,
			)

		if err != nil {

			http.Error(
				w,
				err.Error(),
				http.StatusBadRequest,
			)

			return
		}

		writeJSON(
			w,
			map[string]string{

				"status": "rule created",
			},
		)

	default:

		http.Error(
			w,
			"method not allowed",
			http.StatusMethodNotAllowed,
		)

	}
}
