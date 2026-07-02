# Book API

A simple REST API for managing books built with Go, PostgreSQL, Docker, and Nginx.

## Architecture

```text
Internet
    │
    ▼
Nginx
    │
    ▼
Book API
    │
    ▼
PostgreSQL
```

## Tech Stack

* Go
* PostgreSQL
* Docker
* Docker Compose
* Nginx

## Prerequisites

* Docker
* Docker Compose

## Getting Started

1. Clone the repository.

2. Copy the environment template:

```bash
cp .env.example .env
```

3. Start the application:

```bash
docker compose up --build
```

The API will be available at:

```
http://localhost
```

## Environment Variables

The project uses the following environment variables:

| Variable      | Description              |
| ------------- | ------------------------ |
| `DB_HOST`     | PostgreSQL host          |
| `DB_PORT`     | PostgreSQL port          |
| `DB_USER`     | PostgreSQL username      |
| `DB_PASSWORD` | PostgreSQL password      |
| `DB_NAME`     | PostgreSQL database name |

See `.env.example` for default values.

## Project Structure

```
.
├── db/
│   └── init.sql
├── nginx/
│   └── nginx.conf
├── docker-compose.yml
├── Dockerfile
├── main.go
└── README.md
```
