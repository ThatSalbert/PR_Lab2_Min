FROM golang:alpine

RUN mkdir /producer

WORKDIR /producer

COPY . .

RUN go build -o /go/bin/main

EXPOSE 8000

ENTRYPOINT ["/go/bin/main"]