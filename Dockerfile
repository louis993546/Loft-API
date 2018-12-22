FROM golang:1.11.2-stretch

RUN mkdir -p /loft/graphql
COPY . /loft/graphql/
WORKDIR /loft/graphql/server

RUN go install 

CMD ["./server"]