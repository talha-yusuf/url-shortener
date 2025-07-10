# URL Shortener

A simple, lightweight URL shortening service built with Go. This web application allows users to create shortened URLs and provides basic analytics to track click counts.

## Features

- **URL Shortening**: Convert long URLs into short, manageable links
- **Web Interface**: Clean, responsive web UI for easy URL creation
- **Analytics Dashboard**: Track click counts and view usage statistics
- **REST API**: Programmatic access via HTTP endpoints
- **In-Memory Storage**: Fast access with no database dependencies
- **Click Tracking**: Monitor how many times each shortened URL is accessed
- **Clean Architecture**: Well-organized package structure following Go best practices

## Project Structure

```
url-shortener/
├── cmd/
│   └── server/
│       └── main.go                 # Application entry point
├── internal/
│   ├── handlers/                   # HTTP request handlers
│   │   ├── home.go                 # Home page handler
│   │   ├── create.go               # URL creation handler
│   │   ├── redirect.go             # URL redirection handler
│   │   └── analytics.go            # Analytics handler
│   ├── models/                     # Data structures
│   │   └── url.go                  # URL and Stats models
│   ├── storage/                    # Data storage layer
│   │   ├── storage.go              # Storage interface
│   │   └── memory.go               # In-memory implementation
│   ├── templates/                  # HTML templates
│   │   ├── home.go                 # Home page template
│   │   ├── create.go               # Create form template
│   │   ├── success.go              # Success page template
│   │   ├── error.go                # Error page template
│   │   └── analytics.go            # Analytics page template
│   └── utils/                      # Utility functions
│       └── response.go             # HTTP response helpers
├── go.mod                          # Go module definition
└── README.md                       # This file
```

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
   go run cmd/server/main.go
   ```

3. **Access the web interface**
   Open your browser and navigate to `http://localhost:8080`

### Alternative Commands

```bash
# Build the application
go build cmd/server/main.go

# Run the built binary
./main

# Run with specific port (if you modify the code)
PORT=3000 go run cmd/server/main.go
```

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

## Architecture

### Package Design

The application follows a clean architecture pattern with clear separation of concerns:

- **`cmd/server/`**: Application entry point and server configuration
- **`internal/handlers/`**: HTTP request handlers with dependency injection
- **`internal/models/`**: Data structures and business logic models
- **`internal/storage/`**: Data persistence layer with interface-based design
- **`internal/templates/`**: HTML template management
- **`internal/utils/`**: Shared utility functions

### Data Flow

1. **Request** → HTTP handler receives request
2. **Validation** → Handler validates input data
3. **Business Logic** → Storage layer processes the request
4. **Response** → Template renders and returns HTML response

### Storage Interface

The application uses an interface-based storage design for better testability:

```go
type URLStorage interface {
    Create(originalURL string) (*models.URL, error)
    Get(shortCode string) (*models.URL, error)
    IncrementClicks(shortCode string) error
    GetAll() ([]*models.URL, error)
    GetStats() (*models.Stats, error)
}
```

## Data Structure

The application uses the following data structures:

```go
type URL struct {
    ID          int
    ShortCode   string
    OriginalURL string
    CreatedAt   time.Time
    ClickCount  int
}

type Stats struct {
    TotalURLs     int
    TotalClicks   int
    AverageClicks float64
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
- **Architecture**: Clean architecture with dependency injection
- **Storage**: In-memory map (data is lost on server restart)
- **Server**: Built-in Go HTTP server
- **Port**: 8080 (configurable in main.go)
- **Dependencies**: Standard library only (no external dependencies)
- **Package Structure**: Follows Go project layout conventions

## Development

### Adding New Features

The modular structure makes it easy to add new features:

1. **New Handlers**: Add to `internal/handlers/`
2. **New Models**: Add to `internal/models/`
3. **New Storage**: Implement the `URLStorage` interface
4. **New Templates**: Add to `internal/templates/`

### Testing

Each package can be tested independently:

```bash
# Test all packages
go test ./...

# Test specific package
go test ./internal/handlers
go test ./internal/storage
```

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
- [ ] Unit tests for all packages
- [ ] Integration tests
- [ ] CI/CD pipeline

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Make your changes following the existing package structure
4. Add tests for new functionality
5. Test thoroughly (`go test ./...`)
6. Submit a pull request

### Code Style

- Follow Go conventions and formatting (`gofmt`)
- Use meaningful package and function names
- Add comments for exported functions
- Keep functions small and focused

## License

This project is open source and available under the [MIT License](LICENSE).

## Support

If you encounter any issues or have questions, please open an issue on the project repository.