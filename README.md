# URL Shortener

A simple, lightweight URL shortening service built with Go. This web application allows users to create shortened URLs and provides basic analytics to track click counts.

## Features

- **URL Shortening**: Convert long URLs into short, manageable links
- **Web Interface**: Clean, responsive web UI for easy URL creation
- **Analytics Dashboard**: Track click counts and view usage statistics
- **REST API**: Programmatic access via HTTP endpoints
- **In-Memory Storage**: Fast access with no database dependencies
- **Click Tracking**: Monitor how many times each shortened URL is accessed

## Screenshots

### Home Page
The main interface where users can create shortened URLs through a simple form.

### Analytics Dashboard
View statistics including total URLs, total clicks, average clicks per URL, and detailed breakdown of all shortened URLs.

## Installation

### Prerequisites

- Go 1.22.0 or higher

### Quick Start

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd url-shortener
   ```

2. **Run the application**
   ```bash
   go run main.go
   ```

3. **Access the web interface**
   Open your browser and navigate to `http://localhost:8080`

## Usage

### Web Interface

1. **Create a Short URL**
   - Visit `http://localhost:8080`
   - Enter a URL in the form (must start with `http://` or `https://`)
   - Click "Shorten URL"
   - Copy the generated short URL

2. **View Analytics**
   - Click "View Analytics" from the home page
   - See statistics and detailed information about all shortened URLs

### API Usage

#### Create a Short URL

**POST** `/create`

**Form Data:**
- `url`: The original URL to shorten (required)

**Example with curl:**
```bash
curl -X POST -d 'url=https://google.com' http://localhost:8080/create
```

**Example with curl (GitHub):**
```bash
curl -X POST -d 'url=https://github.com' http://localhost:8080/create
```

**Example with curl (Stack Overflow):**
```bash
curl -X POST -d 'url=https://stackoverflow.com' http://localhost:8080/create
```

#### Access a Shortened URL

**GET** `/{shortCode}`

Simply visit `http://localhost:8080/{shortCode}` in your browser or make a GET request. The service will redirect you to the original URL.

#### View Analytics

**GET** `/analytics`

Visit `http://localhost:8080/analytics` to see the analytics dashboard.

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/` | Home page with URL shortening form |
| GET | `/create` | URL creation form page |
| POST | `/create` | Create a new shortened URL |
| GET | `/{shortCode}` | Redirect to original URL |
| GET | `/analytics` | Analytics dashboard |

## Data Structure

The application uses the following data structure for URLs:

```go
type URL struct {
    ID          int
    ShortCode   string
    OriginalURL string
    CreatedAt   time.Time
    ClickCount  int
}
```

## Features in Detail

### URL Generation
- Short codes follow the pattern `abc{ID}` (e.g., `abc1`, `abc2`, etc.)
- Automatic validation ensures URLs start with `http://` or `https://`
- Unique IDs prevent conflicts

### Analytics
- **Total URLs**: Count of all shortened URLs
- **Total Clicks**: Sum of all click counts
- **Average Clicks**: Mean clicks per URL
- **Detailed Table**: Shows each URL with creation date, click count, and test links

### Error Handling
- Invalid URLs are rejected with helpful error messages
- Non-existent short codes show appropriate error pages
- Form validation ensures required fields are provided

## Technical Details

- **Language**: Go 1.22.0
- **Storage**: In-memory map (data is lost on server restart)
- **Server**: Built-in Go HTTP server
- **Port**: 8080 (configurable in main.go)
- **Dependencies**: Standard library only (no external dependencies)

## Limitations

- **Data Persistence**: URLs are stored in memory and will be lost when the server restarts
- **Scalability**: In-memory storage limits the number of URLs that can be stored
- **No Authentication**: No user accounts or URL ownership
- **Simple Short Codes**: Basic sequential ID-based short codes

## Future Enhancements

Potential improvements for this project:

- [ ] Database integration (PostgreSQL, SQLite)
- [ ] User authentication and URL ownership
- [ ] Custom short codes
- [ ] URL expiration dates
- [ ] Rate limiting
- [ ] HTTPS support
- [ ] Docker containerization
- [ ] Configuration file support
- [ ] More detailed analytics (referrer tracking, geographic data)
- [ ] API rate limiting and authentication

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Test thoroughly
5. Submit a pull request

## License

This project is open source and available under the [MIT License](LICENSE).

## Support

If you encounter any issues or have questions, please open an issue on the project repository.