FROM golang:1.20

WORKDIR /

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
