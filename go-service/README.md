# Go Microservice

A clean architecture Go microservice with PDF generation capabilities for student management.

## Project Structure

```
go-service/
├── cmd/
│   └── student-service/          # Entry point (main.go)
├── internal/
│   ├── config/             # Config loading from env
│   ├── handler/            # HTTP handlers
│   ├── middleware/         # cors, logging
│   ├── routes/             # HTTP routes
│   ├── service/            # Business logic
│   ├── template/           # PDF Template
│   └── types/              # Type definition for Students
├── go.mod
├── .env
├── .env.example
└── README.md
```

## Features

- Clean architecture with separation of concerns
- PDF generation for student reports
- Health check endpoint
- CORS support
- Docker support
- Structured logging
- Domain-driven design (students domain)

## Endpoints

- `GET /` - Service status
- `GET /health` - Health check
- `GET /api/v1/students/:id/report` - Generate PDF report for student

## Running Locally

```bash
# Install dependencies
go mod tidy

# Run the service (didn't test it locally)
go run cmd/report-service/main.go
```

## Running with Docker

```bash
# Build the image
docker build -t go-service .

# Run the container
docker run -p 8080:8080 go-service
```

## Testing

```bash
# Test service status
curl http://localhost:8080/

# Test health check
curl http://localhost:8080/health

# Test PDF generation
curl http://localhost:8080/api/v1/students/123/report --output student_123.pdf
```

## Environment Variables

- `PORT` - Server port (default: 8080)
- `ENV` - Environment (default: development)
- `API_URL` - Backend API URL for fetching student data

## Architecture

- **Config**: Environment configuration management
- **Middleware**: CORS and logging middleware
- **Handler**: HTTP request handlers (StudentsHandler, HealthHandler)
- **Service**: Business logic for PDF generation (StudentsService)
- **Routes**: Route definitions and setup

## Domain Structure

The service follows domain-driven design principles:

- **Students Domain**: All student-related functionality
- **Health Domain**: Service health and status endpoints
- **Future Domains**: Easy to add new domains (teachers, classes, etc.)
