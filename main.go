package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// Data structures
type URL struct {
	ID          int
	ShortCode   string
	OriginalURL string
	CreatedAt   time.Time
	ClickCount  int
}

// Global storage
var urlStorage = make(map[string]URL)
var nextID = 1

// HTML Templates
var homeTemplate = `<!DOCTYPE html>
<html>
<head>
    <title>URL Shortener</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 600px; margin: 50px auto; padding: 20px; }
        h1 { color: #333; }
        .form-container { background: #f5f5f5; padding: 30px; border-radius: 8px; margin: 20px 0; }
        input[type="url"] { width: 70%; padding: 10px; font-size: 16px; border: 1px solid #ddd; border-radius: 4px; }
        button { padding: 10px 20px; font-size: 16px; background: #007bff; color: white; border: none; border-radius: 4px; cursor: pointer; }
        button:hover { background: #0056b3; }
        .links { margin-top: 20px; }
        .links a { color: #007bff; text-decoration: none; margin-right: 20px; }
        .links a:hover { text-decoration: underline; }
    </style>
</head>
<body>
    <h1>URL Shortener</h1>
    <div class="form-container">
        <form method="POST" action="/create">
            <p>Enter a URL to shorten:</p>
            <input type="url" name="url" placeholder="https://example.com" required>
            <button type="submit">Shorten URL</button>
        </form>
    </div>
    
    <div class="links">
        <a href="/analytics">View Analytics</a>
        <a href="/create">Create URL (Form)</a>
    </div>
</body>
</html>`

var createFormTemplate = `<!DOCTYPE html>
<html>
<head>
    <title>Create Short URL</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 600px; margin: 50px auto; padding: 20px; }
        h1 { color: #333; }
        .form-container { background: #f5f5f5; padding: 30px; border-radius: 8px; margin: 20px 0; }
        input[type="url"] { width: 70%; padding: 10px; font-size: 16px; border: 1px solid #ddd; border-radius: 4px; }
        button { padding: 10px 20px; font-size: 16px; background: #007bff; color: white; border: none; border-radius: 4px; cursor: pointer; }
        button:hover { background: #0056b3; }
        .curl-examples { background: #e9ecef; padding: 15px; border-radius: 4px; margin-top: 20px; }
        .curl-examples h3 { margin-top: 0; }
        code { background: #f8f9fa; padding: 2px 4px; border-radius: 3px; font-family: monospace; }
        a { color: #007bff; text-decoration: none; }
        a:hover { text-decoration: underline; }
    </style>
</head>
<body>
    <h1>Create Short URL</h1>
    <div class="form-container">
        <form method="POST" action="/create">
            <p>Enter a URL to shorten:</p>
            <input type="url" name="url" placeholder="https://example.com" required>
            <button type="submit">Shorten URL</button>
        </form>
    </div>
    
    <div class="curl-examples">
        <h3>Or use curl commands:</h3>
        <p><code>curl -X POST -d 'url=https://google.com' http://localhost:8080/create</code></p>
        <p><code>curl -X POST -d 'url=https://github.com' http://localhost:8080/create</code></p>
        <p><code>curl -X POST -d 'url=https://stackoverflow.com' http://localhost:8080/create</code></p>
    </div>
    
    <p><a href="/">← Back to Home</a></p>
</body>
</html>`

// Helper functions
func createShortCode(originalURL string) URL {
	shortCode := fmt.Sprintf("abc%d", nextID)

	newURL := URL{
		ID:          nextID,
		ShortCode:   shortCode,
		OriginalURL: originalURL,
		CreatedAt:   time.Now(),
		ClickCount:  0,
	}

	urlStorage[shortCode] = newURL
	nextID++
	return newURL
}

func incrementClickCount(shortCode string) {
	if url, exists := urlStorage[shortCode]; exists {
		url.ClickCount++
		urlStorage[shortCode] = url
	}
}

func sendHTMLResponse(w http.ResponseWriter, template string, args ...interface{}) {
	w.Header().Set("Content-Type", "text/html")
	if len(args) > 0 {
		fmt.Fprintf(w, template, args...)
	} else {
		fmt.Fprint(w, template)
	}
}

func sendErrorPage(w http.ResponseWriter, errorMsg string) {
	errorHTML := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <title>Error</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 600px; margin: 50px auto; padding: 20px; }
        .error { background: #f8d7da; color: #721c24; padding: 15px; border-radius: 4px; border: 1px solid #f5c6cb; }
        a { color: #007bff; text-decoration: none; }
        a:hover { text-decoration: underline; }
    </style>
</head>
<body>
    <h1>Error</h1>
    <div class="error">
        <p>%s</p>
    </div>
    <p><a href="/create">Try again</a></p>
</body>
</html>`, errorMsg)
	
	sendHTMLResponse(w, errorHTML)
}

func sendSuccessPage(w http.ResponseWriter, shortURL URL) {
	successHTML := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <title>Short URL Created</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 600px; margin: 50px auto; padding: 20px; }
        h1 { color: #28a745; }
        .success { background: #d4edda; color: #155724; padding: 20px; border-radius: 8px; border: 1px solid #c3e6cb; }
        .url-details { background: #f8f9fa; padding: 20px; border-radius: 8px; margin: 20px 0; }
        .short-link { font-size: 20px; font-weight: bold; color: #007bff; }
        .test-button { background: #007bff; color: white; padding: 10px 20px; text-decoration: none; border-radius: 4px; display: inline-block; margin-top: 10px; }
        .test-button:hover { background: #0056b3; text-decoration: none; }
        a { color: #007bff; text-decoration: none; }
        a:hover { text-decoration: underline; }
    </style>
</head>
<body>
    <h1>Short URL Created Successfully!</h1>
    
    <div class="success">
        <p>Your URL has been shortened successfully!</p>
    </div>
    
    <div class="url-details">
        <p><strong>Original URL:</strong> <a href="%s" target="_blank">%s</a></p>
        <p><strong>Short Code:</strong> %s</p>
        <p><strong>Created At:</strong> %s</p>
        <p><strong>Short URL:</strong> <span class="short-link">http://localhost:8080/%s</span></p>
        <a href="/%s" class="test-button">Test the Short Link</a>
    </div>
    
    <p><a href="/">← Create Another URL</a> | <a href="/analytics">View Analytics</a></p>
</body>
</html>`, 
		shortURL.OriginalURL, shortURL.OriginalURL,
		shortURL.ShortCode,
		shortURL.CreatedAt.Format("2006-01-02 15:04:05"),
		shortURL.ShortCode, shortURL.ShortCode)
	
	sendHTMLResponse(w, successHTML)
}

// HTTP Handlers
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		redirectHandler(w, r)
		return
	}
	sendHTMLResponse(w, homeTemplate)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		sendHTMLResponse(w, createFormTemplate)
		return
	}

	if r.Method == "POST" {
		handleCreatePost(w, r)
		return
	}

	sendErrorPage(w, "Method not allowed")
}

func handleCreatePost(w http.ResponseWriter, r *http.Request) {
	// Parse form data
	err := r.ParseForm()
	if err != nil {
		sendErrorPage(w, "Error parsing form data")
		return
	}

	// Get URL from form
	originalURL := r.FormValue("url")
	if originalURL == "" {
		sendErrorPage(w, "URL parameter is required")
		return
	}

	// Validate URL
	if !strings.HasPrefix(originalURL, "http://") && !strings.HasPrefix(originalURL, "https://") {
		sendErrorPage(w, "URL must start with http:// or https://")
		return
	}

	// Create short URL
	shortURL := createShortCode(originalURL)

	// Send success response
	sendSuccessPage(w, shortURL)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortCode := strings.TrimPrefix(r.URL.Path, "/")

	if shortCode == "" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	url, exists := urlStorage[shortCode]
	if !exists {
		sendErrorPage(w, fmt.Sprintf("Short code '%s' not found", shortCode))
		return
	}

	incrementClickCount(shortCode)
	fmt.Printf("Redirecting %s to %s\n", shortCode, url.OriginalURL)
	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}

func analyticsHandler(w http.ResponseWriter, r *http.Request) {
	if len(urlStorage) == 0 {
		emptyHTML := `<!DOCTYPE html>
<html>
<head>
    <title>Analytics - URL Shortener</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 800px; margin: 50px auto; padding: 20px; }
        h1 { color: #333; }
        .empty-state { background: #f8f9fa; padding: 30px; border-radius: 8px; text-align: center; color: #6c757d; }
        a { color: #007bff; text-decoration: none; }
        a:hover { text-decoration: underline; }
        .create-button { background: #28a745; color: white; padding: 10px 20px; text-decoration: none; border-radius: 4px; display: inline-block; margin-top: 15px; }
        .create-button:hover { background: #218838; text-decoration: none; }
    </style>
</head>
<body>
    <h1>Analytics Dashboard</h1>
    <div class="empty-state">
        <h2>No URLs created yet!</h2>
        <p>Create your first shortened URL to see analytics here.</p>
        <a href="/" class="create-button">Create Your First URL</a>
    </div>
    <p><a href="/">← Back to Home</a></p>
</body>
</html>`
		sendHTMLResponse(w, emptyHTML)
		return
	}

	// Calculate totals
	totalClicks := 0
	for _, url := range urlStorage {
		totalClicks += url.ClickCount
	}
	avgClicks := float64(totalClicks) / float64(len(urlStorage))

	// Build table rows
	tableRows := ""
	for shortCode, url := range urlStorage {
		tableRows += fmt.Sprintf(`
            <tr>
                <td><span class="short-code">%s</span></td>
                <td class="original-url"><a href="%s" target="_blank">%s</a></td>
                <td class="date">%s</td>
                <td class="click-count">%d</td>
                <td><a href="/%s" target="_blank">Test Link</a></td>
            </tr>`,
			shortCode,
			url.OriginalURL, url.OriginalURL,
			url.CreatedAt.Format("2006-01-02 15:04"),
			url.ClickCount,
			shortCode)
	}

	// Build analytics page
	analyticsHTML := fmt.Sprintf(`<!DOCTYPE html>
<html>
<head>
    <title>Analytics - URL Shortener</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 1000px; margin: 50px auto; padding: 20px; }
        h1 { color: #333; }
        .summary { background: #e9ecef; padding: 20px; border-radius: 8px; margin-bottom: 30px; display: flex; justify-content: space-around; }
        .stat { text-align: center; }
        .stat-number { font-size: 36px; font-weight: bold; color: #007bff; }
        .stat-label { color: #6c757d; margin-top: 5px; }
        table { width: 100%%; border-collapse: collapse; background: white; border-radius: 8px; overflow: hidden; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
        th, td { padding: 15px; text-align: left; border-bottom: 1px solid #dee2e6; }
        th { background: #f8f9fa; font-weight: bold; color: #495057; }
        .short-code { font-family: monospace; background: #f8f9fa; padding: 4px 8px; border-radius: 4px; }
        .original-url { max-width: 300px; word-break: break-all; }
        .click-count { font-weight: bold; color: #28a745; }
        .date { color: #6c757d; }
        a { color: #007bff; text-decoration: none; }
        a:hover { text-decoration: underline; }
        .actions { margin-top: 20px; }
        .refresh-button { background: #17a2b8; color: white; padding: 8px 16px; text-decoration: none; border-radius: 4px; }
        .refresh-button:hover { background: #138496; text-decoration: none; }
    </style>
</head>
<body>
    <h1>Analytics Dashboard</h1>
    
    <div class="summary">
        <div class="stat">
            <div class="stat-number">%d</div>
            <div class="stat-label">Total URLs</div>
        </div>
        <div class="stat">
            <div class="stat-number">%d</div>
            <div class="stat-label">Total Clicks</div>
        </div>
        <div class="stat">
            <div class="stat-number">%.1f</div>
            <div class="stat-label">Avg Clicks/URL</div>
        </div>
    </div>
    
    <table>
        <thead>
            <tr>
                <th>Short Code</th>
                <th>Original URL</th>
                <th>Created</th>
                <th>Clicks</th>
                <th>Actions</th>
            </tr>
        </thead>
        <tbody>%s</tbody>
    </table>
    
    <div class="actions">
        <a href="/" class="refresh-button">Create New URL</a>
        <a href="/analytics" class="refresh-button">Refresh Analytics</a>
    </div>
    
    <p><a href="/">← Back to Home</a></p>
</body>
</html>`, len(urlStorage), totalClicks, avgClicks, tableRows)

	sendHTMLResponse(w, analyticsHTML)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/create", createHandler)
	http.HandleFunc("/analytics", analyticsHandler)

	fmt.Println("Server starting on http://localhost:8080")
	fmt.Println("Visit http://localhost:8080 in your browser to use the web interface!")

	log.Fatal(http.ListenAndServe(":8080", nil))
}