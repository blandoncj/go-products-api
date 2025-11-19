# Go Products API - Microservices Architecture

A production-ready microservices application built with Go, implementing CRUD operations for product management using a distributed architecture pattern.

## 📋 Table of Contents

- [Overview](#overview)
- [Architecture](#architecture)
- [Features](#features)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Documentation](#api-documentation)
- [Testing](#testing)
- [Docker Deployment](#docker-deployment)
- [CI/CD Pipeline](#cicd-pipeline)
- [Database Backup](#database-backup)
- [Project Structure](#project-structure)
- [Contributing](#contributing)
- [License](#license)

## 🎯 Overview

This project demonstrates a microservices architecture in Go, splitting CRUD operations across four independent services. Each service is containerized, orchestrated with Docker Compose, and includes automated testing, security scanning, and continuous deployment via GitHub Actions.

### Key Technologies

- **Language**: Go 1.25.3
- **Database**: MongoDB
- **Containerization**: Docker & Docker Compose
- **CI/CD**: GitHub Actions
- **Testing**: Testify, GoMock
- **Security**: Gosec
- **Registry**: GitHub Container Registry (GHCR) / Docker Hub

## 🏗️ Architecture

The application follows a microservices architecture with domain-driven design principles:

```
┌─────────────────────────────────────────────────────────┐
│                    API Gateway / Client                  │
└────────┬────────┬────────┬────────┬─────────────────────┘
         │        │        │        │
    ┌────▼───┐ ┌─▼────┐ ┌─▼────┐ ┌─▼──────┐
    │ Create │ │ Read │ │Update│ │ Delete │
    │Service │ │Service│ │Service│ │Service│
    └────┬───┘ └──┬───┘ └──┬───┘ └───┬────┘
         │        │        │          │
         └────────┴────────┴──────────┘
                     │
              ┌──────▼──────┐
              │   MongoDB   │
              └─────────────┘
```

### Service Responsibilities

- **Create Service** (Port 8081): Handles product creation
- **Read Service** (Port 8082): Retrieves product information
- **Update Service** (Port 8083): Updates existing products
- **Delete Service** (Port 8084): Removes products

Each service follows a three-layer architecture:

- **Controller Layer**: HTTP request handling and validation
- **Service Layer**: Business logic implementation
- **Repository Layer**: Data access and MongoDB operations

## ✨ Features

- ✅ Microservices-based CRUD operations
- ✅ MongoDB integration with connection pooling
- ✅ Docker containerization with multi-stage builds
- ✅ Health check endpoints for all services
- ✅ Environment-based configuration
- ✅ Comprehensive unit and integration tests
- ✅ Automated CI/CD pipeline
- ✅ Security scanning with Gosec
- ✅ Automated database backups
- ✅ Test coverage reporting
- ✅ Docker Compose orchestration
- ✅ Volume persistence for data
- ✅ Service dependency management

## 📦 Prerequisites

Before running this project, ensure you have the following installed:

- **Go**: 1.20 or higher
- **Docker**: 20.10 or higher
- **Docker Compose**: 2.0 or higher
- **Make** (optional, for convenience commands)
- **Git**: For version control

## 🚀 Installation

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/go-products-api.git
cd go-products-api
```

### 2. Environment Configuration

Create a `.env` file in the root directory:

```env
# MongoDB Configuration
MONGO_INITDB_ROOT_USERNAME=admin
MONGO_INITDB_ROOT_PASSWORD=secure_password_here
MONGO_HOST=mongodb
MONGO_PORT=27017
MONGO_DATABASE=products_db

# Service Ports
CREATE_SERVICE_PORT=8081
READ_SERVICE_PORT=8082
UPDATE_SERVICE_PORT=8083
DELETE_SERVICE_PORT=8084

# Application Settings
APP_ENV=development
LOG_LEVEL=info
```

### 3. Install Dependencies

For local development:

```bash
# Root module
go mod download

# Individual services
cd services/create-service && go mod download
cd ../read-service && go mod download
cd ../update-service && go mod download
cd ../delete-service && go mod download
```

## ⚙️ Configuration

### Docker Compose Configuration

The `docker-compose.yml` file defines the entire stack:

- **Networks**: Internal bridge network for service communication
- **Volumes**:
  - `mongodb_data`: Persistent MongoDB storage
  - `mongodb_backup`: Database backup storage
- **Health Checks**: All services include health check endpoints
- **Dependencies**: Services wait for MongoDB to be healthy before starting

### Service Configuration

Each service can be configured via environment variables:

| Variable         | Description               | Default                                  |
| ---------------- | ------------------------- | ---------------------------------------- |
| `MONGO_URI`      | MongoDB connection string | `mongodb://admin:password@mongodb:27017` |
| `MONGO_DATABASE` | Database name             | `products_db`                            |
| `PORT`           | Service port              | Varies by service                        |
| `LOG_LEVEL`      | Logging verbosity         | `info`                                   |

## 💻 Usage

### Running with Docker Compose

#### Start All Services

```bash
docker-compose up -d
```

#### View Logs

```bash
# All services
docker-compose logs -f

# Specific service
docker-compose logs -f create-service
```

#### Stop Services

```bash
docker-compose down
```

#### Rebuild After Changes

```bash
docker-compose up -d --build
```

### Running Locally (Development)

```bash
# Terminal 1: Start MongoDB
docker run -d -p 27017:27017 \
  -e MONGO_INITDB_ROOT_USERNAME=admin \
  -e MONGO_INITDB_ROOT_PASSWORD=password \
  mongo:latest

# Terminal 2: Create Service
cd services/create-service
go run cmd/main.go

# Terminal 3: Read Service
cd services/read-service
go run cmd/main.go

# Terminal 4: Update Service
cd services/update-service
go run cmd/main.go

# Terminal 5: Delete Service
cd services/delete-service
go run cmd/main.go
```

## 📚 API Documentation

### Create Service (Port 8081)

#### Create Product

```http
POST /products
Content-Type: application/json

{
  "name": "Product Name",
  "description": "Product Description",
  "price": 99.99,
  "category": "Electronics",
  "stock": 100
}
```

**Response:**

```json
{
  "id": "507f1f77bcf86cd799439011",
  "name": "Product Name",
  "description": "Product Description",
  "price": 99.99,
  "category": "Electronics",
  "stock": 100,
  "created_at": "2024-11-18T10:30:00Z",
  "updated_at": "2024-11-18T10:30:00Z"
}
```

### Read Service (Port 8082)

#### Get All Products

```http
GET /products
```

#### Get Product by ID

```http
GET /products/{id}
```

**Response:**

```json
{
  "id": "507f1f77bcf86cd799439011",
  "name": "Product Name",
  "description": "Product Description",
  "price": 99.99,
  "category": "Electronics",
  "stock": 100,
  "created_at": "2024-11-18T10:30:00Z",
  "updated_at": "2024-11-18T10:30:00Z"
}
```

### Update Service (Port 8083)

#### Update Product

```http
PUT /products/{id}
Content-Type: application/json

{
  "name": "Updated Product Name",
  "price": 89.99,
  "stock": 150
}
```

### Delete Service (Port 8084)

#### Delete Product

```http
DELETE /products/{id}
```

**Response:**

```json
{
  "message": "Product deleted successfully",
  "id": "507f1f77bcf86cd799439011"
}
```

### Health Check (All Services)

```http
GET /health
```

**Response:**

```json
{
  "status": "healthy",
  "service": "create-service",
  "timestamp": "2024-11-18T10:30:00Z"
}
```

## 🧪 Testing

### Run All Tests

```bash
# From root directory
go test ./...

# With coverage
go test -cover ./...

# With verbose output
go test -v ./...

# Generate coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

### Test Individual Services

```bash
# Create Service
cd services/create-service
go test ./... -v

# Read Service
cd services/read-service
go test ./... -v

# Update Service
cd services/update-service
go test ./... -v

# Delete Service
cd services/delete-service
go test ./... -v
```

### Integration Tests

```bash
# Ensure services are running first
docker-compose up -d

# Run integration tests
go test -tags=integration ./tests/integration/...
```

### Test Coverage Requirements

- Unit Tests: Minimum 80% coverage
- Integration Tests: Critical paths covered
- Mock external dependencies using GoMock

## 🐳 Docker Deployment

### Build Individual Images

```bash
# Create Service
docker build -t products-create:latest ./services/create-service

# Read Service
docker build -t products-read:latest ./services/read-service

# Update Service
docker build -t products-update:latest ./services/update-service

# Delete Service
docker build -t products-delete:latest ./services/delete-service
```

### Multi-Stage Build Optimization

Each Dockerfile uses multi-stage builds:

- **Stage 1**: Build the Go binary
- **Stage 2**: Create minimal runtime image with binary

Benefits:

- Smaller image sizes (~20MB vs ~800MB)
- Improved security (fewer attack vectors)
- Faster deployment times

### Docker Ignore

Each service includes a `.dockerignore` file:

```
.git
.gitignore
README.md
.env
*.md
tests/
coverage.out
*.test
```

## 🔄 CI/CD Pipeline

The GitHub Actions workflow (`.github/workflows/ci-cd.yml`) automates:

### On Pull Request

1. **Linting**: Code style checks
2. **Unit Tests**: Run all unit tests
3. **Integration Tests**: Test service interactions
4. **Security Scan**: Gosec vulnerability scanning
5. **Coverage Report**: Generate and upload coverage

### On Push to Main

1. All PR checks
2. **Build Images**: Create Docker images for all services
3. **Tag Images**: Version tagging (semantic versioning)
4. **Push to Registry**: Publish to GHCR/Docker Hub
5. **Create Release**: Automated GitHub releases

### Workflow Example

```yaml
name: CI/CD Pipeline

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v3
      - uses: actions/setup-go@v4
        with:
          go-version: "1.25"

      - name: Run tests
        run: go test -v -cover ./...

      - name: Security scan
        run: |
          go install github.com/securego/gosec/v2/cmd/gosec@latest
          gosec ./...

  build-and-push:
    needs: test
    runs-on: ubuntu-latest
    if: github.event_name == 'push'
    steps:
      - name: Build and push images
        run: |
          docker compose build
          docker compose push
```

## 💾 Database Backup

### Automated Backup Script

The `scripts/backup.sh` script performs MongoDB backups:

```bash
#!/bin/bash
TIMESTAMP=$(date +"%Y%m%d-%H%M")
BACKUP_NAME="backup-${TIMESTAMP}"

docker exec mongodb mongodump \
  --username=$MONGO_INITDB_ROOT_USERNAME \
  --password=$MONGO_INITDB_ROOT_PASSWORD \
  --authenticationDatabase=admin \
  --out=/backup/$BACKUP_NAME

echo "Backup completed: $BACKUP_NAME"
```

### Running Backups

```bash
# Manual backup
./scripts/backup.sh

# Schedule with cron (daily at 2 AM)
0 2 * * * /path/to/scripts/backup.sh >> /var/log/mongo-backup.log 2>&1
```

### Restore from Backup

```bash
BACKUP_DATE="20241118-0200"

docker exec mongodb mongorestore \
  --username=admin \
  --password=password \
  --authenticationDatabase=admin \
  /backup/backup-${BACKUP_DATE}
```

## 📁 Project Structure

```
go-products-api/
├── pkg/
│   └── model/
│       └── product.go              # Shared product model
├── services/
│   ├── create-service/
│   │   ├── cmd/
│   │   │   └── main.go            # Service entry point
│   │   ├── internal/
│   │   │   ├── controller/        # HTTP handlers
│   │   │   ├── service/           # Business logic
│   │   │   └── repository/        # Data access
│   │   ├── Dockerfile
│   │   ├── .dockerignore
│   │   ├── go.mod
│   │   └── go.sum
│   ├── read-service/              # Similar structure
│   ├── update-service/            # Similar structure
│   └── delete-service/            # Similar structure
├── scripts/
│   └── backup.sh                  # MongoDB backup script
├── docker-compose.yml             # Orchestration config
├── .env.example                   # Environment template
├── .gitignore
├── go.mod                         # Root module
├── go.sum
├── LICENSE
└── README.md
```

## 🤝 Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Development Guidelines

- Follow Go best practices and idioms
- Write tests for new features
- Update documentation as needed
- Ensure all tests pass before submitting PR
- Keep commits atomic and well-described

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Go community for excellent tooling
- MongoDB team for robust database
- Docker for containerization platform
- GitHub Actions for CI/CD automation

---

**Maintained by**: Your Name  
**Repository**: <https://github.com/yourusername/go-products-api>  
**Issues**: <https://github.com/yourusername/go-products-api/issues>
