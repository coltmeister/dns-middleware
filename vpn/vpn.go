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
    state := request.Request{W: w, Req: r}
    return 0, nil
}
