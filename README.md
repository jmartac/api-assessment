# API assessment

## Description

This is a simple API that allows you to manage films.

## Decisions taken

- First, I decided to develop the API endpoints, then the JWT authentication and finally _dockerize_ the API.
- I will use [Gorilla Mux](https://github.com/gorilla/mux) for the API endpoints.
    - First, because it widely used in the Go community (documentation, examples, etc.)
    - Second, because it is the one I am most familiar with.
- I will use [GORM](https://github.com/go-gorm/gorm) for the database management, because it is the one I am most
  familiar with.
- I will use the `github.com/golang-jwt/jwt/v4` package for the JWT authentication, as it seems to be the most used one.
- I will use the `github.com/go-playground/validator` package for validations, as it seems to be widely used in the
  community.

## Instructions

First, clone the repository.

Now, at the project root:

1. Run `docker build -t films-api .` to build the API image.
2. Run `docker-compose up -d` to start both the API and the database. It will start the API on port
   **8080** and the database on port **3360**, both exposed to the host, and it may take a few seconds for the database
   to be ready.
3. (Optional) Run `mysql --protocol=TCP -u film -pfilm -D films < seed.sql` to seed the database with some data.

## Endpoints

I have added a Postman collection with the endpoints and some examples.

You can log in with an existing user (endpoint `/login existing`) or create a new one to log in with it, and then use
the JWT token to access the other endpoints (only the `/login` and `/register` endpoints are public). The JWT token must
be sent in the `Authorization` header, with the `Bearer` prefix, so I recommend setting it in the collection
Authorization tab.

A collection variable `URL` is used on the endpoints, it should be set to `localhost:PORT` (in this
case, `localhost:8080`).