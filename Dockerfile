FROM golang

ADD . /go/src/github.com/jasosa/wemper

RUN go install ./...

ENTRYPOINT /go/bin/wemp

#ARG buildtime_variable="tcp(host.docker.internal:3306)"
#ENV DBHOST=$buildtime_variable

EXPOSE 8080
