package storage

import "github.com/talha-yusuf/url-shortener/internal/models"

// URLStorage defines the interface for URL storage operations
type URLStorage interface {
	Create(originalURL string) (*models.URL, error)
	Get(shortCode string) (*models.URL, error)
	IncrementClicks(shortCode string) error
	GetAll() ([]*models.URL, error)
	GetStats() (*models.Stats, error)
}
