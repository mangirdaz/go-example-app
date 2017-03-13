# Container Development Example App

## Introduction

> This is small test application to show containers development capabilities comparing docker-compose and Openshift Container Platform. 

## Pre-Requisites 
    docker 1.13.1+
    docker-compose 1.11+
    Openshift cluster running. Minishift, Openshift Dedicates, Origin, etc 1.4/3.4+
    
## Installation

Build steps:

        1. Build Containers
        2. Deploy container
        3. Expose container

        4. Debug container

## Openshift Container Platform:
        1. make ocp-build
        2. make ocp-run

## Docker-Compose:
        1. make docker-build
        2. make docker-run

## Under the hood: 

#### Docker
Build Binaries:

        CGO_ENABLED=0 GOOS=linux go build -o api .
        CGO_ENABLED=0 GOOS=linux go build -o fe .
        
Build Images:

        docker-compose build

Run Images:

    	docker-compose up 	

### Openshift Container Platform
Build Binaries:

        CGO_ENABLED=0 GOOS=linux go build -o api .
        CGO_ENABLED=0 GOOS=linux go build -o fe .
        
(one time activity) Create build artifacts:

        oc new-build --name fe --binary
        oc new-build --name api --binary
        
Execute build:

	    oc start-build fe --from-file=fe/ --follow
        oc start-build api --from-file=api/ --follow

Run Images:

        oc new-app --name api --image-stream=api -e=API_IP=0.0.0.0 -e=API_PORT=8080
        oc new-app --name fe --image-stream=fe -e=FE_IP=0.0.0.0 -e=FE_PORT=8080 -e=API_SVC=http://api-demo.apps.192.168.2.187.xip.io

Expose:

        oc expose service fe 
        oc expose service api
