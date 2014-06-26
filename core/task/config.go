package task

import (
	"github.com/idealeak/goserver/core"
	"github.com/idealeak/goserver/core/basic"
)

var Config = Configuration{}

type WorkerConfig struct {
	basic.Options
	WorkerCnt int
}

type Configuration struct {
	basic.Options
	Worker WorkerConfig
}

func (c *Configuration) Name() string {
	return "task"
}

func (c *Configuration) Init() error {
	if c.Options.QueueBacklog <= 0 {
		c.Options.QueueBacklog = 1024
	}
	if c.Options.MaxDone <= 0 {
		c.Options.MaxDone = 1024
	}
	if c.Worker.Options.QueueBacklog <= 0 {
		c.Worker.Options.QueueBacklog = 1024
	}
	if c.Worker.Options.MaxDone <= 0 {
		c.Worker.Options.MaxDone = 1024
	}
	if c.Worker.WorkerCnt <= 0 {
		c.Worker.WorkerCnt = 8
	}
	return nil
}

func (c *Configuration) Close() error {
	return nil
}

func init() {
	core.RegistePackage(&Config)
}