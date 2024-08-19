# Google OAuth Go Server

This project is a Go server that implements Google OAuth 2.0 authentication, with user data stored in a PostgreSQL database. The server uses JWT (JSON Web Token) for session management, and all authenticated routes are protected using middleware.

## Features

- Google OAuth 2.0 Authentication
- JWT-based session management
- PostgreSQL for user data storage
- Database migrations with `golang-migrate`
- Protected routes using middleware
- Error handling with custom utilities

## Project Structure

.
├── cmd
│ └── server
│ └── main.go
├── db
│ └── migrations
│ ├── 000001_create_users_table.down.sql
│ └── 000001_create_users_table.up.sql
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│ ├── auth
│ │ ├── auth.go
│ │ └── middleware.go
│ ├── config
│ │ └── config.go
│ ├── google
│ │ └── google.go
│ ├── handlers
│ │ ├── auth_handler.go
│ │ ├── health_handler.go
│ │ └── user_handler.go
│ └── models
│ └── user.go
└── pkg
├── db
│ └── db.go
└── utils
└── utils.go

## Requirements

- Go 1.17 or higher
- Docker and Docker Compose
- PostgreSQL

## Setup

1. **Clone the repository:**

   ```sh
   git clone https://github.com/bsrisompong/google-oauth-go-server.git
   cd google-oauth-go-server
   ```

2. **Create a `.env` file:**

   Create a `.env` file in the project root with the following contents:

   CLIENT_ID=your-google-client-id
   CLIENT_SECRET=your-google-client-secret
   REDIRECT_URL=http://localhost:8080/auth/google/callback
   JWT_SECRET_KEY=your-secret-key
   DATABASE_URL=postgres://postgres:postgres@postgres:5432/database?sslmode=disable

3. **Run the application:**

   You can use Docker Compose to build and run the application:

   ```sh
   docker-compose up --build
   ```

   This will start the PostgreSQL database and the Go server.

4. **Access the server:**

   The server will be accessible at `http://localhost:8080`.

## Database Migrations

This project uses `golang-migrate` to handle database migrations. Migration files are located in the `db/migrations` directory.

To apply migrations:

```sh
docker-compose exec go-server ./main migrate up
```

To roll back the last migration:

```sh
docker-compose exec go-server ./main migrate down
```

## Routes

### Public Routes

- `POST /auth/google`: Handles Google OAuth authentication and returns a JWT.

### Protected Routes

- `GET /auth/me`: Returns the authenticated user's information. Requires a valid JWT.

## Issues

### Automatic Database Migration Issue

Currently, database migrations are applied automatically when the server starts. However, there are scenarios where this process might fail silently or not apply the expected migrations due to configuration or path issues.

**Possible Enhancements:**

- **Verbose Logging:** Enhance logging during the migration process to provide more visibility into what is being applied.
- **Manual Trigger Option:** Consider adding an option to manually trigger migrations as part of the deployment or startup process.
