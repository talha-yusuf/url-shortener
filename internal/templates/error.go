package templates

import "fmt"

// ErrorTemplate returns the HTML template for error pages
func ErrorTemplate(errorMsg string) string {
	return fmt.Sprintf(`<!DOCTYPE html>
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
}
