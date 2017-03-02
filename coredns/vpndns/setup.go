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

var cache ConcurrentMap = NewConcurrentMap()

func init() {
    caddy.RegisterPlugin("vpndns", caddy.Plugin{
        ServerType: "dns",
        Action: setup,
    })
}

func setup(c *caddy.Controller) error {
    if !c.NextArg() {
        return middleware.Error("vpndns", c.ArgErr())
    }

    c.Next()

    if c.NextArg() {
        return middleware.Error("vpndns", c.ArgErr())
    }

    conf := c.Val()
    go watchFile(conf, 2000, func(path string) error { return loadCache(path, &cache) })

    dnsserver.GetConfig(c).AddMiddleware(func(next middleware.Handler) middleware.Handler {
        return VpnDns{}
    })

    return nil
}

func loadCache(path string, kv KeyValueStore) error {
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
        kv.Put(domainName, ipAddress)
    }

    if err := scanner.Err(); err != nil {
        return err
    }

    return nil
}
