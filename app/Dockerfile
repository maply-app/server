FROM golang:1.19-alpine AS builder

WORKDIR /usr/src/app

COPY ./go.mod ./go.sum ./
RUN go mod download

COPY ./entrypoint.sh /usr/src/app/entrypoint.sh
RUN chmod +x entrypoint.sh

COPY . .

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

ENTRYPOINT ["sh", "/usr/src/app/entrypoint.sh"]
