FROM golang:alpine

RUN apk --update add git bash

WORKDIR /app

ENV GO111MODULE=on

COPY . ./
RUN go build

EXPOSE 8080 8081

CMD ["/app/go-react-example"]