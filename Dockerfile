FROM golang:1.13 as build

ARG GOARCH=amd64

WORKDIR /go/src/app
COPY . .

RUN GO111MODULE=on go get
RUN GO111MODULE=on GOOS=linux GOARCH=${GOARCH} go build -ldflags="-w -s"
RUN chmod +x /go/src/app/base36

FROM debian:buster-slim

COPY --from=build /go/src/app/base36 /usr/local/bin/base36
ENTRYPOINT [ "base36" ]
