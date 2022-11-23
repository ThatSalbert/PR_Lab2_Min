FROM golang:alpine

RUN mkdir /aggregator

WORKDIR /aggregator

COPY . .

RUN go build -o /go/bin/main

EXPOSE 8001

ENTRYPOINT ["/go/bin/main"]