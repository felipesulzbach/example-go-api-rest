# GoLang Web Service REST

API REST example in GoLang.

---

## About Golang

![](https://raw.githubusercontent.com/felipesulzbach/grpc-go-example/master/things/go.png)

### What is?

Golang, or simply Go, is an open source language created in 2009 by [Google](https://about.google/intl/en_US/) (by engineers [Rob Pike](https://en.wikipedia.org/wiki/Rob_Pike) and [Ken Thompson](https://en.wikipedia.org/wiki/Ken_Thompson)). The Go language was created with the goal of having **C** language performance but also focusing more readable and easier to program from more robust languages like **Java**.

### Some advantages of language

- Incredibly light in terms of memory usage;
- Suppose several concurrent processing because it uses goroutines instead of threads that are found in most programming languages. Competition is one of the language's strengths;
- Compiles very fast;
- Has garbage collector (has been incorporated into its core in order to prioritize performance);
- It is strongly typed.

GoLang intentionally leaves out many features of modern _OOP_ languages. Everything is divided into packages. [Google](https://about.google/intl/en_US/) technology has only _structs_ instead of _classes_.

### Some companies that have adopted Golang:

- Netflix
- The Economist
- IBM
- GitHub
- Uber
- Docker
- Dropbox
- OpenShift
- Twitter
- [Complete list by country (link here)](https://github.com/golang/go/wiki/GoUsers)

## About this project

### Prerequisites

- [Visual Studio Code](https://code.visualstudio.com/)
- [Golang](https://golang.org/)
- [Docker](https://hub.docker.com/search?offering=community&type=edition&operating_system=linux%2Cwindows)
- [Postman](https://www.postman.com/downloads/)

### Preparing the environment

The following technologies are critical for running/compiling application sources:

- Access the application directory from the terminal and execute:

  - > go get -u github.com/gorilla/mux
  - > go get github.com/lib/pq
  - > go get -u gorm.io/gorm
  - > go get github.com/spf13/viper

### Execute endpoints in Postman

- open *Postman* and import the `example-go-api-rest/resources/goLang.postman_collection.json` file;
- just select the endpoints and execute by clicking the **Send** button;

### Access PgAdmin

- open your browser end access the link `http://localhost:9090`;
- enter the access data:
  - user: `admin@admin`;
  - password: `admin`;
- right click on **Servers**, and **Create > Server...**;
- In the **General** tab inform:
  - Name: `local`;
- In the **Connection** tab inform:
  - Host name/address: `db_postgres`;
  - Port: `5432`;
  - Maintenance database: `postgres`;
  - Username: `postgres`;
  - Password: `postgres`;
- Click the **Save** button;
