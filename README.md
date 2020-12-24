# httpAPIserver

## RESTful HTTP API server using [Go](https://github.com/golang), [Cobra CLI](https://github.com/spf13/cobra), [gorilla mux](https://github.com/gorilla/mux)

## Description

This is a basic RESTful API server, build with Golang. In this API server I have implemented Cobra CLI for running the API from the CLI and also used gorilla mux instead of Go net/http.

------------ 

### Installation 

`go install github.com/shahincsejnu/httpAPIserver`


--------------

### The Endpoints of this REST API

|Endpoint | Function | Method | StatusCode | Authentication|
|-----|-----|-----|-----|-----|
|`/api/articles` | getAllArticles | GET | StatusOK, StatusUnauthorized | Basic|
|`/api/article` | addNewArticle | POST | StatusCreated, StatusUnauthorized | Basic|
|`/api/article/{id}` | deleteArticle | DELETE | StatusOK, StatusNoContent, StatusUnauthorized | Basic|
|`/api/article/{id}` | updateArticle | PUT | StatusCreated, StatusNoContent, StatusUnauthorized | Basic|
|`/api/article/{id}` | getSingleArticle | GET | StatusOK, StatusNoContent, StatusUnauthorized | Basic|


----------------

### Basic Authentication

- implemented basic authentication middleware
- give username : `admin` and password : `admin` for each query to the api endpoint otherwise access will be denied

----------------

### JWT Authentication

- implemented JWT authentication
- first of all user need to hit `/api/login` endpoint with basic authentication then a token will be given and with that token for specific time user can do other request
----------------


### Run by CLI Commands

- start the API in default port : 8080 by `httpAPIserver start`
- start the API in your given port by `httpAPIserver start -p=8088`, give your port number in the place of 8088

---------------


### curl Commands

#### Read all articles

`curl --user admin:admin -s -X GET http://localhost:8080/api/articles`

#### Read an article

`curl --user admin:admin -s -X GET http://localhost:8080/api/article/{id}`



----------------

### API Endpoints Testing

- Primarily tested the API endpoints by [Postman](https://github.com/postmanlabs)
- E2E Testing.
    - added unit testing for this API
    - Checked response status code with our expected status code
