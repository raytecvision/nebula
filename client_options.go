package nebula

import "github.com/sirupsen/logrus"

func WithLogger(l *logrus.Logger) ClientOption {
	return func(c *Client) {
		c.logger = l
	}
}

func WithBuildVersion(version string) ClientOption {
	return func(c *Client) {
		c.version = version
	}
}
