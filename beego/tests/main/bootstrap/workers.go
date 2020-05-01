package bootstrap

import (
		"fmt"
		"github.com/WebGameLinux/bootstraps/beego/worker"
		"time"
)

type BeeGoWebWorkerBootStrap struct {
		worker.WorkersBootStrap
}

const BeeGoWorkerBootstrap = "web-bee-go"

func BeeGoWorkBootStrapName() string {
		return worker.Prefix() + BeeGoWorkerBootstrap
}

func NewBeeGoWebWorkerBootstrap() *BeeGoWebWorkerBootStrap {
		var boot = worker.NewWorkerBootstrap(BeeGoWorkBootStrapName())
		if boot.StartersLen() == 0 {
				boot.InitStarters(getBeeGoStarters())
		}
		if b, ok := boot.(*BeeGoWebWorkerBootStrap); ok {
				return b
		}
		if b, ok := boot.(*worker.WorkersBootStrap); ok {
				w := new(BeeGoWebWorkerBootStrap)
				b.BaseBootStrapDto.AWait = false
				w.WorkersBootStrap = *b
				return w
		}
		return nil
}

func getBeeGoStarters() []worker.Starter {
		var starters []worker.Starter

		return starters
}

func (this *BeeGoWebWorkerBootStrap) Start() {
		fmt.Println("hello beego start")
		// this.WorkersBootStrap.Start()
		run()
}

func (this *BeeGoWebWorkerBootStrap) Boot() {
		this.WorkersBootStrap.BaseBootStrapWrapper.BaseBootStrapDto.Booted = true
}

func run() {
		for {
				select {
				case v := <-time.NewTimer(2 * time.Second).C:
						fmt.Println("run now : ", v.String())
				}
		}
}
