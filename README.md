# Mekdi

A small web app with HTML views, made with Go and PostgreSQL, and a little help
from HTMX for the frontend.

## Development

To develop locally first you need to [install Go](https://go.dev/dl) and
[PostgreSQL](https://www.postgresql.org/download/) on your machine.

First copy the .env.example and rename the new file to `.env`. Fill the file with
the appropriate data. Then you can recover the database from the dump file at the
project root (`mekdi.sql`), and make sure PostgreSQL is also running (`sudo service
postgresql start` for Linux).

After getting all dependencies, you can run the development server via the
commands below:

```bash
go mod tidy
go run .
```

The dev server will start in `localhost:8080`.
