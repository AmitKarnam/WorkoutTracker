# WorkoutTracker

A Go-based workout tracking application that helps users manage and track their fitness routines.

## Features

- User authentication and authorization
- Workout tracking and management
- Progress monitoring
- Email notifications
- Secure and scalable architecture

## Tech Stack

- **Backend**: Go (Golang)
- **Database**: MySQL
- **Containerization**: Docker
- **Authentication**: JWT-based
- **Email**: SMTP integration

## Project Structure

```
WorkoutTracker/
├── cmd/                 # Main execution package
├── internal/           # Core application code
│   ├── auth/          # Authentication/Authorization
│   ├── config/        # Configuration management
│   ├── controller/    # HTTP request handlers
│   ├── models/        # Data models
│   ├── repository/    # Database operations
│   ├── services/      # Business logic
│   ├── server/        # Server initialization
│   └── mailer/        # Email functionality
├── database/          # Database configurations
├── logger/            # Logging functionality
├── logs/             # Application logs
├── docker-compose.yml # Docker configuration
└── go.mod            # Go dependencies
```

## Getting Started

### Prerequisites

- Go 1.20+ (latest stable version recommended)
- MySQL 8.0+
- Docker and Docker Compose

### Installation

1. Clone the repository:
```bash
git clone https://github.com/AmitKarnam/WorkoutTracker.git
cd WorkoutTracker
```

2. Set up environment variables (copy .env.example to .env and modify as needed)

3. Set up MySQL container:

```bash
# Pull the MySQL image
docker pull mysql:8.0

# Run MySQL container
docker run --name workout-tracker-db \
    -e MYSQL_ROOT_PASSWORD=your_secure_password \
    -e MYSQL_DATABASE=workout_tracker \
    -e MYSQL_USER=your_username \
    -e MYSQL_PASSWORD=your_secure_password \
    -p 3306:3306 \
    -d mysql:8.0
```

4. Verify MySQL is running:
```bash
docker ps | grep mysql
```

5. Connect to MySQL:
```bash
docker exec -it workout-tracker-db mysql -u your_username -p
```

6. Start the application:
```bash
go run cmd/main.go
```

### Running Locally

1. Start the development server:
```bash
go run cmd/main.go
```

2. The server will be available at `http://localhost:8080`

### Running Tests

```bash
go test ./...
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

For support, please open an issue in the GitHub repository.