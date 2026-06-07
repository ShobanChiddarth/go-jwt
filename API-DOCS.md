# API DOCS

## POST `/signup`

body:
```json
{
    "email":"sample@example.com",
    "password":"PASSWORD"
}
```

## POST `/login`

body:
```json
{
    "email":"sample@example.com",
    "password":"PASSWORD"
}
```
Cookie will be set

## GET `/validate`

GET it with the cookie set by `/login`

