FROM golang:1.20-alpine

WORKDIR /app

ADD . ./
RUN go mod download

# ENV GIN_MODE=release

RUN go build -o build/program/app main.go

EXPOSE 80

CMD ["/app/build/program/app"]