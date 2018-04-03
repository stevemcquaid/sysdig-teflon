FROM golang:1.10.1-alpine3.7

MAINTAINER Steve McQuaid <steve@stevemcquaid.com>

ENV VERSION=1.0.0

RUN apk update && apk upgrade && \
    apk add --no-cache bash git

WORKDIR /go/src/teflon

# Caching large packages to speed up build
#RUN go-wrapper download -u github.com/golang/glog

RUN go get -d -v "k8s.io/apimachinery/pkg/api/errors"
RUN go get -d -v "k8s.io/apimachinery/pkg/apis/meta/v1"
RUN go get -d -v "k8s.io/client-go/kubernetes"
RUN go get -d -v "k8s.io/client-go/tools/clientcmd"

COPY src/ .

RUN go get -d -v ./...

RUN go install -v ./...

CMD ["teflon"]
