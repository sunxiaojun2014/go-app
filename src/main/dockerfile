FROM sunxiaojun2014/go-app:0.2
MAINTAINER sunxiaojun:go-app

#ENV GOPATH /go
#ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN go install main
ENTRYPOINT /go/bin/main

EXPOSE 8080
