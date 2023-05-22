
### Build Image
docker build -t mongodb-alerts .

### Run image

docker run -d --name mongodb -p 27017:27017 mongodb-alerts
