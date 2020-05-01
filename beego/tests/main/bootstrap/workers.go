package bootstrap

import (
		"fmt"
		base "github.com/WebGameLinux/bootstraps/beego"
		"github.com/WebGameLinux/bootstraps/beego/worker"
		"github.com/astaxie/beego"
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
				b.BaseBootStrapDto.Async = false
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
		beego.Run()
}

func (this *BeeGoWebWorkerBootStrap) Boot() {
		this.WorkersBootStrap.BaseBootStrapWrapper.BaseBootStrapDto.Booted = true
		this.Container("bootstrap.startAt", time.Now())
		this.Container("bootstrap.class", base.ClassName(this))
}

