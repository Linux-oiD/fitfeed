# Configuration Reference

FitFeed is configured using a unified `config.toml` file.

## Core Configuration

| Key | Description | Default |
| :--- | :--- | :--- |
| `api.port` | Port for the API service | 8082 |
| `auth.port` | Port for the Auth service | 8081 |
| `auth.secret` | JWT signing secret | (generated) |
| `database.postgres.host` | DB Host | localhost |

## OAuth Providers

OAuth providers are configured under `[auth.providers.{name}]`. You must enable them and provide Client ID/Secret from the respective provider's developer console.
