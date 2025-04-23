# Smart Bridge Service

A Go-based microservice that acts as a bridge between different services, built with modern Go practices and Docker support.

## Overview

This service is designed to handle HTTP requests and act as a bridge between different services. It's built using:
- Go 1.21
- Gin web framework
- Docker support
- Logrus for logging
- Environment-based configuration

## Project Structure

```
.
├── cmd/
│   └── main.go           # Application entry point
├── pkg/
│   ├── handler/          # HTTP handlers
│   └── service/          # Business logic
├── Dockerfile            # Production Docker configuration
├── Dockerfile.local      # Local development Docker configuration
├── go.mod               # Go module definition
├── go.sum               # Go module checksums
└── server.go            # Server configuration
```

## Prerequisites

- Go 1.21 or higher
- Docker (for containerized deployment)

## Environment Variables

The service requires the following environment variables:
- `SERVER_PORT`: Port number for the HTTP server
- `SMART_BRIDGE_URL`: URL for the smart bridge service

## Getting Started

### Local Development

1. Clone the repository
2. Set up environment variables:
   ```bash
   export SERVER_PORT=8080
   export SMART_BRIDGE_URL=your_bridge_url
   ```
3. Run the service:
   ```bash
   go run cmd/main.go
   ```

### Docker Deployment

#### Local Development
```bash
docker build -f Dockerfile.local -t smartbridge-service:local .
docker run -p 8080:8080 --env-file .env smartbridge-service:local
```

#### Production
```bash
docker build -t smartbridge-service .
docker run -p 8080:8080 --env-file .env smartbridge-service
```

## Features

- HTTP server with configurable port
- Graceful shutdown handling
- JSON-formatted logging
- Docker support for both development and production
- Environment-based configuration

## API Documentation

The service exposes HTTP endpoints through the handler package. For detailed API documentation, please refer to the handler package documentation.

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details. 