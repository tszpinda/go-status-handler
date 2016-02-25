## To run:

```go
go run server.go
```

## To test:
```bash
go test

or:

wget -qO- --header="Content-Type: application/json" --header="x-status-code: 200" localhost:8080
wget -qO- --header="Content-Type: application/json" --header="x-status-code: 404" localhost:8080
wget -qO- --header="Content-Type: application/wrong" localhost:8080
```


