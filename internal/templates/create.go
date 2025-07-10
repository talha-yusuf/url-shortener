package templates

// CreateFormTemplate returns the HTML template for the create form page
func CreateFormTemplate() string {
	return `<!DOCTYPE html>
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
}
