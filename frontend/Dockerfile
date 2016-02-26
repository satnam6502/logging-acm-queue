FROM google/golang:latest

ADD main.go /go/src/frontend/main.go
WORKDIR /go/src/frontend
RUN go get
RUN go build -o frontend
ENTRYPOINT /go/src/frontend/frontend

EXPOSE 8080
