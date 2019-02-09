FROM golang:1.11.5

WORKDIR /go/src/app

ENV GO111MODULE=on
EXPOSE 8080

COPY . .

RUN go get
RUN go build

CMD ["blog"]