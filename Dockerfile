FROM golang

RUN go get -u github.com/DanielHauge/goSpace
RUN go get -u github.com/beevik/guid


COPY . .

RUN go build -o app

EXPOSE 31415
EXPOSE 8080
EXPOSE 80

CMD ["./app"]