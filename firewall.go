package nebula

type Firewall struct {
	UnsafeRoutes []struct {
		Route  string
		Via    string
		MTU    int
		Metric int
	}

	InboundRules []struct {
		Port  string
		Proto string
		Host  string
		Group string
	}

	OutboundRules []struct {
		Port  string
		Proto string
		Host  string
		Group string
	}
}

// DefaultFirewall allows outbound connections by default.
func DefaultFirewall() *Firewall {
	return &Firewall{
		OutboundRules: []struct {
			Port  string
			Proto string
			Host  string
			Group string
		}{{"any", "any", "any", "any"}},
	}
}

func (cc *Firewall) AddUnsafeRoute(Route, Via string, MTU, metric int) {
	cc.UnsafeRoutes = append(cc.UnsafeRoutes, struct {
		Route  string
		Via    string
		MTU    int
		Metric int
	}{
		Route:  Route,
		Via:    Via,
		MTU:    MTU,
		Metric: metric,
	})
}

func (cc *Firewall) AddInboundRule(Port, Proto, Host, Group string) {
	cc.InboundRules = append(cc.InboundRules, struct {
		Port  string
		Proto string
		Host  string
		Group string
	}{
		Port:  Port,
		Proto: Proto,
		Host:  Host,
		Group: Group,
	})
}

func (cc *Firewall) AddOutboundRule(Port, Proto, Host, Group string) {
	cc.OutboundRules = append(cc.OutboundRules, struct {
		Port  string
		Proto string
		Host  string
		Group string
	}{
		Port:  Port,
		Proto: Proto,
		Host:  Host,
		Group: Group,
	})
}
