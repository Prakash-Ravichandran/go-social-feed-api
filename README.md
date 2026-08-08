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

### 5. Posts CRUD

**JSON Marshalling Responses**

[JSON Marshalling Responses](https://github.com/Prakash-Ravichandran/go-social-feed-api/commit/1e3c41c605a6b4a7e4a36ebc20caacb8faec37ae)

#### Create a post by having 1 user created in database

- [create a post](https://github.com/Prakash-Ravichandran/go-social-feed-api/commit/cbfc71b003fefc9007eb44e527e9ad742438a9a6)

create a user, where id, created and updated are automatically created.

```SQL
INSERT INTO
  users (email, username, password)
VALUES
  ('mail@email.com', 'Tom', 'pwpw')
```

<img width="1667" height="622" alt="Image" src="https://github.com/user-attachments/assets/33ce9df5-ea71-4d16-a78d-89b0bff4dd17" />

create post in which tags column are created using alter table syntax.

- create post using below POST parameters

```JSON
{

  "title": "Best Movies in Hollywood",
  "content": "Mission Impossible, Fast and Furious",
  "tags": ["Hollywood"]
}
```

Response:

```JSON
[
  {
    "id": 4,
    "title": "Best Movies in Hollywood",
    "user_id": 1,
    "content": "Mission Impossible, Fast and Furious",
    "tags": "{Hollywood}",
    "created_at": "2026-08-08 13:31:08+00",
    "updated_at": "2026-08-08 13:31:08+00"
  }
]
```
