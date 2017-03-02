FROM ubuntu:xenial

WORKDIR /opt

ADD ./coredns ./coredns
ADD ./Corefile.docker ./Corefile
ADD ./vpndns.conf /etc/vpndns.conf

EXPOSE 53

ENTRYPOINT ["./coredns"]
