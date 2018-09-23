# Restful-API
Simple Restful API on AWS
This project implements a simple Restful API on AWS using the following tech stack:
## Pre-requisites

### Linux Packages

##### zip & unzip Package
```
apt-get install -y zip
apt-get install -y unzip
```

##### Python
Installation :
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
Installation :
```
apt-get install -y nodejs
```
To verify installation :
```
npm --version
```
##### Serverless
[Install Serverless framework](https://serverless.com/framework/docs/providers/aws/guide/quick-start/)
```
npm install -g serverless
npm install serverless-pseudo-parameters
```
##### Git
Installation :
```
apt-get install -y git
```

##### Go Dep
On MacOS you can install or upgrade to the latest released version with Homebrew:
```
$ brew install dep
$ brew upgrade dep
```
On other platforms you can use the install.sh script:
```
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```
## Build
    make
### Unit Test
    Unit Test
## Deploy
    serverless deploy

## Integration Test
    Integration Test
## Logs
