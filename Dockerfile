FROM golang:1.18.4-alpine3.10
RUN mkdir /sampleapp
ADD . /sampleapp
WORKDIR /sampleapp

