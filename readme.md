### App structure
- main.go
- go.mod
- controller/
    - alertController.go
- repository/
    - alertRepository.go
- models/
    - alert.go
- docker/
    - Dockerfile
    - prepopulate.js

### Commands
go mod tidy
go run main.go


### Build Docker Service Image
> Change database.url value to "mongodb" in config.yaml file before building images.
#### Service alone
docker build --tag alert-go .
docker run --rm -p 3000:3000 alert-go
#### Mongo alone
docker build --tag mongodb docker/mongo
docker run --rm -p 27017:27017 mongodb
#### Services with docker-compse 
docker-compose build .
docker-compose up -d