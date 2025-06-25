package core

import "time"

type Config struct {
	Scheme   string
	Endpoint string
	Timeout  time.Duration
}

func NewConfig() *Config {
	return &Config{
		Scheme:   "https",
		Endpoint: DefaultEndpoint,
		Timeout:  time.Second * 30,
	}
}

func (c *Config) WithScheme(scheme string) *Config {
	c.Scheme = scheme
	return c
}

func (c *Config) WithEndpoint(endpoint string) *Config {
	c.Endpoint = endpoint
	return c
}

func (c *Config) WithTimeout(timeout time.Duration) *Config {
	c.Timeout = timeout
	return c
}
