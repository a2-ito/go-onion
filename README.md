# go-onion

Onion Architecture Sample with Golang

## Todo

- [ ] "Hello world" app with echo
- [ ] env configuration
- [ ] Build as a Container via buildpacks
- [ ] format with go fmt
- [ ] Build with a dockerfile
- [ ] lint for dockerfile
- [ ] Logging - zap
- [ ] disable linter structcheck due to warninng
- [ ] lint with golangci-lint
- [ ] go test

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

