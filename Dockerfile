
FROM golang:1.23-alpine

WORKDIR /app

RUN go mod download

COPY app/go.mod app/go.sum ./

COPY app/ .

EXPOSE 8080

CMD ["go", "run", "main.go"]