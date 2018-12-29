FROM golang:1.11.2-stretch

RUN mkdir -p /loft
COPY . /loft/
WORKDIR /loft/

# RUN go install 

# CMD ["./server"]

CMD ["go run server/server.go"]