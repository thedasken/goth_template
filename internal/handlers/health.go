package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/thedasken/goth_template/internal/database"
)

func HealthHandler(db database.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonResp, _ := json.Marshal(db.Health())
		_, _ = w.Write(jsonResp)
	}
}
