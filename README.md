# API assessment
## Description
This is a simple API that allows you to manage films.

## Decisions taken

- First, I decided to develop the API endpoints, then the JWT authentication and finally _dockerize_ the API.
- I will use Gorilla Mux for the API endpoints. 
  - First, because it widely used in the Go community (documentation, examples, etc.)
  - Second, because it is the one I am most familiar with.
- I will use GORM for the database management, because it is the one I am most familiar with.
- I will use the `github.com/golang-jwt/jwt/v4` package for the JWT authentication, as it seems to be the most used one.
- I will use the `github.com/go-playground/validator` package for validations, as it seems to be widely used in the
  community.
