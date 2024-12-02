FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

USER vscode

# Install golangci-lint
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

# Install air
RUN go install github.com/air-verse/air@latest