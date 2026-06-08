package handlers

import (
	"log/slog"
	"net/http"

	"github.com/thedasken/goth_template/internal/views"
)

func HelloWebHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	component := views.HelloPost(name)
	err = component.Render(r.Context(), w)
	if err != nil {
		slog.Error("failed to render HelloPost", "error", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
