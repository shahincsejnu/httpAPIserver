#--------------------------- used the binary file for creting docker image -------------------

# start from the latest golang base image
FROM golang:latest

# add maintainer info
LABEL maintainer="Sahadat Hossain"

# set the current working directory inside the container
WORKDIR /app

# copy the binary file of the API server in current directory
COPY server .

# expose port 8080 to outside the world
EXPOSE 8080

# command to run the executable
CMD ["./server", "start"]


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