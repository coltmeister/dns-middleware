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

    /* Resolve A and AAAA records to a Google IP address */
    var rr dns.RR
    var g_ip string = "74.125.21.100"

    switch state.Family() {
        case 1:
            rr = new(dns.A)
            rr.(*dns.A).Hdr = dns.RR_Header{Name: state.QName(), Rrtype: dns.TypeA, Class: state.QClass()}
            rr.(*dns.A).A = net.ParseIP(g_ip).To4()
        case 2:
            rr = new(dns.AAAA)
            rr.(*dns.AAAA).Hdr = dns.RR_Header{Name: state.QName(), Rrtype: dns.TypeAAAA, Class: state.QClass()}
            rr.(*dns.AAAA).AAAA = net.ParseIP(g_ip)
    }

    a.Extra = []dns.RR{rr}

    /* Set response bits and write response to client */
    state.SizeAndDo(a)
    w.WriteMsg(a)

    return 0, nil
}
