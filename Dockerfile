FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/server

RUN go build -o films-api

EXPOSE 8080

CMD ["./films-api", "-p=8080"]