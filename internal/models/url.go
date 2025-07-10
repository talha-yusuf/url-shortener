package models

import "time"

// URL represents a shortened URL with its metadata
type URL struct {
	ID          int
	ShortCode   string
	OriginalURL string
	CreatedAt   time.Time
	ClickCount  int
}

// Stats represents analytics statistics
type Stats struct {
	TotalURLs     int
	TotalClicks   int
	AverageClicks float64
}
