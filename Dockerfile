FROM golang:1.16

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY *.go .

go build -o /GoWWW

EXPOSE 8080

CMD [ "/GoWWW" ]
