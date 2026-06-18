FROM golang:1.26.4-bookworm

WORKDIR /app

COPY . .

RUN go build -o exe .

CMD ["./exe"]