# Restful-API
Simple Restful API on AWS
This project implements a simple Restful API on AWS using the following tech stack:
## Pre-requisites
Install Serverless framework:
https://serverless.com/framework/docs/providers/aws/guide/quick-start/
### Linux Packages

##### zip & unzip Package
```
apt-get install -y zip
apt-get install -y unzip
```

##### Python
```
apt install -y python3
apt install -y python3-pip
```
##### Curl
```
apt -y install  curl  
```
##### AWS CLI
[Installing the AWS Command Line Interface](https://docs.aws.amazon.com/cli/latest/userguide/installing.html)
```
curl "https://s3.amazonaws.com/aws-cli/awscli-bundle.zip" -o "awscli-bundle.zip"
```
To verify installation :
```
aws --version
```
##### Nodejs
[Installing Nodejs ](https://nodejs.org/en/)

Download :
```
curl -sL https://deb.nodesource.com/setup_8.x | bash -
```
```
apt-get install -y nodejs
```
To verify installation :
```
npm --version
```
##### Serverless
```
npm install -g serverless
```
```
npm install serverless-pseudo-parameters
```
##### Git
```
apt-get install -y git
```
	
##### Go Dep
	curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh

## Build
    make
### Unit Test
    Unit Test
## Deploy
    serverless deploy

## Integration Test
    Integration Test
## Logs
