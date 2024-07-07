FROM golang:alpine3.20
LABEL maintainer="Adeboye"
RUN apk update && mkdir /go/src/app
WORKDIR /go/src/app
COPY . .
EXPOSE 8080
CMD ["go", "run", "http_server.go"]