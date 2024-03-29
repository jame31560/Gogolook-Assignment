FROM golang:1.20

WORKDIR /app

COPY . .

RUN go mod download && go test -v ./...

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
