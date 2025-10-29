# Start with an official Golang image as the base image
FROM golang:1.25.3 AS base

# Install dependencies
RUN apt-get update \
    && apt-get install -y \
    libcurl4-openssl-dev \
    libxml2-dev \
    mime-support \
    tzdata \
    && rm -rf /var/lib/apt/lists/*

# Set timezone to UTC
ENV TZ=UTC
RUN ln -sf /usr/share/zoneinfo/UTC /etc/localtime \
    && echo "UTC" > /etc/timezone \
    && dpkg-reconfigure -f noninteractive tzdata \
    && date

# Set the working directory
WORKDIR /app/mp

# Copy Go module files first for dependency caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go application
RUN go build -o ./out/mp main.go

# Expose application port
EXPOSE 8001

# Run the app
CMD ["./out/mp"]