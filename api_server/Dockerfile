# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH environment variable at /go.
FROM golang:1.18 as builder

# Create and change to the app directory.
WORKDIR /app

# Retrieve application dependencies using go mod.
# This allows the container build to be cached unless the go.mod or go.sum files change.
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy local code to the container image.
COPY . ./

# Build the binary.
# -o myapp specifies the name of the compiled application.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o myapp

# Use the official lightweight Alpine image for a lean production container.
# https://hub.docker.com/_/alpine
FROM alpine:3.14
WORKDIR /app

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/myapp /app/myapp

# Create a volume for persistent data storage.
VOLUME /app/data

# Run the web service on container startup.
CMD ["/app/myapp"]
