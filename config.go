package nebula

type firewallRules struct {
	Port  string `yaml:"port"`
	Proto string `yaml:"proto"`
	Host  string `yaml:"host,omitempty"`
	Group string `yaml:"group,omitempty"`
}

// This struct mimics a YAML configuration.
type yamlConfig struct {
	PKI struct {
		CA   string `yaml:"ca"`
		Cert string `yaml:"cert"`
		Key  string `yaml:"key"`
	} `yaml:"pki"`
	StaticHostMap map[string][]string `yaml:"static_host_map"`
	Lighthouse    struct {
		AmLighthouse bool     `yaml:"am_lighthouse"`
		Interval     int      `yaml:"interval"`
		Hosts        []string `yaml:"hosts"`
	} `yaml:"lighthouse"`
	Listen struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"listen"`
	Punchy struct {
		Punch   bool `yaml:"punch"`
		Respond bool `yaml:"respond"`
	} `yaml:"punchy"`
	Tun struct {
		Disabled           bool   `yaml:"disabled"`
		Dev                string `yaml:"dev"`
		DropLocalBroadCast bool   `yaml:"drop_local_broadcast"`
		DropMulticast      bool   `yaml:"drop_multicast"`
		TxQueue            int    `yaml:"tx_queue"`
		MTU                int    `yaml:"mtu"`
		UnsafeRoutes       []struct {
			Route  string `yaml:"route"`
			Via    string `yaml:"via"`
			MTU    int    `yaml:"mtu"`
			Metric int    `yaml:"metric"`
		} `yaml:"unsafe_routes"`
	} `yaml:"tun"`
	Logging struct {
		Level  string `yaml:"level"`
		Format string `yaml:"format"`
	} `yaml:"logging"`
	Relay struct {
		UseRelays bool     `yaml:"use_relays"`
		Relays    []string `yaml:"relays"`
	} `yaml:"relay"`
	Firewall struct {
		Conntrack struct {
			TCPTimeout     string `yaml:"tcp_timeout"`
			UDPTimeout     string `yaml:"udp_timeout"`
			DefaultTimeout string `yaml:"default_timeout"`
			MaxConnections int    `yaml:"max_connections"`
		} `yaml:"conntrack"`
		Outbound []firewallRules `yaml:"outbound"`
		Inbound  []firewallRules `yaml:"inbound"`
	} `yaml:"firewall"`
}

func defaultConfig() *yamlConfig {
	y := &yamlConfig{}

	y.Lighthouse.AmLighthouse = false
	y.Lighthouse.Interval = 60
	y.Listen.Host = "0.0.0.0"
	y.Listen.Port = 0
	y.Punchy.Punch = true
	y.Punchy.Respond = true
	y.Tun.Disabled = false
	y.Tun.Dev = "nebula1"
	y.Tun.DropLocalBroadCast = false
	y.Tun.DropMulticast = false
	y.Tun.TxQueue = 500
	y.Tun.MTU = 1300
	y.Logging.Format = "text"
	y.Logging.Level = "info"

	y.Firewall.Conntrack = struct {
		TCPTimeout     string "yaml:\"tcp_timeout\""
		UDPTimeout     string "yaml:\"udp_timeout\""
		DefaultTimeout string "yaml:\"default_timeout\""
		MaxConnections int    "yaml:\"max_connections\""
	}{
		TCPTimeout:     "12m",
		UDPTimeout:     "3m",
		DefaultTimeout: "10m",
		MaxConnections: 100000,
	}

	// Empty firewall rules.
	y.Firewall.Inbound = []firewallRules{}
	y.Firewall.Outbound = []firewallRules{}

	return y
}