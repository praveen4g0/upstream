FROM golang:1.18-alpine

WORKDIR /app

COPY * ./

RUN go mod tidy && \
    go mod download && \
    go build -o randomGenerator

CMD [ "./randomGenerator" ]
