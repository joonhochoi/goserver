package signal

import (
	"github.com/idealeak/goserver/core"
)

var Config = Configuration{}

type Configuration struct {
	SupportSignal bool
}

func (c *Configuration) Name() string {
	return "signal"
}

func (c *Configuration) Init() error {
	return nil
}

func (c *Configuration) Close() error {
	return nil
}

func init() {
	core.RegistePackage(&Config)
}