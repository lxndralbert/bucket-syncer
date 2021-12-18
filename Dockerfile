FROM golang:1.17.4-buster

RUN apt update && apt upgrade -y

LABEL maintainer="Alexander Albert <alexander@vilkrig.com>"

WORKDIR /.

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]