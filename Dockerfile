# syntax=docker/dockerfile:1

# This is a multi-stage image. First we build an image with the executable binary.
# There is a problem, the resulting image is too heavy because it contains the go toolkit, 
# which is useless after obtaining the executable binary.
# In order to obtain a smaller image, we build a second image, which will be used
# for deploying the app

### ------------- Deploy Image ------------- ###
FROM golang:1.19-alpine as build

WORKDIR /app

# Resolve project dependencies
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy all source files
COPY . ./

# Build executable from source files
RUN go build -o /go-seed

### ------------- Deploy Image ------------- ###
FROM alpine

WORKDIR /

COPY --from=build /go-seed /go-seed

EXPOSE 8080

ENTRYPOINT ["/go-seed"]