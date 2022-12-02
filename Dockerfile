FROM golang:1.19

RUN mkdir /app

WORKDIR /app

COPY go.mod /app
COPY go.sum /app
RUN go mod download

ADD . /app

RUN go build -o main ./server/cmd/

CMD ["/app/main"]