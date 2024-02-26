# Mekdi

A small web app with HTML views, made with Go and PostgreSQL, and a little help
from HTMX for the frontend.

## Development

To develop locally you need to [install Go](https://go.dev/dl) and [PostgreSQL](https://www.postgresql.org/download/)
on your machine.

After getting all dependencies, you can run the development server via the
commands below (Make sure PostgreSQL is also running (`sudo service postgresql start`
for Linux)):

```bash
go mod tidy
go run .
```

The dev server will start in `localhost:8080`.
