# Web Service

The `web` service is the frontend client for FitFeed, built with React and Ant Design.

## Architecture

The frontend follows a modern, component-based architecture:

- **Components:** Reusable UI components for layout, features, and feedback.
- **Context:** Global state management for authentication and configuration.
- **Services:** API integration for OAuth, Passkeys, and data fetching.
- **Pages:** Main views of the application (Home, Profile, etc.).

## Key Components

- **AppHeader:** Sticky header with navigation, user profile dropdown, and login/register modal.
- **AppFooter:** Simple footer with copyright and link to FitFeed.
- **MainLayout:** Layout wrapper providing consistent spacing and structure across all pages.
- **AuthContext:** Manages authentication state, user session, and application configuration.

## Development

The frontend uses Vite for fast development and hot-reloading. All dependencies are managed using Bun.

### Commands

- `bun run dev` - Start the Vite development server.
- `bun run build` - Build the production-ready client.
- `bun run lint` - Run ESLint to check for code quality.

## Theme & Styling

FitFeed uses Ant Design's modern styling and provides a consistent theme throughout the application. The primary color is a vibrant orange, inspired by fitness social platforms.
