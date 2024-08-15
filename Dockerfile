# Build stage
FROM golang:1.22.5-alpine AS builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/server ./main.go

# Final stage
FROM alpine:3.20.2

WORKDIR /go/src/app

COPY --from=builder /build/bin/server .
COPY --from=builder /build/.env .

CMD [ "./server" ]