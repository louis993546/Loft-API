FROM golang:1.11.2-stretch

RUN mkdir -p $GOPATH/src/github.com/louistsaitszho/loft
COPY . $GOPATH/src/github.com/louistsaitszho/loft/
WORKDIR $GOPATH/src/github.com/louistsaitszho/loft/server

RUN go build

CMD ["./server"]