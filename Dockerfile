FROM golang:1.21.3

WORKDIR /app

COPY . .

RUN go get -d -v ./...
RUN go install -v ./...
RUN go build main.go

EXPOSE 1234

CMD ["./main"]
