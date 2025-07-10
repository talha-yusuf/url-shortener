package handlers

import (
	"net/http"

	"github.com/talha-yusuf/url-shortener/internal/templates"
	"github.com/talha-yusuf/url-shortener/internal/utils"
)

// HomeHandler handles the home page requests
type HomeHandler struct {
	redirectHandler *RedirectHandler
}

// NewHomeHandler creates a new home handler
func NewHomeHandler(redirectHandler *RedirectHandler) *HomeHandler {
	return &HomeHandler{
		redirectHandler: redirectHandler,
	}
}

// ServeHTTP handles HTTP requests for the home page
func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.redirectHandler.ServeHTTP(w, r)
		return
	}
	utils.SendHTMLResponse(w, templates.HomeTemplate())
}
