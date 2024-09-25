# Kusatoko

![GitHub License](https://img.shields.io/github/license/nicolaiort/kusatoko?style=for-the-badge) ![GitHub top language](https://img.shields.io/github/languages/top/nicolaiort/kusatoko?style=for-the-badge) ![GitHub commit activity](https://img.shields.io/github/commit-activity/m/nicolaiort/kusatoko?style=for-the-badge)

A simple go server for testing connection stuff.

## Run

### With docker

```bash
docker pull ghcr.io/nicolaiort/kusatoko
docker run -p 8080:8080 ghcr.io/nicolaiort/kusatoko
```

### With go for dev 🛠️

```bash
go run main.go
```

## Endpoints

| Endpoint         | Description                                                               |
| ---------------- | ------------------------------------------------------------------------- |
| `/`              | Returns a simple text message                                             |
| `/healthz`       | Returns `ok`                                                              |
| `/whatsmyip`     | Returns the requesting client's IP                                        |
| `/headers`       | Returns all request headers                                               |
| `/status/{code}` | Returns the status code provided via `{code}` and it's associated message |

## Config

| ENV-Var        | Default         | Description                 |
| -------------- | --------------- | --------------------------- |
| `PORT`         | 8080            | The server's listening port |
| `ROOT_MESSAGE` | "Hello, World!" | The text displayed on `/`   |
