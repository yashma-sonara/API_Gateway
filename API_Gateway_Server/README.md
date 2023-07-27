# Hertz_Demo

This is the Hertz server of the API Gateway. When a request is made to the server, the decode function is invoked which makes a RPC call based on the service and method name, and return a response. The server listens for incoming requests on the host 127.0.0.1:8888.

## how to run
* `go run main.go`

## how to send request to server
### request to call backend server
* `curl -X GET http://localhost:8888/[serviceName]/[method] -d '[{"userID":"id"}]' -H "Content-Type: application/json"`

### request to update IDL
* `curl -X [HTTP_REQUEST_METHOD] http://localhost:8888/[serviceName]/ -d “{\“file\”:\”[idl_filepath]\”}” -H "Content-Type: application/json"`


Modify content in [ ] according to service running and IDL file
