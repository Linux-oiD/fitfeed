# API Service

The `api` service provides the main interface for the FitFeed web client, offering profile management and configuration.

## Project Structure

The service follows a flat module structure:
- `cmd/api/main.go`: Application entry point.
- `internal/config/`: Configuration loading logic.
- `internal/controller/http/`: HTTP handlers and routing.
- `internal/entity/`: Domain models.
- `internal/repo/`: Data access layer.
- `internal/usecase/`: Business logic.
- `pkg/`: Shared utility packages.

## API Endpoints

- **GET /v1/config:** Returns the current configuration (auth_url, api_url).
- **GET /v1/users/{username}:** Retrieves a user's profile information.
- **PUT /v1/users/profile:** Updates the currently logged-in user's profile (requires JWT).

## Internal Structure

The `api` service also follows a layered architecture:

- **Controller:** HTTP handlers and routing.
- **UseCase:** Business logic for user and profile management.
- **Repo:** Data access for `User` and `Profile` tables.
- **JWTMiddleware:** Middleware for validating tokens and extracting user context.
