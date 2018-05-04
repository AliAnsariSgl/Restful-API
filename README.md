# Restful-API
Simple Restful API on AWS
This project implements a simple Restful API on AWS using the following tech stack:

    Serverless Framework ( https://serverless.com )
    Go language ( https://golang.org )
    AWS API Gateway
    AWS Lambda
    AWS DynamoDB

The API accepts the following JSON requests and produces the corresponding HTTP responses:
Request 1:

      HTTP POST
      URL: https://`API-GATEWAY-URL`/api/devices
      Body (application/json):
      {
      "id": "/devices/id1",
      "deviceModel": "/devicemodels/id1",
      "name": "Sensor",
      "note": "Testing a sensor.",
      "serial": "A020000102"
      }

Response 1 - Success:

        HTTP 201 Created
        Body (application/json):
        {
        "id": "/devices/id1",
        "deviceModel": "/devicemodels/id1",
        "name": "Sensor",
        "note": "Testing a sensor.",
        "serial": "A020000102"
        }

Response 1 - Failure 1:

        HTTP 400 Bad Request
        
If any of the payload fields are missing. Response body should have a descriptive error message for the client to be able to detect         the problem.
        Response 1 - Failure 2:

        HTTP 500 Internal Server Error

If any exceptional situation occurs on the server side.
         
Request 2:

      HTTP GET
      URL: https://`API-GATEWAY-URL`/api/devices/{id}
      Example: GET https://api123.amazonaws.com/api/devices/id1

Response 2 - Success:

      HTTP 200 OK
      Body (application/json):
      {
      "id": "/devices/id1",
      "deviceModel": "/devicemodels/id1",
      "name": "Sensor",
      "note": "Testing a sensor.",
      "serial": "A020000102"
       }

Response 2 - Failure 1:

     HTTP 404 Not Found

     If the request id does not exist.
     Response 2 - Failure 2:

     HTTP 500 Internal Server Error
     
 If any exceptional situation occurs on the server side.
 
## Getting Started
First things first, you'll be needing the [Golang](http://golang.org/) and [Serverless](http://serverless.com/) Framework installed, and an [AWS account](https://aws.amazon.com/) 

### Prerequisites
1. GO
2. Serverless

##### setup Go
[Download](https://golang.org/dl/) Go and follow the installation instructions.
To test your Go installation, open a new terminal and enter:
```
$ go version
```
##### setup Serverless
The intial setup was straightforward:

  1. Install NodeJS: [download](https://nodejs.org/en/download/) or using package manager.
  2. Install the Serverless Framework: 
   ```
   npm install -g serverless
   ```
  3. Setup an AWS account
  4. Sign up for AWS
  5. Create an IAM User
  6. [Install and setup AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/installing.html)
  7. [Go Dep](https://golang.github.io/dep/)
## Design

AWS Lambda functions are responsible for mentioned tasks:
```
   1. Create New Device for handling POST request
   2. getDevice for handling GET request
   3. These functions connect to a DynamoDB table named "devices"
``` 
## Deployment

1- Create a new service based on this template
```
serverless create -u template aws-go-dep -p myservice
````
2-  Compile function
````
cd myservice
%USERPROFILE%\go\bin\dep ensure
set GOOS=linux
go build -o bin/CreateDevice CreateDevice/main.go
go build -o bin/GetDevice GetDevice/main.go
%USERPROFILE%\go\bin\build-lambda-zip.exe -o bin/CreateDevice.zip bin/CreateDevice
%USERPROFILE%\go\bin\build-lambda-zip.exe -o bin/GetDevice.zip bin/GetDevice
``````
3-Deploy
```
serverless deploy
````
## Test

For testing by URL, replace you API-GATEWAY-URL with following curl commands:

#### NOTE :
Also you can use [POSTMAN](https://www.getpostman.com/) for testing
Create a new device:
```
curl -i -H "Content-Type: application/json" -X POST -d '{"id":"/devices/id1","deviceModel":"/devicemodels/id1","name":"Sensor","note":"note1","serial":"s1"}' https://API-GATEWAY-URL/devices
```

Get request using id:
```
curl -i https://API-GATEWAY-URL/devices/id1
```

