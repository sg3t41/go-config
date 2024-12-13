package config

import (
	"github.com/sg3t41/gocnf/strategy"
)

type config struct {
	FilePath string
	Strategy strategy.IStrategy
}

func NewConfig() *config {
	return &config{}
}

func (c *config) SetStrategy(strategy strategy.IStrategy) *config {
	c.Strategy = strategy
	return c
}

func (c *config) SetFilePath(path string) *config {
	c.FilePath = path
	return c
}

func (c *config) Unmarshal(out any) error {
	in, err := c.Strategy.Load(c.FilePath)
	if err != nil {
		return err
	}
	return c.Strategy.Unmarshal(in, out)
}
