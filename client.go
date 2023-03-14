package nebula

import "sync"

type Node struct {
	m   sync.Mutex
	cfg *yamlConfig
}

func NewNode(priv, cert, ca []byte, opts ...Option) *Node {
	l := &Node{
		cfg: defaultConfig(),
	}

	l.cfg.Lighthouse.AmLighthouse = false

	l.cfg.PKI.Key = string(priv)
	l.cfg.PKI.Cert = string(cert)
	l.cfg.PKI.CA = string(ca)

	// Apply all the additional options.
	for _, f := range opts {
		f(l.cfg)
	}

	return l
}
