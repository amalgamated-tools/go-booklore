# Development Docker Setup

This directory includes a Docker Compose configuration for running the Booklore API server locally for development and testing.

## Quick Start

### Start the Server

```bash
make docker-up
```

This will:
1. Start a PostgreSQL database container
2. Start the Booklore API server
3. Pause briefly to allow services to start up

The Booklore API will be available at `http://localhost:8080`

### Run Tests Against Docker Server

```bash
make docker-test
```

This runs all tests against the local Docker-based Booklore server using credentials from `.env.docker`.

### View Logs

```bash
make docker-logs
```

### Stop Services

```bash
make docker-down
```

To also remove volumes (database data):

```bash
make docker-clean
```

## Manual Docker Compose Commands

If you prefer direct docker-compose commands:

```bash
# Start services in background
docker-compose up -d

# Stop services
docker-compose down

# View logs
docker-compose logs -f

# View specific service logs
docker-compose logs -f booklore

# Run a command in the running container
docker-compose exec booklore curl http://localhost:8080/api/v1/healthcheck
```

## Configuration

### Environment Variables

The Docker setup uses environment variables from `.env.docker`:

```env
BOOKLORE_SERVER="http://localhost:8080"
BOOKLORE_USERNAME="demo"
BOOKLORE_PASSWORD="demo"
```

To use different credentials, update `.env.docker` or set environment variables before running tests:

```bash
BOOKLORE_SERVER=http://localhost:8080 \
BOOKLORE_USERNAME=myuser \
BOOKLORE_PASSWORD=mypass \
go test -v ./pkg/booklore/...
```

### Booklore Server Configuration

The Booklore server is configured via environment variables in `docker-compose.yml`:

- `SPRING_DATASOURCE_*` - Database connection details
- `SERVER_PORT` - API server port (default: 8080)
- `JWT_SECRET` - JWT signing key for authentication
- `LOGGING_LEVEL_*` - Log levels

Adjust these in `docker-compose.yml` as needed for your development environment.

## Database

The PostgreSQL database:
- Runs on `localhost:5432`
- Database: `booklore`
- Username: `booklore`
- Password: `booklore`

Data is persisted in a Docker volume (`postgres_data`) and survives container restarts.

To reset the database:

```bash
make docker-clean
make docker-up
```

## Troubleshooting

### Server not responding on startup

The Booklore server may take 30+ seconds to start on first run. Check logs:

```bash
make docker-logs
```

### Port already in use

If port 8080 or 5432 is already in use, edit `docker-compose.yml` and change the port mappings:

```yaml
ports:
  - "8081:8080"  # Use 8081 instead of 8080
```

### Database connection errors

Ensure PostgreSQL is healthy before the API starts. The `depends_on` condition should handle this, but if you see errors:

```bash
docker-compose restart booklore
```

### Image not found

If the Booklore image isn't available:

```bash
# Pull the latest image
docker pull ghcr.io/booklore-app/booklore:latest

# Rebuild
docker-compose down
docker-compose up -d
```

## Integration with Go Tests

The generated tests in `pkg/booklore/client_test.go` read the `BOOKLORE_SERVER` environment variable. To run them against the Docker server:

```bash
# Load .env.docker and run tests
env $(cat .env.docker | grep -v '^#' | xargs) go test -v ./pkg/booklore/...
```

Or use the make target:

```bash
make docker-test
```

## Production Deployment

**Important:** The credentials and configuration in this docker-compose setup are for local development only. Do NOT use in production.

For production deployments:
- Use strong JWT secrets
- Enable HTTPS/TLS
- Use external database services
- Configure proper logging and monitoring
- Use secrets management tools
- Follow the [Booklore deployment guide](https://github.com/booklore-app/booklore)

## Further Reading

- [Booklore Repository](https://github.com/booklore-app/booklore)
- [Docker Compose Documentation](https://docs.docker.com/compose/)
- [PostgreSQL Docker Image](https://hub.docker.com/_/postgres)
