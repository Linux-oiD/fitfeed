# API Service

The `api` service provides the main interface for the FitFeed web client, offering profile management and configuration.

## Features

- **Profile Management:** Users can view and update their personal information (name, avatar, email).
- **Service Configuration:** Provides the frontend with necessary backend URLs and settings.
- **Route Protection:** All profile-related routes are protected by JWT middleware.

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
