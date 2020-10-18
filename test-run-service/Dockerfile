FROM golang:alpine as builder

RUN apk update && apk upgrade && \
    apk add --no-cache git

RUN mkdir /app
WORKDIR /app

ENV GO111MODULE=on

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o test-run-service *.go

# Run container
FROM alpine:latest

RUN mkdir /app
WORKDIR /app

COPY --from=builder /app .

ENTRYPOINT ["./test-run-service"]