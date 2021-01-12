FROM golang

RUN go get -u github.com/DanielHauge/goSpace
RUN go get -u github.com/rs/xid

# Copy the server code into the container
WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# Make port 8787 available to the host
EXPOSE 31415, 8080, 80

CMD ["app"]