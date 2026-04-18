# Relay Engine

Relay Engine is a lightweight, distributed workflow and task execution engine built with Go and PostgreSQL. It allows you to define complex workflows and execute them with built-in task locking, retries, and persistence.

## Key Features

- **Workflow Definitions**: Define structured workflows with JSON definitions.
- **Distributed Workers**: Scale task processing by running multiple worker instances.
- **Task Locking**: Ensures tasks are processed by only one worker at a time using database-level locking.
- **Retry Mechanism**: Built-in support for task retries with configurable limits.
- **Persistence**: Full execution history and state management via PostgreSQL.
- **REST API**: Simple interface for managing workflows and starting executions.

## Tech Stack

- **Languange**: Go (Golang)
- **Framework**: Gin Gonic (Web)
- **Database**: PostgreSQL
- **ORM**: GORM
- **Containerization**: Docker & Docker Compose

## Project Structure

```text
.
├── cmd/
│   ├── api/          # Entry point for the REST API
│   └── worker/       # Entry point for the Background Worker
├── internal/
│   ├── database/     # DB connection and setup
│   ├── handlers/     # HTTP and Worker handlers
│   ├── models/       # Database schemas (GORM)
│   ├── repositories/ # Data access layer
│   └── services/     # Business logic
├── Dockerfile        # Multi-stage build for API and Worker
└── docker-compose.yml # Full stack orchestration
```

## Getting Started

### Prerequisites

- [Docker](https://www.docker.com/get-started) and Docker Compose installed.

### Running with Docker

The easiest way to get started is using Docker Compose, which sets up the API, Worker, and PostgreSQL database automatically.

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd relay-engine
   ```

2. **Start the services**:
   ```bash
   docker compose up --build
   ```

The API will be available at `http://localhost:8000`.

## API Documentation

### Workflows

#### Create a Workflow
`POST /api/v1/workflows`

**Payload:**
```json
{
  "name": "Social Media Post Workflow",
  "version": 1,
  "definition": {
    "steps": [
      {
        "name": "publish_post",
        "type": "activity",
        "properties": {
          "channel": "twitter"
        }
      }
    ]
  }
}
```

### Executions

#### Start a Workflow Execution
`POST /api/v1/executions/start`

**Payload:**
```json
{
  "workflow_id": "<workflow-uuid>",
  "input": {
    "content": "Hello World!"
  }
}
```

## Configuration

Environment variables can be set in the `.env` file or directly in the `docker-compose.yml`:

| Variable | Description | Default |
|----------|-------------|---------|
| `PORT` | API server port | `8000` |
| `DATABASE_URL` | PostgreSQL connection string | `postgres://user:password@db:5432/relay_db` |
| `WORKER_ID` | Identifier for the worker instance | `worker-1` |

## Development

If you want to run the project locally without Docker:

1. Install Go 1.23+
2. Set up a PostgreSQL database.
3. Configure your `.env` file based on `.env.example`.
4. Run the API:
   ```bash
   go run cmd/api/main.go
   ```
5. Run the Worker:
   ```bash
   go run cmd/worker/main.go
   ```
