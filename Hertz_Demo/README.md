# Hertz_Demo

This is the Hertz server of the API Gateway. When a request is made to the server, the decode function is invoked which extracts the serviceName and method from the URL path of the request. This information is then printed to the console. The server listens for incoming requests on the host 127.0.0.1:8080.

## how to run
* `go mod init Hertz_Demo`

* `go mod tidy`

* `go build -o Hertz_Demo`

* `go run main.go`

## how to send request to server
* `curl -X GET http://localhost:8080/[serviceName]/[method] -d '[{"userID":"id"}]' -H "Content-Type: application/json"`

Modify content in [] according to service running
