FROM alpine:latest

ARG GOPATH
WORKDIR /opt
ADD $GOPATH/bin/coredns coredns
ADD Corefile.docker
EXPOSE 53

CMD ["./coredns", "-conf", "Corefile.docker"]
