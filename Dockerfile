FROM golang:alpine

WORKDIR /go/src/app

COPY ./main.go ./main.go

RUN apk update && apk add git

RUN go-wrapper download
RUN go-wrapper install

CMD ["go-wrapper", "run"]
