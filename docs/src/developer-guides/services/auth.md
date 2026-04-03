# Auth Service

The `auth` service handles registration, login, and token generation for FitFeed. It provides multiple authentication methods:

- **OAuth:** Users can register and log in using third-party providers. If an existing user logs in with a new provider, it is automatically linked to their account.
- **Passkeys (WebAuthn):** Users can register and log in using biometric authentication or security keys.
- **JWT:** Once authenticated, the service generates a JWT and returns it to the client as a cookie.

## API Endpoints

- **GET /v1/oauth/{provider}/auth:** Initiates the OAuth flow.
- **GET /v1/oauth/{provider}/callback:** Handles the OAuth callback from the provider.
- **GET /v1/passkey/register/begin:** Initiates Passkey registration.
- **POST /v1/passkey/register/finish:** Finalizes Passkey registration.
- **GET /v1/passkey/login/begin:** Initiates Passkey login.
- **POST /v1/passkey/login/finish:** Finalizes Passkey login.
- **GET /v1/oauth/{provider}/logout:** Clears the JWT cookie and logs the user out.

## Internal Structure

The `auth` service follows a layered architecture:

- **Controller:** HTTP handlers and routing.
- **UseCase:** Business logic for authentication and provider management.
- **Repo:** Data access for `User`, `Profile`, and `OauthProvider` tables.
- **JWTManager:** Logic for generating and validating tokens.
- **PasskeyManager:** Integration with WebAuthn for biometric authentication.
