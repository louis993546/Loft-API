# Use go 1.11.2
FROM golang:1.11.2-stretch

# Copy all files into the image's GOPATH
RUN mkdir -p $GOPATH/src/github.com/louistsaitszho/loft
COPY . $GOPATH/src/github.com/louistsaitszho/loft/
WORKDIR $GOPATH/src/github.com/louistsaitszho/loft/server

# Build the application (which generates a executable binary)
RUN go build

# Command for the image to run when it triggers
CMD ["./server"]