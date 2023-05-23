# Tmp base image
FROM golang:alpine AS build

# container workspace
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy source code and compile app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

### Base image###
FROM alpine:latest

# container workspace
WORKDIR /app

# Copy binary files from compilation stage
COPY config.yaml .
COPY --from=build /app/app .

# Expose app port
EXPOSE 3000

# Execute app 
CMD ["./app"]