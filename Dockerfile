FROM golang:1.21.3

WORKDIR /app

COPY . .

RUN go get
RUN go build main.go
RUN apt-get update && apt-get install -y redis-server

EXPOSE 1234
EXPOSE 6379

CMD ["redis-server", "--daemonize yes"]
CMD ["./main"]
