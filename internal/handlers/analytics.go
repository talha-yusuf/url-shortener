package handlers

import (
	"net/http"

	"github.com/talha-yusuf/url-shortener/internal/storage"
	"github.com/talha-yusuf/url-shortener/internal/templates"
	"github.com/talha-yusuf/url-shortener/internal/utils"
)

// AnalyticsHandler handles analytics page requests
type AnalyticsHandler struct {
	storage storage.URLStorage
}

// NewAnalyticsHandler creates a new analytics handler
func NewAnalyticsHandler(storage storage.URLStorage) *AnalyticsHandler {
	return &AnalyticsHandler{
		storage: storage,
	}
}

// ServeHTTP handles HTTP requests for the analytics page
func (h *AnalyticsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Get all URLs
	urls, err := h.storage.GetAll()
	if err != nil {
		utils.SendHTMLResponse(w, templates.ErrorTemplate("Error retrieving URLs"))
		return
	}

	// Get statistics
	stats, err := h.storage.GetStats()
	if err != nil {
		utils.SendHTMLResponse(w, templates.ErrorTemplate("Error calculating statistics"))
		return
	}

	// Send analytics response
	utils.SendHTMLResponse(w, templates.AnalyticsTemplate(stats, urls))
}
