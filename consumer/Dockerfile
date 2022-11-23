FROM golang:alpine

RUN mkdir /consumer

WORKDIR /consumer

COPY . .

RUN go build -o /go/bin/main

EXPOSE 8002

ENTRYPOINT ["/go/bin/main"]