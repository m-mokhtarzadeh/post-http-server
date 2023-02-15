FROM golang:1.20.0-alpine as build
COPY . $GOPATH/src/go-http-server
WORKDIR $GOPATH/src/go-http-server
RUN go build -v -o /go-http-server/server main.go

FROM alpine:3.17.1
WORKDIR /data
COPY --from=build /go-http-server/server /server
ENV SERVER_PORT=80
EXPOSE 80
ENTRYPOINT ["/server"]
