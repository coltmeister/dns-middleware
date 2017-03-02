package vpndns

import (
    "github.com/coredns/coredns/core/dnsserver"
    "github.com/coredns/coredns/middleware"
    "github.com/mholt/caddy"
)

func init() {
    caddy.RegisterPlugin("vpndns", caddy.Plugin{
        ServerType: "dns",
        Action: setup,
    })
}

func setup(c *caddy.Controller) error {
    c.Next()

    if c.NextArg() {
        return middleware.Error("vpndns", c.ArgErr())
    }

    dnsserver.GetConfig(c).AddMiddleware(func(next middleware.Handler) middleware.Handler {
        return Vpn{}
    })

    return nil
}
