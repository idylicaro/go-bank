FROM qmcgaw/godevcontainer AS devcontainer

FROM golang:1.19 as base
WORKDIR /go-ms-upload

# Fetch dependencies
COPY go.mod go.sum ./
RUN go mod download

# Build
# CGO_ENABLED=0 because the alpine version
COPY . ./
RUN CGO_ENABLED=0 go build -v ./...

FROM base AS test

# Test
CMD go test -v ./...

