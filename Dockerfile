#
# Builder
#

# Use the official Golang image as the base image
FROM golang:latest as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download && go mod verify

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN make build

#
# Runner
#

FROM busybox

WORKDIR /

# Expose port 8080 to the outside world
EXPOSE 8080

# Copy from builder the final binary
COPY --from=builder /app/bin/nhlscores/nhlscores .

ENTRYPOINT [ "./nhlscores" ]