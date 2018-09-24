# CRUD APIs 
Simple Restful API on AWS
This project implements a simple Restful API on AWS using the following tech stack:
## Pre-requisites
* Serverless
* AWS CLI
* Go Dep
* Linux Packages zip & unzip Package
* NodeJs
* Git
##### Serverless
[Install Serverless framework](https://serverless.com/framework/docs/providers/aws/guide/quick-start/)
```
npm install -g serverless
npm install serverless-pseudo-parameters
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
##### Linux Packages zip & unzip Package
```
apt-get install -y zip
apt-get install -y unzip
```
##### NodeJs
[Installing Nodejs](https://nodejs.org/en/download/)
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
##### Git
Installation :
```
apt-get install -y git
```
## Build
Build the project using the following command:

    ```
    make bild
    ```
### Unit Test
Running UnitTests using the following command:

    ```
    make unit-test
    ```
## Deploy
Deploy the project using the following command:

    ```
    make deploy
    ```
## Integration Test
Running Integration Test using the following command

    ```
    make integration-test
    ```
## Logs
