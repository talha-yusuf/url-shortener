package storage

import (
	"fmt"
	"time"

	"github.com/talha-yusuf/url-shortener/internal/models"
)

// MemoryStorage implements URLStorage using in-memory storage
type MemoryStorage struct {
	urlStorage map[string]*models.URL
	nextID     int
}

// NewMemoryStorage creates a new memory storage instance
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		urlStorage: make(map[string]*models.URL),
		nextID:     1,
	}
}

// Create generates a new short URL and stores it
func (m *MemoryStorage) Create(originalURL string) (*models.URL, error) {
	shortCode := fmt.Sprintf("abc%d", m.nextID)

	newURL := &models.URL{
		ID:          m.nextID,
		ShortCode:   shortCode,
		OriginalURL: originalURL,
		CreatedAt:   time.Now(),
		ClickCount:  0,
	}

	m.urlStorage[shortCode] = newURL
	m.nextID++
	return newURL, nil
}

// Get retrieves a URL by its short code
func (m *MemoryStorage) Get(shortCode string) (*models.URL, error) {
	url, exists := m.urlStorage[shortCode]
	if !exists {
		return nil, fmt.Errorf("short code '%s' not found", shortCode)
	}
	return url, nil
}

// IncrementClicks increases the click count for a URL
func (m *MemoryStorage) IncrementClicks(shortCode string) error {
	url, exists := m.urlStorage[shortCode]
	if !exists {
		return fmt.Errorf("short code '%s' not found", shortCode)
	}
	url.ClickCount++
	return nil
}

// GetAll returns all stored URLs
func (m *MemoryStorage) GetAll() ([]*models.URL, error) {
	urls := make([]*models.URL, 0, len(m.urlStorage))
	for _, url := range m.urlStorage {
		urls = append(urls, url)
	}
	return urls, nil
}

// GetStats calculates and returns analytics statistics
func (m *MemoryStorage) GetStats() (*models.Stats, error) {
	totalClicks := 0
	for _, url := range m.urlStorage {
		totalClicks += url.ClickCount
	}

	var avgClicks float64
	if len(m.urlStorage) > 0 {
		avgClicks = float64(totalClicks) / float64(len(m.urlStorage))
	}

	return &models.Stats{
		TotalURLs:     len(m.urlStorage),
		TotalClicks:   totalClicks,
		AverageClicks: avgClicks,
	}, nil
}
