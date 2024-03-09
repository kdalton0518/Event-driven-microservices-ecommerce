# Base
FROM golang:1.22-alpine3.19 as base

WORKDIR /go/src/app
COPY ./apps/payment-svc/go.mod ./
COPY ./apps/payment-svc/go.sum ./
COPY ./apps/payment-svc/ ./
RUN go mod download
RUN go build -o main ./cmd/queue

# Binary
FROM alpine:3.19 as binary
COPY --from=base /go/src/app/main .
CMD ["./main"]
