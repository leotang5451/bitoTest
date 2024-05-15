FROM golang:1.22-alpine as builder

WORKDIR /app

COPY . .

RUN go build -o bitotest .

FROM alpine:latest

COPY --from=builder /app/bitotest /usr/local/bin/bitotest

EXPOSE 8080

CMD ["bitotest"]