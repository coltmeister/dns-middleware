package vpndns

import (
    "os"
    "bufio"
    "strings"
    "errors"
    "github.com/coredns/coredns/core/dnsserver"
    "github.com/coredns/coredns/middleware"
    "github.com/mholt/caddy"
)

type NameMap map[string]string
var cache NameMap = make(NameMap)

func init() {
    caddy.RegisterPlugin("vpndns", caddy.Plugin{
        ServerType: "dns",
        Action: setup,
    })
}

func loadCache(path string, m NameMap) error {
    f, err := os.Open(path)

    if err != nil {
        return err
    }

    defer f.Close()
    scanner := bufio.NewScanner(f)

    for scanner.Scan() {
        line := scanner.Text()
        split := strings.Split(line, "\t")

        if len(split) != 2 {
            return errors.New("Malformed vpndns.conf")
        }

        domainName, ipAddress := split[0], split[1]
        m[domainName] = ipAddress
    }

    if err := scanner.Err(); err != nil {
        return err
    }

    return nil
}

func setup(c *caddy.Controller) error {
    loadCache("/etc/vpndns.conf", cache)
    c.Next()

    if c.NextArg() {
        return middleware.Error("vpndns", c.ArgErr())
    }

    dnsserver.GetConfig(c).AddMiddleware(func(next middleware.Handler) middleware.Handler {
        return VpnDns{}
    })

    return nil
}
