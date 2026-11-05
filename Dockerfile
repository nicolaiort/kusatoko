FROM --platform=$BUILDPLATFORM mirror.gcr.io/library/golang:1.26-alpine AS builder

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGETARCH} go build -o server

FROM scratch

LABEL org.opencontainers.image.source="https://github.com/nicolaiort/kusatoko"
LABEL org.opencontainers.image.description="A simple testing webserver written in go"
LABEL org.opencontainers.image.licenses="MIT"

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /app/server /server

ENTRYPOINT ["/server"]
