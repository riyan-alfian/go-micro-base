# go-micro-base

sample service using golang

## Folder Project Structures

```js
.
├── cmd/                     # Entrypoints for the application
│   └── server/              # Main service entrypoint
│       └── main.go          # Starts the HTTP server and loads dependencies
│
├── internal/                # Core application logic (private Go packages)
│   ├── db/                  # Database-related files
│   │   ├── migration/       # Database migration scripts
│   │   ├── seed/            # Seeders for initial data
│   │   ├── migrate.go       # Migration runner (go run)
│   │   └── seed.go          # Seeder runner (go run)
│   │
│   ├── entity/              # Domain entities and models
│   │   ├── mongo/           # MongoDB entity definitions
│   │   └── postgres/        # PostgreSQL entity definitions
│   │
│   ├── modules/             # Business feature modules (bounded contexts)
│   │   ├── auth/            # Authentication module
│   │   │   ├── controller/  # HTTP handlers / delivery layer
│   │   │   ├── dto/         # Data Transfer Objects (request/response)
│   │   │   ├── repository/  # Repository interfaces and DB adapters
│   │   │   ├── route/       # Router registration
│   │   │   └── usecase/     # Business logic (application layer)
│   │   └── user/            # Example user module
│   │       ├── controller/
│   │       ├── dto/
│   │       ├── repository/
│   │       ├── route/
│   │       └── usecase/
│   │
│   └── util/                # Shared utilities and infrastructure helpers
│       ├── broker/          # Message broker (NATS, RabbitMQ, etc.)
│       └── bucket/          # Object storage (S3, GCS, MinIO)
│
├── infra/                   # Deployment and environment infrastructure
│   ├── docker/              # Dockerfiles, docker-compose setup
│   ├── jenkins/             # Jenkins CI/CD pipeline configs
│   └── k8s/                 # Kubernetes manifests / Helm charts
│
├── environment/             # Environment configuration files (.env/.yaml)
│   ├── development.yaml
│   ├── staging.yaml
│   └── production.yaml
│
├── docs/                    # Documentation and architecture diagrams
│   └── flowchart.puml
│
├── test/                    # Unit and integration tests
│   └── sample_test.go
│
├── Makefile                 # Build/run/test/migrate commands
├── go.mod
├── go.sum
└── README.md
```
