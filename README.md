# go-social-feed-api

A Go-based social feed API backend.

## Folder Structure

- **`/bin`** – Compiled binaries.
- **`/cmd`** – Application entry points and main executables.
- **`/api`** – HTTP handling, transport layer, middlewares, and servers.
- **`/migrate`** – SQL migration configurations and scripts.
- **`/internal`** – Database interaction and data validations.
- **`/docs`** – Swagger documentation.
- **`/scripts`** – Deployment and server setup scripts.
- **`/web`** – Frontend assets or static applications.

## Development Log & Commits

### 1. HTTP Server Setup

- [Create an HTTP server with http.ServeMux](https://github.com/Prakash-Ravichandran/go-social-feed-api/commit/768f0417b2b2a991acfef872e1eedee57a3afb09)
- [Migrate from standard mux to chi router](https://github.com/Prakash-Ravichandran/go-social-feed-api/commit/8dd9b26295bea17b627f4cce0d3954fca539c382)

### 2. Hot Reloading with Air

- [Configure hot reloading with Air](https://github.com/Prakash-Ravichandran/go-social-feed-api/commit/fa79ea8f2f8ac4105a5b18c3275da970cbecb838)

Start the server locally using:

```bash
air -v
air
```

### 3. Environment Configuration

- [Setup direnv on Windows using Git Bash](https://gist.github.com/Prakash-Ravichandran/1cd0ea17671702c6a4b71eb0b4fdfe06)

### 4. Repository Pattern & Database

- [Establish a DB Connection Pool](https://github.com/Prakash-Ravichandran/go-social-feed-api/commit/d81a30a8b76a88c6d23633284dbdf366af30d6da)
- [Implement the Repository Pattern Skeleton](https://github.com/Prakash-Ravichandran/go-social-feed-api/commit/7f2b7123484a9bcb3211bf51d5ef299abd663d61)
- [Add Post Model to Repository](https://github.com/Prakash-Ravichandran/go-social-feed-api/commit/6b731760b10c25c476e7b4c0833fdce2757c37ec)
- [Add User Model to Repository](https://github.com/Prakash-Ravichandran/go-social-feed-api/commit/72fc9283bafa4f232079363822b3ee488ad18312)

---

## SQL Migrations

### Raw CLI Usage

**Create a migration:**

```bash
migrate create -seq -ext sql -dir ./cmd/migrate/migration create_users

```

**Run up migrations:**

```bash
migrate -path=./cmd/migrate/migrations -database="postgres://admin:adminpassword@localhost/socialfeed?sslmode=disable" up

```

### Makefile Shortcuts

**Step 1: Create a new SQL migration file**

```bash
make migration alter_post_table

```

**Step 2: Apply pending migrations (Up)**

```bash
make migrate-up

```

**Step 3: Roll back migrations (Down)**

```bash
make migrate-down

```
