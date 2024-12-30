FROM golang:1.23

WORKDIR /app

COPY ./cmd ./cmd

RUN go build -o main ./cmd/main.go

EXPOSE 3000

CMD ["./main"]
