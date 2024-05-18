FROM golang:latest

WORKDIR /app
COPY go.mod .
COPY go.sum .
COPY proto/ /app/proto
COPY servidor/main.go .

RUN go build main.go

EXPOSE 50051

CMD [ "./main" ]
