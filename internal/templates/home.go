package templates

// HomeTemplate returns the HTML template for the home page
func HomeTemplate() string {
	return `<!DOCTYPE html>
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
}
