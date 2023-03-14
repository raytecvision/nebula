package nebula

import (
	"sync"

	"github.com/slackhq/nebula/cert"
	"gopkg.in/yaml.v2"
)

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
func (n *Node) Cert() (*cert.NebulaCertificate, error) {
	c, _, err := cert.UnmarshalNebulaCertificateFromPEM([]byte(n.cfg.PKI.Cert))
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (n *Node) UpdateCredentials(priv, cert []byte) {
	n.m.Lock()
	defer n.m.Unlock()

	n.cfg.PKI.Key = string(priv)
	n.cfg.PKI.Cert = string(cert)
}

func (l *Node) UpdateCA(crt []byte) {
	l.m.Lock()
	defer l.m.Unlock()

	l.cfg.PKI.CA = string(crt)
}

func (n *Node) config() ([]byte, error) {
	return yaml.Marshal(n.cfg)
}
