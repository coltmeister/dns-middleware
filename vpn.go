package vpn

import (
    "github.com/coredns/coredns/request"
    "github.com/miekg/dns"
    "golang.org/x/net/context"
)

type Vpn struct {}

func (vpn Vpn) Name() string {
    return "vpn"
}

func (vpn Vpn) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
    /* Get current request state */
    state := request.Request{W: w, Req: r}

    /* Create response message */
    a := &dns.Msg{}
    a.SetReply(r)
    a.Compress = true
    a.Authoritative = true

    /* Set response bits and write response to client */
    state.SizeAndDo(a)
    w.WriteMsg(a)

    return 0, nil
}
