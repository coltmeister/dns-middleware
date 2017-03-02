package vpndns

import (
    "net"
    "github.com/coredns/coredns/request"
    "github.com/miekg/dns"
    "golang.org/x/net/context"
)

type VpnDns struct {}

func (vpndns VpnDns) Name() string {
    return "vpndns"
}

func (vpndns VpnDns) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
    /* Get current request state */
    state := request.Request{W: w, Req: r}

    /* Create response message */
    a := &dns.Msg{}
    a.SetReply(r)
    a.Compress = true
    a.Authoritative = true

    /* Cache lookup */
    ip, ok := cache.Get(state.Name())

    if !ok {
        return dns.RcodeServerFailure, nil
    }

    /* Build response */
    var rr dns.RR

    switch state.Family() {
    case 1: // ipv4
        rr = &dns.A{}
        rr.(*dns.A).Hdr = dns.RR_Header {
            Name: state.QName(),
            Rrtype: dns.TypeA,
            Class: state.QClass(),
        }
        rr.(*dns.A).A = net.ParseIP(ip).To4()
    case 2: // ipv6
        rr = &dns.AAAA{}
        rr.(*dns.AAAA).Hdr = dns.RR_Header{
            Name: state.QName(),
            Rrtype: dns.TypeAAAA,
            Class: state.QClass(),
        }
        rr.(*dns.AAAA).AAAA = net.ParseIP(ip)
    }

    a.Extra = []dns.RR{rr}

    /* Write response out to client */
    state.SizeAndDo(a)
    w.WriteMsg(a)

    return 0, nil
}
