FROM ubuntu:xenial

WORKDIR /opt

ADD ./coredns ./coredns
ADD ./Corefile.docker ./Corefile

EXPOSE 53

ENTRYPOINT ["./coredns"]
