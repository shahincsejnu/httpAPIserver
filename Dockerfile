#------------building an optimzed docker image for the api server using the binary of the api and busybox:glibc image"-----
## created the binary of the api server using normal go build -o server .
FROM busybox:glibc

WORKDIR /app

COPY server .

EXPOSE 8080

CMD ["./server", "start"]


#------------building an optimized docker image for the server using the binary of the API server and golang:alpine image---------
#FROM golang:1.15-alpine3.7
## created the binary of the API server in my local machine with CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .
#FROM alpine:latest
#
#WORKDIR /app
#
#COPY server .
#
#EXPOSE 8080
#
#CMD ["./server", "start"]


#------------building an optimized docker image for the server using the binary of the API server and ubuntu image-----
#FROM ubuntu:latest
#
#WORKDIR /app
#
#COPY server .
#
#EXPOSE 8080
#
#CMD ["./server", "start"]


#-------------- building an optimized docker image for the server using multi-stage builds -----------
##--first stage of the multi-stage build will use the golang:latest image and build the application--
# start from the latest golang base image
#FROM golang:latest as builder
#
## add miantainer info
#LABEL maintainer="Sahadat Hossain"
#
## set the current working directory inside the container
#WORKDIR /app
#
## copy go mod and sum files
#COPY go.mod go.sum ./
#
## download all dependencies, dependencies will be cached if the go.mod and go.sum files are not changed
#RUN go mod download
#
## Copy the source from the current directory to the Working Directory inside the container
#COPY . .
#
## build the Go app (API server)
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server .
#
############ start a new stage from scracthc ###########
#FROM alpine:latest
#
#RUN apk --no-cache add ca-certificates
#
#WORKDIR /root/
#
## copy the pre-built binary file from the previous stage
#COPY --from=builder /app/server .
#
## Expose port 8080 to the outside world
#EXPOSE 8080
#
## command to run the executable
#CMD ["./server", "start"]


#--------------------- used the binary file for creting docker image, a little lighter than the next one ------------

## start from the latest golang base image
#FROM golang:latest
#
## add maintainer info
#LABEL maintainer="Sahadat Hossain"
#
## set the current working directory inside the container
#WORKDIR /app
#
## copy the binary file of the API server in current directory
#COPY server .
#
## expose port 8080 to outside the world
#EXPOSE 8080
#
## command to run the executable
#CMD ["./server", "start"]


#-------------------------- heavy weight and copied all the source file in docker container ------------------
## start from the latest golang base image
#FROM golang:latest
#
## add maintainer info
#LABEL maintainer="Sahadat Hossain"
#
## set the current working directory inside the container
#WORKDIR /app
#
## copy go mod and sum files
#COPY go.mod go.sum ./
#
## download all dependencies, dependencies will be cached if the go.mod and go.sum files are not changed
#RUN go mod download
#
## copy the source from the current directory to the working directory inside the container
#COPY . .
#
## build the Go app (API server)
#RUN go build -o server .
#
## expose port 8080 to outside the world
#EXPOSE 8080
#
#
## command to run the executable
#CMD ["./server", "start"]