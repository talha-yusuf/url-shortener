package handlers

import (
	"net/http"
	"strings"

	"github.com/talha-yusuf/url-shortener/internal/storage"
	"github.com/talha-yusuf/url-shortener/internal/templates"
	"github.com/talha-yusuf/url-shortener/internal/utils"
)

// CreateHandler handles URL creation requests
type CreateHandler struct {
	storage storage.URLStorage
}

// NewCreateHandler creates a new create handler
func NewCreateHandler(storage storage.URLStorage) *CreateHandler {
	return &CreateHandler{
		storage: storage,
	}
}

// ServeHTTP handles HTTP requests for URL creation
func (h *CreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		utils.SendHTMLResponse(w, templates.CreateFormTemplate())
		return
	}

	if r.Method == "POST" {
		h.handleCreatePost(w, r)
		return
	}

	utils.SendHTMLResponse(w, templates.ErrorTemplate("Method not allowed"))
}

// handleCreatePost processes POST requests for URL creation
func (h *CreateHandler) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		utils.SendHTMLResponse(w, templates.ErrorTemplate("Error parsing form data"))
		return
	}

	// Get URL from form
	originalURL := r.FormValue("url")
	if originalURL == "" {
		utils.SendHTMLResponse(w, templates.ErrorTemplate("URL parameter is required"))
		return
	}

	// Validate URL
	if !strings.HasPrefix(originalURL, "http://") && !strings.HasPrefix(originalURL, "https://") {
		utils.SendHTMLResponse(w, templates.ErrorTemplate("URL must start with http:// or https://"))
		return
	}

	// Create short URL
	shortURL, err := h.storage.Create(originalURL)
	if err != nil {
		utils.SendHTMLResponse(w, templates.ErrorTemplate("Error creating short URL"))
		return
	}

	// Send success response
	utils.SendHTMLResponse(w, templates.SuccessTemplate(shortURL))
}
