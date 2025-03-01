# Shoto

A modern Go web application template with clean architecture and best practices.

## Features
- Clean Architecture
- RESTful API
- PostgreSQL Integration
- JWT Authentication
- Docker Support
- Hot Reload Development
- Dependency Injection
- Structured Logging

## Tech Stack
- Go 1.21+
- Echo Framework
- GORM
- PostgreSQL
- Docker & Docker Compose
- Air (Hot Reload)
- Wire (Dependency Injection)

## Prerequisites
- Go 1.21 or higher
- Docker & Docker Compose
- PostgreSQL
- Make (optional)

## Installation

1. Clone the repository
```bash 
git clone https://github.com/yourusername/shoto.git 
cd shoto
```
2. Install dependencies
```bash
go mod download
```
3. Set up environment variables
```bash
cp .env.example .env
```
4. Start Application
```bash
docker-compose up -d
```

## Development
1. Directory Structure
.
├── app/
│   └── user/
│       ├── handler.go
│       ├── service.go
│       ├── repository.go
│       └── model.go
├── pkg/
│   ├── database/
│   └── config/
├── docker-compose.yml
├── Dockerfile
└── main.go
