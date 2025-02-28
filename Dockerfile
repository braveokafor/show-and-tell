# checkov:skip=CKV_DOCKER_7: "Ensure the base image uses a non latest version tag"
# checkov:skip=CKV_DOCKER_2: "Ensure that HEALTHCHECK instructions have been added to container images"
FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o show-and-tell ./cmd/server


FROM cgr.dev/chainguard/static:latest

WORKDIR /app

COPY --from=builder /app/show-and-tell .
COPY --from=builder /app/ui ./ui

USER nobody
EXPOSE 8080

ENTRYPOINT ["./show-and-tell"]