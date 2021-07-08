FROM golang:1.16.5-alpine

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY *.go .

RUN go build -o /GoWWW

EXPOSE 81

CMD [ "/GoWWW" ]
