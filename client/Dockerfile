FROM google/golang:latest

ADD main.go /go/src/client/main.go
WORKDIR /go/src/client
RUN go get
RUN go build -o client
ENTRYPOINT /go/src/client/client
