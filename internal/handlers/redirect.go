package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/talha-yusuf/url-shortener/internal/storage"
	"github.com/talha-yusuf/url-shortener/internal/templates"
	"github.com/talha-yusuf/url-shortener/internal/utils"
)

// RedirectHandler handles URL redirections
type RedirectHandler struct {
	storage storage.URLStorage
}

// NewRedirectHandler creates a new redirect handler
func NewRedirectHandler(storage storage.URLStorage) *RedirectHandler {
	return &RedirectHandler{
		storage: storage,
	}
}

// ServeHTTP handles HTTP requests for URL redirections
func (h *RedirectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	shortCode := strings.TrimPrefix(r.URL.Path, "/")

	if shortCode == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	url, err := h.storage.Get(shortCode)
	if err != nil {
		utils.SendHTMLResponse(w, templates.ErrorTemplate(fmt.Sprintf("Short code '%s' not found", shortCode)))
		return
	}

	// Increment click count
	h.storage.IncrementClicks(shortCode)

	fmt.Printf("Redirecting %s to %s\n", shortCode, url.OriginalURL)
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}
