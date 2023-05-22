# Base image
FROM golang:latest

# workspace container
WORKDIR /go/src/app

# Download Go modules
COPY . .

# Download project's dependecies
RUN go mod download



# Compila la aplicación Go
RUN go build -o main .

# Expone el puerto en el que se ejecutará la aplicación
EXPOSE 3000

# Define el comando para ejecutar la aplicación
CMD ["./main"]