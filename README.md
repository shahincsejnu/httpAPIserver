# httpAPIserver

## RESTful HTTP API server using [Go](https://github.com/golang), [Cobra CLI](https://github.com/spf13/cobra), [gorilla mux](https://github.com/gorilla/mux)

## Description

This is a basic RESTful API server, build with Golang. In this API server I have implemented Cobra CLI for running the API from the CLI and also used gorilla mux instead of Go net/http.

### Installation 

`go install github.com/shahincsejnu/httpAPIserver`

### Basic Authentication

- give username : `admin` and password : `admin` for each query to the api endpoint otherwise access will be denied

### Run by CLI Commands

- start the API in default port : 8080 by `httpAPIserver apiStart`
- start the API in your given port by `httpAPIserver apiStart -p=8088`, give your port number in the place of 8088


### API Endpoints Testing

- Primarily tested the API endpoints by [Postman](https://github.com/postmanlabs)
- E2E Testing.
    - added unit testing for this API
    - Checked response status code with our expected status code


