FROM golang:1.12

WORKDIR /usr/src
COPY . /usr/src

RUN go build

ENTRYPOINT ["./microservice-course"]
