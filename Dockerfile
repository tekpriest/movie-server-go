FROM golang:1.12-alpine

RUN apk add --no-cache git

WORKDIR /app/movie-server-go

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 5013

CMD ["movie-server"]
