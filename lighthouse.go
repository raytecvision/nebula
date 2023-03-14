package nebula

import (
	"sync"

	"github.com/slackhq/nebula/cert"
	"gopkg.in/yaml.v2"
)

type Lighthouse struct {
	m   sync.Mutex
	cfg *yamlConfig
}

func NewLighthouse(priv, cert, ca []byte, opts ...Option) *Lighthouse {
	l := &Lighthouse{
		cfg: defaultConfig(),
	}

	l.cfg.Lighthouse.AmLighthouse = true

	l.cfg.PKI.Key = string(priv)
	l.cfg.PKI.Cert = string(cert)
	l.cfg.PKI.CA = string(ca)

	// Apply all the additional options.
	for _, f := range opts {
		f(l.cfg)
	}

	return l
}

func (l *Lighthouse) Cert() (*cert.NebulaCertificate, error) {
	c, _, err := cert.UnmarshalNebulaCertificateFromPEM([]byte(l.cfg.PKI.Cert))
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (l *Lighthouse) UpdateNodeCredentials(priv, cert []byte) {
	l.m.Lock()
	defer l.m.Unlock()

	l.cfg.PKI.Key = string(priv)
	l.cfg.PKI.Cert = string(cert)
}

func (l *Lighthouse) UpdateCA(crt []byte) {
	l.m.Lock()
	defer l.m.Unlock()

	l.cfg.PKI.CA = string(crt)
}

func (l *Lighthouse) config() ([]byte, error) {
	return yaml.Marshal(l.cfg)
}
