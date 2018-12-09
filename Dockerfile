FROM golang:1.11.2-stretch

RUN mkdir -p /loft/api
COPY main.go /loft/api
COPY go.mod /loft/api
COPY go.sum /loft/api
WORKDIR /loft/api

RUN go build

CMD ["loft"]