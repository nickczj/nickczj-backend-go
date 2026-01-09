# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Build and Run Commands

```bash
# Run the application (module mode)
go run .

# Build
go build .

# Generate (tidy modules, download dependencies)
go generate
```

## Environment Configuration

The app uses Viper for configuration with environment-based config files:
- Set `APP_ENV` environment variable to select config (`local`, `prod`)
- Config files are in `config/` directory (e.g., `config/local.yml`, `config/prod.yml`)
- Defaults to `local` when `APP_ENV` is not set

Required environment variables for production:
- `APP_ENV` - Environment name
- `GCP_PROJECT_ID` - Google Cloud project ID
- `GCP_BUCKET_NAME` - Google Cloud bucket name

## Architecture

**HTTP Framework**: Gin-gonic with middleware chain (CORS → JWT auth → routes)

**Global State** (`global/global.go`): Shared instances for Database (gorm), Cache (go-redis/cache), and HTTP Client (resty)

**Initialization Flow** (`main.go`):
1. Load environment config via Viper
2. Initialize database (CockroachDB/PostgreSQL via GORM)
3. Initialize cache (Redis in prod, TinyLFU locally)
4. Register routes
5. Start server on :8080

**Key Packages**:
- `auth/` - Auth0 JWT middleware
- `cache/` - Redis caching with generic `GetElse` pattern for cache-aside
- `config/` - Environment config and GCP Secret Manager integration
- `database/` - GORM setup with PostgreSQL/MySQL dialects
- `service/` - Business logic handlers (weather, flights, finances)
- `model/` - Data models and request structs

**API Routes** (see `routes.go`):
- `/p/v1/nw/*` - Private finance endpoints (JWT protected)
- `/weather/now` - Public weather data from data.gov.sg
- `/flights/*` - Flight search endpoints
- `/stats` - Gin-stats metrics endpoint
