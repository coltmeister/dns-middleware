## TODO
- Come up with better name
- API to store domain names and associated protocols
- Persistent storage
- TLS connections
- Dynamically provision IP addresses
- Manage SSH tunnels

Each domain name is created with a name and private IP address. When a new
domain name is created, a new address must be provisioned on the vpn box. Each
protocol associated with a domain name will spawn an SSH tunnel to the actual
server. When the DNS server receives a request for a recognized domain, it
will respond with the provisioned IP address.

    Client                              Server
--------------------------------------------------------------------------------
1.  (domainName, privateIp)     ->
                                <-      Store data, provision IP address
--------------------------------------------------------------------------------
2.  (domainId, protocol)        ->
                                        Store data, open ssh tunnel on
                                <-      cooresponding port
--------------------------------------------------------------------------------
3.  DNS request                 ->
                                <-      Resolve to provisioned IP address
