FROM golang:1.8.5-alpine3.6

MAINTAINER Steve McQuaid <steve@stevemcquaid.com>

ENV VERSION=1.0.0

RUN apk update && apk upgrade && \
    apk add --no-cache bash git

WORKDIR /go/src/teflon

# Caching large packages to speed up build
RUN go-wrapper download -u github.com/golang/glog

COPY src/ .

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

CMD ["go-wrapper", "run"] # ["app"]
