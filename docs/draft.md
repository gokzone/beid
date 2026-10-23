# Registry

## Signup

`POST /registry/signup`

request json:
```json
{
  "username": "...",  // required, min_length=2, max_length=64
  "email": "...",     // required, format=email
  "password": "..."   // required, min_length=6, max_length=256
}
```

response json:
```json
{
  "uid": "...",   // if success
  "error": "..."  // if failed
}
```

# Auth

## Login

`POST /auth/login`

## Verify

`POST /auth/verify`

## Refresh

`POST /auth/refresh`

