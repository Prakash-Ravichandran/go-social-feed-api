# go-social-feed-api

### Folder Structure

- /bin - going to keep our code into compiled binaries.
- /cmd - the main executables or the entry points of the application.
- /api - anything related to HTTP, Transport layer, Middlewares, servers.
- /migrate - related to sql migrations.
- /internal - Interacting with databases, data validations.
- /docs - swagger doc for the application.
- /scripts - Scripts for setting up the servers.
- /web - Frontend of the application will go here/ any static apps will go here.

### Create an HTTP server

- [create an HTTP server with http.ServeMux](https://github.com/Prakash-Ravichandran/go-social-feed-api/commit/768f0417b2b2a991acfef872e1eedee57a3afb09)

- [mux -> chi](https://github.com/Prakash-Ravichandran/go-social-feed-api/commit/8dd9b26295bea17b627f4cce0d3954fca539c382)

### Hot reloading with air

[hot reloading with air](https://github.com/Prakash-Ravichandran/go-social-feed-api/commit/fa79ea8f2f8ac4105a5b18c3275da970cbecb838)

- Start the server using

```bash
    air -v
    air
```

### Setup direnv

[direnv: windows using gitbash](https://gist.github.com/Prakash-Ravichandran/1cd0ea17671702c6a4b71eb0b4fdfe06)

### Repository Pattern

[Repository Pattern: Implement a Skeleton](https://github.com/Prakash-Ravichandran/go-social-feed-api/commit/7f2b7123484a9bcb3211bf51d5ef299abd663d61)

[Repository Pattern: Post Model](https://github.com/Prakash-Ravichandran/go-social-feed-api/commit/6b731760b10c25c476e7b4c0833fdce2757c37ec)

[Repository Pattern: User Model](https://github.com/Prakash-Ravichandran/go-social-feed-api/commit/72fc9283bafa4f232079363822b3ee488ad18312)

### Establish a DB Connection Pool

[Establish a DB Connection Pool](https://github.com/Prakash-Ravichandran/go-social-feed-api/commit/d81a30a8b76a88c6d23633284dbdf366af30d6da)

### SQL Migrations

cmd:

```bash
migrate create -seq -ext sql -dir ./cmd/migrate/migration create_users
```

create a table using migrate

```bash
migrate -path=./cmd/migrate/migrations -database="postgres://admin:adminpassword@localhost/socialfeed?sslmode=disable" up
```

Result: A table skeleton will be created.

- Run migration with help of Makefile.

Step 1: Execute for creating sql file: Ex.alter_post_table argument is assigned to MAKECMDGOALS

```bash
make migration alter_post_table
```

```makefile
.PHONY: migrate-create
migration:
	@migrate create -seq -ext sql -dir $(MIGRATIONS_PATH) $(filter-out $@,$(MAKECMDGOALS))
```

Step 2: up the migration to execute the query in .sql

```bash
make migrate-up
```

Step 3: down the migration to execute the query in .sql

```bash
make migrate-down
```
