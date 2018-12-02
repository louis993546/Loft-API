FROM golang:1.11.2-stretch

WORKDIR /go/src/app
COPY ..

RUN go get -d -v github.com/gorilla/mux
RUN go install -v github.com/gorilla/mux

CMD ["app"]