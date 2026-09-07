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

### GET: Implement POST GetById

[Implement POST GetById](https://github.com/Prakash-Ravichandran/go-social-feed-api/commit/2a5c8962a0687a731443cbfa627cba31d861b65e#diff-96ee07431493095d4c7e7ed42090aad5945b7ef6ee635ce6453890f0108337cb)

Test it: http://127.0.0.1:3000/posts/4

Response

```JSON
{
    "id": 4,
    "content": "Mission Impossible, Fast and Furious",
    "title": "Best Movies in Hollywood",
    "user_id": "1",
    "tags": [
        "Hollywood"
    ],
    "created_at": "2026-08-08T13:31:08Z",
    "updated_at": "2026-08-08T13:31:08Z"
}
```

### PUT: Implement PUT: updatePostById

[Implement PUT updatePostById](https://github.com/Prakash-Ravichandran/go-social-feed-api/commit/4f9c8bb9636ac5302df7936958d7f7a807256f3d)

Test it:PUT: http://127.0.0.1:3000/posts/2

Input

```JSON
{

  "title": "Best Songs of 2000",
  "content": "Best songs of 2000 in Tamil, Telugu, Hindi",
  "tags": ["Kollywod"]
}

```

Response

```JSON
{
    "id": 2,
    "content": "Best songs of 2000 in Tamil, Telugu, Hindi",
    "title": "Best Songs of 2000",
    "user_id": "1",
    "tags": [
        "Kollywod"
    ],
    "created_at": "2026-08-08T07:38:58Z",
    "updated_at": "2026-08-08T07:38:58Z"
}
```

### HTTP Payload Validation

[HTTP Payload Validation](https://github.com/Prakash-Ravichandran/go-social-feed-api/commit/d98584f5ce74e07a152ece4bd2422fd1e246d5aa)

### Attach Comments[] to Post

[attaching comments to post](https://github.com/Prakash-Ravichandran/go-social-feed-api/commit/f449e1eacf4a935a4c1072f9af38590d2eea855d)

Query 1: Query 1 is preferred over 2

```sql
SELECT * FROM COMMENTS c
JOIN users ON users.id = c.user_id WHERE c.post_id = 2;
```

Individual columns for query 1

```sql
SELECT c.id, c.post_id, c.user_id, c.content, c.created_at,  users.id, users.username FROM COMMENTS c
JOIN users ON users.id = c.user_id WHERE c.post_id = 2 ORDER BY c.created_at DESC;
```

Query 2:

```sql
SELECT * FROM posts
JOIN comments ON comments.post_id = posts.id;
```

### Updating Title, content individually

- because updating one field is a valid thing
- default value of string variable is a empty "" string and a empty string is valid value in updating a post

The HTTP validation here cannot have both the fields to be required while updating the post. Hence it was removed.

- user only updates the title, check for `tempUpdatePost.Title != nil` is not nil and update to post.
- user only updates the content, check for `tempUpdatePost.Content != nil` is not nil and update to post.

**By adding if tempUpdatePost.Content != nil:**

- Go checks: "Did the user send a new content field?"

- It sees `tempUpdatePost.Content` is nil (No).

- It skips the update and keeps the old content untouched.

<img width="1160" height="801" alt="Image" src="https://github.com/user-attachments/assets/df46eaae-d352-4d2d-a95f-972874e46ac1" />

### Request Validation: `required` vs `omitempty`

| Feature           | `required`               | `omitempty`                        |
| :---------------- | :----------------------- | :--------------------------------- |
| **Primary Use**   | `POST` (Create)          | `PATCH` (Partial Update)           |
| **Missing Field** | ❌ **Fails** (Mandatory) | ✅ **Passes** (Validation skipped) |
| **Present Field** | Enforces rules           | Enforces rules                     |
| **Go Field Type** | Value types (`string`)   | Pointer types (`*string`)          |

#### Quick Rules

- **`required`**: The JSON field **must** be present in the request. Use for creating resources where data is mandatory.
- **`omitempty`**: The JSON field is **optional**. If omitted (`nil`), validation is skipped. If provided, constraints (e.g., `max=100`) are enforced.

> **Rule of Thumb:** Use `required` for `POST` payloads and `omitempty` with pointer types (`*string`) for `PATCH` partial updates.
