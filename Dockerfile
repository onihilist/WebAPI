
FROM golang:1.23-alpine

WORKDIR /app

COPY app/go.mod app/go.sum ./

RUN go mod download

RUN apk add --no-cache nodejs npm

RUN npm install -g sass

COPY app/ .

RUN sass public/scss:public/css

EXPOSE 8080

CMD ["go", "run", "main.go"]