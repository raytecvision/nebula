// This package is a complete utter mess. Please don't judge me too harshly.
// Nebula does not expose clean APIs, so I needed to get creative.
package nebula

import (
	"github.com/sirupsen/logrus"
	"github.com/slackhq/nebula"
	"github.com/slackhq/nebula/config"
)

type Configer interface {
	config() ([]byte, error)
}

type ClientOption func(*Client)

type Client struct {
	ctrl    *nebula.Control
	cfg     *config.C
	logger  *logrus.Logger
	version string
	node    Configer
}

func NewClient(opts ...ClientOption) *Client {
	c := &Client{logger: logrus.StandardLogger(), version: "raytecvision-nebula-client"}
	for _, o := range opts {
		o(c)
	}

	return c
}

func (c *Client) Reload(cf Configer) error {
	c.node = cf

	ymlstr, err := c.node.config()
	if err != nil {
		return err
	}

	return c.cfg.ReloadConfigString(string(ymlstr))
}

func (c *Client) Start(cfg Configer) error {
	c.node = cfg

	conf, err := c.node.config()
	if err != nil {
		return nil
	}

	c.cfg = config.NewC(logrus.StandardLogger())

	if err := c.cfg.LoadString(string(conf)); err != nil {
		return err
	}

	c.ctrl, err = nebula.Main(c.cfg, false, c.version, c.logger, nil)
	if err != nil {
		return err
	}

	c.ctrl.Start()

	return nil
}

func (r *Client) Stop() {
	// stop nebula client
	if r.ctrl != nil {
		r.ctrl.Stop()
	}
}
