# How to Run:

## Server:

You can run with your own host and port.
Defaults are:
  host: localhost
  port: 8080

```bash
HTTP_HOST={host} HTTP_PORT={port} go run cmd/main.go
```

## Tests:

```bash
go test -v -p 1 ./...
```
If it fails, try to delete db, then run it again:
```bash
rm internal/handlers/cities.db
```
Don't forget to delete (above) after tests!