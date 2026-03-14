FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -o ssh-portfolio .

FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/ssh-portfolio .
EXPOSE 23234
CMD ["./ssh-portfolio"]
