FROM golang:1.23-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY *.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -o server

FROM scratch

LABEL org.opencontainers.image.source="https://github.com/nicolaiort/kusatoko"
LABEL org.opencontainers.image.description="A simple testing webserver written in go"
LABEL org.opencontainers.image.licenses=MIT

COPY --from=builder /app/server /server
ADD https://curl.haxx.se/ca/cacert.pem /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT [ "/server" ]