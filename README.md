# Team Management

This application is responsible for team management

# Overview

A member can be an Employee or a Contractor. If it's an Employee we stores its role together, otherwise, we store its contract duration.

A member also has a name and tags.

---

# Table of content

- [Team Management](#team-management)
- [Overview](#overview)
- [Pre Requisites](#pre-requisites)
- [Running locally](#running-locally)
- [API Documentation](#api-documentation)
- [Unit Tests](#unit-tests)
- [Folders Structure](#folders-structure)
- [K8s](#k8s)
- [Exposing k8s app](#exposing-k8s-app)
- [K8s considerations](#k8s-considerations)
- [Logs](#logs)
- [Thoughts about improvement](#thoughts-about-improvement)

---

# Pre Requisites

Before running this application, you must have:

`Docker version 20.10`

`Docker-compose version 1.25`

`golang 1.18`

---

# Running locally

First of all, let's create an .env file:

```bash
cp .env.example .env
```

Now let's create a docker network, it is required for connecting our application with its dependencies:

```bash
docker network create team-management
```

Let's start a mongodb and mongo-express:

*Note: Mongo will be exposed at port `:27017` and mongo-express at `:8081`.*

```bash
docker-compose up -d
```

Now, let's build our application

```bash
docker build . -t team-management
```

And finally, let's run it:

```bash
docker run -p 8080:8080 --network=team-management team-management
```

**Now our API is ready to receive requests at http://localhost:8080/**

---

# API documentation

Here you can find an [OpenAPI spec file](.docs/team-management.postman_collection.json-OpenApi3Yaml.yaml) with all resources at [docs](.docs/) folder. Try it out in [Swagger Editor](https://editor.swagger.io/) or install [Swagger Plugin](https://marketplace.visualstudio.com/items?itemName=42Crunch.vscode-openapi) in your vscode to test it.

Also a [postman collection](.docs/team-management.postman_collection.json) can be found at [docs](.docs/) folder. See how to [import it](https://learning.postman.com/docs/getting-started/importing-and-exporting-data/#importing-postman-data) into your postman.

---

# Unit Tests

Execute the following command to run all unit tests:

```bash
go test ./...
```

---

# Folders structure

## Top level folder structure

    .
    ├── 
    ├── .docs                   # Documentation 
    ├── .k8s                    # Kubernetes manifests 
    ├── cmd                     # Application entrypoints 
    ├── members                 # Source files 
    ├── mocks                   # Mock files for our tests
    └── ...

## Members' folder structure

    .
    ├── ...
    ├── members                 # Source files
    │   ├── api                 # Api router, handlers, middlewares, etc
    │   ├── domain              # Entities, factories, etc
    │   ├── infrastructure      # Drivers, configs, adapters, etc
    │   ├── repository          # Repositories to serve our use cases
    │   ├── usecase             # All functionalities that our application supports
    │   └── utils               # Basic stuffs shared between all layers 
    └── ...

---

# K8s

First of all, you must have a `k8s cluster` running (for local tests, you can use either [minikube](https://minikube.sigs.k8s.io/docs/start/) or another one).

All the commands bellow are creating objects in namespace `default`.

To run our application in a kubernetes cluster, follow the nexts steps:

Deploy our mongo db:

```bash
make k8s-mongo
```

And then deploy our app:

```bash
make k8s-app
```

*Note: All the manifests which will be executed are found at [k8s](./.k8s) folder*

---

# Exposing k8s app

```bash
make k8s-app-port-forward
```

**Now our API at k8s is exposed at http://localhost:8080/**

Be sure that you aren't running anything at port `8080`.

If you want to bind another port, changes makefile at command `k8s-app-port-forward`.

You can also bind mongo db running the following command:

```bash
make k8s-mongodb-port-forward
```

It will be binded to port `27017`.

---

# K8s considerations

The k8s [Deployment manifest](./.k8s/app/deployment.yaml) found at [k8s app](./.k8s/app/) is using my own image: [dantunes/team-management](https://hub.docker.com/r/dantunes/team-management)

If you want to build your own image and push it to your registry, you will have to modify the [app deployment manifest](./.k8s/app/deployment.yaml) to use the new image.

For example:

Tag the image which you've built before:

```bash
docker tag team-management {your-account}/team-management
```

Push it:

```bash
docker push {your-account}/team-management
```

Modify the [deployment manifest](./.k8s/app/deployment.yaml) at line 19.

Then deploy it again:

```bash
make k8s-app
```

---

# Logs

All logs are logged in the stdout following the [ECS Logging standards](https://www.elastic.co/guide/en/ecs-logging/overview/current/intro.html).

---

# Thoughts about improvement

Given more time and incentives, I would like to:

* Implement integration tests for a better coverage of the features requirements.

* Implement a better filtering strategy and pagination for `GET /members` resource (which at moment are retrieving all the items matched in db)

* Use ConfigMap or Secrets at k8s mongodb manifests, which have credentials hard coded.

* Find another better solution for `type_data` field problemn. In this case, my solution needs to decode and encode it many times in different places(To decouple api from domain, domain from database access, and so on).

* Make a better use of golang context throughout the application.

---