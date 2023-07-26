<h1 align = "center"> API Gateway Project </h1>
<h3 align = "center"> Jia Xin and Yashma </h3>
<h3> Project Summary</h3>
This project is an API Gateway implementation that serves as a centralized entry point for accessing and managing multiple backend services. The API Gateway acts as a mediator between clients and the backend services, providing a unified interface and offering various features such as  load balancing, request transformation, and service discovery.

# USER GUIDE
## To clone the repository, use this command on terminal: 
	git clone https://github.com/yashma-sonara/API_Gateway

## To start Hertz Server
  ### Navigate into Hertz_Demo folder 
  Run following command: 

* `go run main.go`

## To start nacos server: 
  ### Follow Guide to Install nacos: 
	https://nacos.io/en-us/docs/quick-start.html

  ### Navigate into nacos directory
  * `cd distribution/target/nacos-server-[version]/nacos/bin` OR
  * `cd nacos/bin`

  ### Run the following command to start nacos: 
  * Windows
     * `startup.cmd -m standalone`
  * Linux/Mac
     * `sh startup.sh -m standalone`

## To register RPC servers: 
  1.   Open a new terminal 
  2.   Navigate to RPC_Server directory
  3.   Run following command: 
    
  * Windows: 
    *  `bash build.sh`
    *  `bash output/bootstrap.sh`
  * Linux/Git Bash: 
     *  `sh build.sh`
     *  `sh output/bootstrap.sh`
   

  You can view nacos registered services on: http://localhost:8848/nacos

## Send a HTTP request to the server
 Open a new terminal
 
 Run following command: 
  * request to call backend server:

	  curl -X GET "http://localhost:8888/ServiceA/methodB" -d "{\"userId\":\"12312\", \"message\":\"Hello World!\"}" -H "Content-type:application/json"

  * request to update IDL:

    curl -X [HTTP_REQUEST_METHOD] http://localhost:8888/[serviceName]/ -d “{\“file\”:\”[idl_filepath]\”}” -H "Content-Type: application/json"

---
 <h3 align="left">Languages and Tools:</h3>
<p align="left"> <a href="https://golang.org" target="_blank" rel="noreferrer"> <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="go" width="40" height="40"/> </a> <a href="https://www.w3.org/html/" target="_blank" rel="noreferrer"> <img src="https://avatars.githubusercontent.com/u/79236453?s=200&v=4" alt="html5" width="40" height="40"/> </a> <a href="https://reactjs.org/" target="_blank" rel="noreferrer"> <img src="https://avatars.githubusercontent.com/u/41446552?s=280&v=4" alt="react" width="40" height="40"/> </a> </p>
