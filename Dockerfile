
FROM golang:1.23-alpine

WORKDIR /app

COPY app/go.mod app/go.sum ./

RUN go mod download

COPY app/ .

EXPOSE 8080

CMD ["go", "run", "main.go"]