ARG GO_VERSION=1.24.0

######################################################### Base
FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-alpine AS base
RUN apk add --no-cache git tree dumb-init ca-certificates

######################################################### Build
FROM base as build
WORKDIR /app

## swagger
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .
RUN swag init && \
    go mod download -x && \
    CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o main .

######################################################### Runner
FROM scratch AS runner

WORKDIR /app

## Fiber PreFork
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /usr/bin/dumb-init /usr/bin/dumb-init
COPY --from=build /app/main ./

EXPOSE 3000

ENTRYPOINT ["/usr/bin/dumb-init", "--"]
CMD ["./main"]
