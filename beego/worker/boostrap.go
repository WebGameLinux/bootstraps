package worker

import (
		base "github.com/WebGameLinux/bootstraps/beego"
		"strings"
)

type WorkersBootStrap struct {
		base.BaseBootStrapWrapper
		Starters []Starter
}

type WorkersBootstrap interface {
		base.BootStrap
		Start()
		RemoveStarter(int)
		InitStarters([]Starter)
		StartersLen() int
}

// 创建一个worker bootstrap
func NewWorkerBootstrap(name string) WorkersBootstrap {
		if !strings.Contains(name, Prefix()) {
				name = Prefix() + name
		}
		var (
				worker *WorkersBootStrap
				boot   = base.NewBootstrap(name)
		)
		if w, ok := boot.(*WorkersBootStrap); ok {
				return w
		}
		if w, ok := boot.(*base.BaseBootStrapWrapper); ok {
				worker = new(WorkersBootStrap)
				worker.BaseBootStrapWrapper = *w
		}
		return worker
}

func (this *WorkersBootStrap) StartersLen() int {
		return len(this.Starters)
}

func (this *WorkersBootStrap) Block() bool {
		return this.BaseBootStrapDto.Async
}

func (this *WorkersBootStrap) InitStarters(starters []Starter) {
		if this.StartersLen() == 0 {
				this.Starters = starters
				return
		}
		this.Starters = append(this.Starters, starters...)
}

//  启动器
func (this *WorkersBootStrap) Start() {
		if len(this.Starters) == 0 {
				panic("impl worker start")
		}
		for i := 0; i < len(this.Starters); i++ {
				starter := this.Starters[i]
				if starter == nil {
						this.RemoveStarter(i)
						i--
						continue
				}
				starter.Init()
				if starter.AWait() {
						go starter.Start()
						continue
				}
				starter.Start()
		}
}

// 移除stater
func (this *WorkersBootStrap) RemoveStarter(i int) {
		var size = len(this.Starters)
		if size <= i {
				this.Starters = this.Starters[:i]
		}
		if i == 0 {
				this.Starters = this.Starters[i+1:]
		}
		if size > i && i != 0 {
				this.Starters = append(this.Starters[:i], this.Starters[i+1:]...)
		}
}
