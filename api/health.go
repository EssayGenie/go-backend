package api

import (
	ghttp "go-backend/http"
	"net/http"
)

func (rs *Restful) HealthCheck(w http.ResponseWriter, _ *http.Request) error {
	return ghttp.SendJSON(w, http.StatusOK, map[string]string{
		"status":      "ok",
		"description": "EssayGenie is up and running",
	})
}
