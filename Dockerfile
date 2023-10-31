FROM golang:1.21

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code
COPY *.go ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /go_hello

# Bind port
EXPOSE 8080

# Run
CMD ["/go_hello"]