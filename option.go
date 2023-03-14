package nebula

type Option func(*yamlConfig)

func WithListenHost(ip string, port uint16) Option {
	return func(o *yamlConfig) {
		o.Listen.Host = ip
		o.Listen.Port = int(port)
	}
}

func WithDevName(name string) Option {
	return func(o *yamlConfig) {
		o.Tun.Dev = name
	}
}

func WithLighthouse(nip, addr string) Option {
	return func(o *yamlConfig) {
		o.Lighthouse.Hosts = append(o.Lighthouse.Hosts, nip)
		o.StaticHostMap[nip] = append(o.StaticHostMap[nip], addr)
	}
}

func WithRelay(addrs []string) Option {
	return func(o *yamlConfig) {
		o.Relay.UseRelays = true
		o.Relay.Relays = addrs
	}
}

func WithFirewall(c *Firewall) Option {
	return func(y *yamlConfig) {
		for _, cfg := range c.InboundRules {
			y.Firewall.Inbound = append(y.Firewall.Inbound, struct {
				Port  string "yaml:\"port\""
				Proto string "yaml:\"proto\""
				Host  string "yaml:\"host,omitempty\""
				Group string "yaml:\"group,omitempty\""
			}(cfg))
		}

		for _, cfg := range c.OutboundRules {
			y.Firewall.Outbound = append(y.Firewall.Outbound, struct {
				Port  string "yaml:\"port\""
				Proto string "yaml:\"proto\""
				Host  string "yaml:\"host,omitempty\""
				Group string "yaml:\"group,omitempty\""
			}(cfg))
		}

		for _, cfg := range c.UnsafeRoutes {
			y.Tun.UnsafeRoutes = append(y.Tun.UnsafeRoutes, struct {
				Route  string `yaml:"route"`
				Via    string `yaml:"via"`
				MTU    int    `yaml:"mtu"`
				Metric int    `yaml:"metric"`
			}(cfg))
		}
	}
}