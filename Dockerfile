FROM golang:1.11.2-stretch

RUN mkdir -p /go/src/github.com/louistsaitszho/loft
COPY main.go /go/src/github.com/louistsaitszho/loft
WORKDIR /go/src/github.com/louistsaitszho/loft

RUN go get -d -v github.com/gorilla/mux
RUN go install -v github.com/gorilla/mux
RUN go install

CMD ["loft"]