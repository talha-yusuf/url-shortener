package templates

import (
	"fmt"

	"github.com/talha-yusuf/url-shortener/internal/models"
)

// SuccessTemplate returns the HTML template for the success page
func SuccessTemplate(shortURL *models.URL) string {
	return fmt.Sprintf(`<!DOCTYPE html>
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
}
