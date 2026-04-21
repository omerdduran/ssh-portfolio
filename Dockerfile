# syntax=docker/dockerfile:1.7

FROM golang:1.26-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/ssh-portfolio .

FROM alpine:3.21 AS runner
WORKDIR /app
RUN apk add --no-cache tini ca-certificates && \
    addgroup -S app && adduser -S app -G app && \
    mkdir -p /app/.ssh && chown -R app:app /app
ENV PORTFOLIO_URL=https://www.omerduran.dev
COPY --from=builder /out/ssh-portfolio /app/ssh-portfolio
USER app
EXPOSE 23234
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["/app/ssh-portfolio"]
