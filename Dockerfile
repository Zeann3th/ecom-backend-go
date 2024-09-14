# Alpine-based Go image
FROM golang:1.22.7-alpine3.20

# Install make
Run apk add --no-cache make

WORKDIR /app

# copy all project files, directories
COPY ..

# Build
Run make build

# Start App
Run make run


