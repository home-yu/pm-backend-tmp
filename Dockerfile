FROM golang:latest

WORKDIR /app
COPY ./src /app

RUN go mod tidy

CMD [ "go", "run", "." ]