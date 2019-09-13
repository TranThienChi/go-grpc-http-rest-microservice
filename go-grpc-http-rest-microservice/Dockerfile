# Start from the latest golang base image
FROM golang:latest as builder

# Add maintainer info
LABEL maintainer="Tran Thien Chi <thienchi.tr@outlook.com.vn>"

# Set the current working directory inside the container
WORKDIR /app

# Set the environment variable
ENV GO111MODULE=on

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the working directory inside the container
COPY . /app

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/app cmd/server/main.go


# Start a new stage from scratch #
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /app

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/bin/app ./

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./app"] 