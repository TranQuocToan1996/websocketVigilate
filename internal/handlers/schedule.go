package handlers

import (
	"github.com/TranQuocToan1996/websocketVigilate/internal/helpers"
	"net/http"
)

// ListEntries lists schedule entries
func (repo *DBRepo) ListEntries(w http.ResponseWriter, r *http.Request) {
	err := helpers.RenderPage(w, r, "schedule", nil, nil)
	if err != nil {
		printTemplateError(w, err)
	}
}
