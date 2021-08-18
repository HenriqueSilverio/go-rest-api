FROM golang:1.16.7-alpine3.14@sha256:3e6a2def9a57f23344a75bd71d9cd79726f0fbaf4b75330be5669773df0e9d4c AS build

ENV GOOS=linux \
  GOARCH=amd64

WORKDIR /go/go-rest-api

COPY . .

RUN go mod download && go build -o /bin/go-rest-api

FROM alpine:3.14.1@sha256:eb3e4e175ba6d212ba1d6e04fc0782916c08e1c9d7b45892e9796141b1d379ae

EXPOSE 8080/tcp

COPY --from=build /bin/go-rest-api /bin/go-rest-api

CMD ["/bin/go-rest-api"]
