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

## Project Goal

This repository demonstrates the evolution of a backend service into a production-like infrastructure.

The project gradually grows from a simple REST API to a fully automated and observable platform featuring monitoring, infrastructure automation, Kubernetes, and GitOps.

## Roadmap

### v1.0 — Foundation 
- [x] REST API
- [x] PostgreSQL
- [x] Docker
- [x] Docker Compose
- [x] Nginx
- [x] HTTPS
- [x] CI/CD

### v2.0 — Observability 
- [ ] Prometheus
- [ ] Grafana
- [ ] Node Exporter
- [ ] Loki
- [ ] Promtail
- [ ] Alertmanager

### v3.0 — Reliability
- [ ] Backup
- [ ] Restore
- [ ] Zero-downtime deployment

### v4.0 — Automation
- [ ] Ansible
- [ ] Infrastructure automation

### v5.0 — Kubernetes
- [ ] k3s
- [ ] Helm
- [ ] Ingress
- [ ] Cert-Manager

### v6.0 — GitOps
- [ ] ArgoCD
- [ ] Canary
- [ ] Rollback

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
