# Users JWT Authentication

Users JWT Authentication using Refresh Token Rotation mechanism.

## Getting started

### Clone the repo

```bash
git clone https://github.com/shaikrasheed99/golang-user-jwt-authentication.git
cd golang-user-jwt-authentication/
```

### Environment variables

For environment variables, create a `.env` file in home directory of this project.

```
DB_HOST="localhost"
DB_PORT=5432
DB_USER="postgres"
DB_PASSWORD="postgres"
DB_NAME="users"
JWT_SECRET="[jwt secret key]"
JWT_ISSUER="[issuer name]"
JWT_ACCESS_TOKEN_EXPIRATION_IN_MINUTES=10
JWT_REFRESH_TOKEN_EXPIRATION_IN_MINUTES=15
```

## Localhost server

To start the localhost server, execute the below command in the terminal.

```bash
make run
```

## API endpoints

### Signup

##### Request

```
curl --location --request POST 'http://localhost:8080/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "first_name": "Iron",
    "last_name": "Man",
    "username": "ironman123",
    "password": "ironman@123",
    "email": "ironman@gmail.com",
}'
```

##### Response

```
{
    "status": "success",
    "code": "OK",
    "message": "successfully saved user details",
    "data": null
}
```

`Access Token` and `Refresh Token` values would be returned through the `httpOnly` cookies.

### Login

##### Request

```
curl --location --request POST 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username": "ironman123",
    "password": "ironman@123"
}'
```

##### Response

```
{
    "status": "success",
    "code": "OK",
    "message": "successfully logged in",
    "data": null
}
```

`Access Token` and `Refresh Token` values would be returned through the `httpOnly` cookies.

### Logout

##### Request

User needs to provide `Access Token` in the request header to access this api.

```
curl --location --request POST 'http://localhost:8080/logout' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer [User's access Token]' \
--data '{
    "username": "ironman123"
}'
```

##### Response

```
{
    "status": "success",
    "code": "OK",
    "message": "successfully logged out",
    "data": null
}
```

Empty `Access Token` and `Refresh Token` values would be returned through the `httpOnly` cookies.

### Refresh Access Token

##### Request

User needs to provide `Refresh Token` in the request header to access this api.

```
curl --location --request POST 'http://localhost:8080/refresh' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer [User's refresh token]' \
--data '{
    "username": "ironman123"
}'
```

##### Response

```
{
    "status": "success",
    "code": "OK",
    "message": "successfully received access token",
    "data": null
}
```

`Access Token` and `Refresh Token` values would be returned through the `httpOnly` cookies.

### Fetch all users

This api is only accessed by Admins.

##### Request

Admin needs to provide `Access Token` in the request header to access this api.

```
curl --location --request GET 'http://localhost:8080/users' \
--header 'Authorization: Bearer [Admin's access token]' \
--data ''
```

##### Response

```
{
    "status": "success",
    "code": "OK",
    "message": "successfully got list of users",
    "data": [
        {
            "id": 1,
            "first_name": "Captain",
            "last_name": "America",
            "username": "captain12",
            "email": "captainamerica@gmail.com",
            "role": "user"
        },
        {
            "id": 2,
            "first_name": "Iron",
            "last_name": "Man",
            "username": "ironman123",
            "email": "ironman@gmail.com",
            "role": "admin"
        }
    ]
}
```

### Fetch users by username

This api can be accessed by Admins and particular user.

##### Request

User needs to provide `Access Token` in the request header to access this api.

```
curl --location --request GET 'http://localhost:8080/users/ironman123' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer [User's access token]' \
--data-raw '{
    "username": "ironman123",
    "password": "ironman@123"
}'
```

##### Response

```
{
    "status": "success",
    "code": "OK",
    "message": "successfully got user details",
    "data": {
        "id": 1,
        "first_name": "Iron",
        "last_name": "Man",
        "username": "ironman123",
        "email": "ironman@gmail.com",
        "role": "admin"
    }
}
```
