# httpAPIserver

## RESTful HTTP API server using [Go](https://github.com/golang), [Cobra CLI](https://github.com/spf13/cobra), [gorilla mux](https://github.com/gorilla/mux)

### Description

This is a basic RESTful API server, build with Golang. In this API server I have implemented Cobra CLI for running the API from the CLI and also used gorilla mux instead of Go net/http.


[![Go Report Card](https://goreportcard.com/badge/github.com/shahincsejnu/httpAPIserver)](https://goreportcard.com/report/github.com/shahincsejnu/httpAPIserver)

------------ 

### Installation

- `git clone https://github.com/shahincsejnu/httpAPIserver.git`
- `cd httpAPIserver`
- `go install httpAPIserver`

---------------

### Run by CLI Commands

- start the API in default port : 8080 by `httpAPIserver start`
- start the API in your given port by `httpAPIserver start -p=8088`, give your port number in the place of 8088


--------------

### The Endpoints of this REST API

|Endpoint | Function | Method | StatusCode | Authentication|
|-----|-----|-----|-----|-----|
|`/api/login`| logIn | GET | StatusOK, StatusUnauthorized | Basic
|`/api/articles` | getAllArticles | GET | StatusOK, StatusUnauthorized | JWT|
|`/api/article` | addNewArticle | POST | StatusCreated, StatusUnauthorized | JWT|
|`/api/article/{id}` | deleteArticle | DELETE | StatusOK, StatusNoContent, StatusUnauthorized | JWT|
|`/api/article/{id}` | updateArticle | PUT | StatusCreated, StatusNoContent, StatusUnauthorized | JWT|
|`/api/article/{id}` | getSingleArticle | GET | StatusOK, StatusNoContent, StatusUnauthorized | JWT|


----------------

### Data Model

* Article Model
```
    type Article struct {
    	ID       string    `json:"id"`
    	Title    string    `json:"title"`
    	Body     string    `json:"body"`
    	Author   Author    `json:"author"`
    }
```

* Author Model
```
    type Author struct {
    	ID       string    `json:"id"`
    	Name     string    `json:"name"`
    	Rating   float64   `json:"rating"`
    }
```

----------------

### Basic Authentication

- implemented basic authentication middleware
- give username : `admin` and password : `admin` for each query to the api endpoint otherwise access will be denied

----------------

### JWT Authentication

- implemented JWT authentication
- first of all user need to hit `/api/login` endpoint with basic authentication then a token will be given and with that token for specific time user can do other request
----------------



### curl commands without authentication (it is valid for version : v1.0.0)

#### Run the API server

- `httpAPIserver start --port=8080`

#### Get all articles

- `curl -X GET http://localhost:8080/api/articles`

#### Get single article

- `curl -X GET http://localhost:8080/api/article/1`

#### Create new article

```
curl -X POST -H "Content-Type:application/json" -d '{"id":"10","title":"oka","body":"none","author":{"id":"11","name":"nobody","rating":10}}' http://localhost:8080/api/article
```

#### Update any article

```
curl -X PUT -H "Content-Type:application/json" -d '{"id":"10","title":"update","body":"me","author":{"id":"11","name":"somebody","rating":9}}' http://localhost:8080/api/article/10
```

#### Delete an article

- `curl -X DELETE -H http://localhost:8080/api/article/1`

----------------


### curl commands with Basic Authentication (it is valid for version : v1.0.1 and login function of version : v1.0.2)

#### Get all articles

- `curl -X GET --user admin:admin http://localhost:8080/api/articles`

#### Get single article

- `curl -X GET --user admin:admin http://localhost:8080/api/article/1`

#### Create new article

```
curl -X POST --user admin:admin -H "Conten-Type:application/json" -d '{{"id":"10","title":"update","body":"me","author":{"id":"11","name":"somebody","rating":9}}' http://localhost:8080/api/article
```

#### Update an article

```
curl -X PUT --user admin:admin -H "Conten-Type:application/json" -d '{{"id":"10","title":"update","body":"me","author":{"id":"11","name":"somebody","rating":9}}' http://localhost:8080/api/article/1
```

#### Delete an article

- `curl -X DELETE --user admin:admin http://localhost:8080/api/article/1`

----------------


### curl commands with JWT authentication

#### Get jwt token via login with basic authentication

- `curl -X GET --user admin:admin http://localhost:8080/api/login`

#### Get all articles

- `curl -X GET -H "Token: <jwt_token>" http://localhost:8080/api/articles`

#### Get single article of id 1

- `curl -X GET -H "Token: <jwt_token>" http://localhost:8080/api/article/1`

#### Create new article

```
curl -X POST -H "Token: <jwt_token>" -H "Content-Type:application/json" -d '{{"id":"10","title":"update","body":"me","author":{"id":"11","name":"somebody","rating":9}}' http://localhost:8080/api/article
```

#### Update an article

```
curl -X POST -H "Token: <jwt_token>" -H "Content-Type:application/json" -d '{{"id":"10","title":"update","body":"me","author":{"id":"11","name":"somebody","rating":9}}' http://localhost:8080/api/article/1
```

#### Delete an article whose id is 1

- `curl -X DELETE -H "Token: <jwt_token>" http://localhost:8080/api/article/1`


----------------

### API Endpoints Testing

- Primarily tested the API endpoints by [Postman](https://github.com/postmanlabs)
- E2E Testing.
    - added unit testing for this API
    - Checked response status code with our expected status code
