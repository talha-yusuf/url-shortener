package templates

import (
	"fmt"

	"github.com/talha-yusuf/url-shortener/internal/models"
)

// AnalyticsTemplate returns the HTML template for the analytics page
func AnalyticsTemplate(stats *models.Stats, urls []*models.URL) string {
	if len(urls) == 0 {
		return emptyAnalyticsTemplate()
	}

	// Build table rows
	tableRows := ""
	for _, url := range urls {
		tableRows += fmt.Sprintf(`
            <tr>
                <td><span class="short-code">%s</span></td>
                <td class="original-url"><a href="%s" target="_blank">%s</a></td>
                <td class="date">%s</td>
                <td class="click-count">%d</td>
                <td><a href="/%s" target="_blank">Test Link</a></td>
            </tr>`,
			url.ShortCode,
			url.OriginalURL, url.OriginalURL,
			url.CreatedAt.Format("2006-01-02 15:04"),
			url.ClickCount,
			url.ShortCode)
	}

	return fmt.Sprintf(`<!DOCTYPE html>
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
</html>`, stats.TotalURLs, stats.TotalClicks, stats.AverageClicks, tableRows)
}

// emptyAnalyticsTemplate returns the HTML template for empty analytics state
func emptyAnalyticsTemplate() string {
	return `<!DOCTYPE html>
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
}
