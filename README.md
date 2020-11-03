# Go gin quickstart

# How it works
```go
go run app.go
```

# Getting started

## Building
```
go build
./go-gin-quickstart
```

## Verify
```
# register an user
curl -v -X POST \
  http://localhost:8080/api/users/ \
  -H "Accept: application/json" \
  -H 'Content-Type: application/json' \
  -d '{
    "user": {
        "username": "User",
        "password": "12345678",
        "email": "user@gmail.com",
        "bio": "an user is for good",
        "image": "https://pic.url"
    }
}'

# login
curl -X POST \
  http://localhost:8080/api/users/login \
  -H "Accept: application/json" \
  -H 'Content-Type: application/json' \
  -d '{
    "user": {
        "email": "user@gmail.com",
        "password": "12345678"
    }
}'

# follow user
curl -X POST \
  -H "Accept: application/json" \
  -H 'Content-Type: application/json' \
  'http://localhost:8080/api/profiles/User2/follow?access_token=${ACCESS_TOKEN}'


# unfollow user
curl -X DELETE \
  -H "Accept: application/json" \
  -H 'Content-Type: application/json' \
  'http://localhost:8080/api/profiles/User1/follow?access_token=${ACCESS_TOKEN}' \
```

## Testing
```
➜  go test -v ./... -cover
```