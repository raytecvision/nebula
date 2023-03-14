// This package is a complete utter mess. Please don't judge me too harshly.
// Nebula does not expose clean APIs, so I needed to get creative.
package nebula

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/slackhq/nebula"
	"github.com/slackhq/nebula/config"
)

type Node interface {
	config() ([]byte, error)
}

type Client struct {
	ctx context.Context

	ctrl *nebula.Control
	cfg  *config.C

	node Node
}

func NewClient(ctx context.Context) *Client {
	return &Client{ctx: ctx}
}

func (c *Client) Reload(node Node) error {
	c.node = node

	ymlstr, err := c.node.config()
	if err != nil {
		return err
	}

	return c.cfg.ReloadConfigString(string(ymlstr))
}

func (c *Client) Start(node Node) error {
	c.node = node

	conf, err := c.node.config()
	if err != nil {
		return nil
	}

	c.cfg = config.NewC(logrus.StandardLogger())

	if err := c.cfg.LoadString(string(conf)); err != nil {
		return err
	}

	c.ctrl, err = nebula.Main(c.cfg, false, "raytecvision-nebula-client", logrus.StandardLogger(), nil)
	if err != nil {
		return err
	}

	c.ctrl.Start()

	return nil
}

func (r *Client) Stop() {
	// stop nebula client
	r.ctrl.Stop()
}
