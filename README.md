# go-onion

Onion Architecture Sample with Golang

## libs
- echo

## Manual Test

### Create User
```
curl -XPOST localhost:1323/users \
  -H 'Content-Type: application/json' \
  -d '{"id":"550e8400-e29b-41d4-a716-446655440000", "name":"hogename"}'
```

## Usage

```
make
```
- go test
- go run

