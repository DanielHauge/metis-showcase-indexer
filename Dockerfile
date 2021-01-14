FROM golang

RUN go get -u github.com/DanielHauge/goSpace
RUN go get -u github.com/beevik/guid
RUN go get github.com/go-git/go-git
RUN go get -u github.com/droptheplot/abcgo
RUN go get -u github.com/elastic/go-elasticsearch
RUN go get github.com/robfig/cron


COPY . .

RUN go build -o app

EXPOSE 31415
EXPOSE 8080
EXPOSE 80

CMD ["./app"]