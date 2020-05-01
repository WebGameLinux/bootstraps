package worker

import (
		base "github.com/WebGameLinux/bootstraps/beego"
		"github.com/WebGameLinux/bootstraps/beego/exports"
		"regexp"
)

var workerMatcher = regexp.MustCompile("^" + Prefix())

// 自动注册相关逻辑
func Autoloader() {
		var lists = exports.Exports()
		if lists == nil {
				return
		}
		var (
				hookBefore = exports.GetHook(exports.BeforeHookName)
				hookAfter  = exports.GetHook(exports.AfterHookName)
		)
		if hookBefore != nil {
				hookBefore(string(exports.BeforeHookName))
		}
		// 遍历外表配置器中所需要启动的worker
		lists.Range(Register)
		if hookAfter != nil {
				hookAfter(string(exports.AfterHookName))
		}
}

// work 前缀
func Prefix() string {
		return exports.WorkerPrefix
}

// 是否worker bootstrap group
func IsWorkerBootstrapName(name string) bool {
		return workerMatcher.MatchString(name)
}

// worker init and register start
func Register(key, value interface{}) bool {
		if key == nil || value == nil {
				return true
		}
		var (
				ok   bool
				name string
				boot base.BootStrap
		)
		if name, ok = key.(string); !ok {
				return true
		}
		if boot, ok = value.(base.BootStrap); !ok {
				return true
		}
		if boot.Booted() {
				return true
		}
		if boot.Name() != name {
				return true
		}
		if ! IsWorkerBootstrapName(name) {
				return true
		}
		if boot.Block() {
				go Start(boot)
				base.BootNameReset(name, boot)
				return true
		}
		Start(boot)
		base.BootNameReset(name, boot)
		return true
}

// worker 启动逻辑
func Start(boot base.BootStrap) {
		if w, ok := boot.(WorkersBootstrap); ok {
				w.Boot()
				w.Start()
		} else {
				base.Start(boot)
		}
}

// 主进程阻塞服务
func Daemon() {
		handler, ok := exports.Exports().Load(exports.MainProcessName)
		if !ok {
				return
		}
		if fn, ok := handler.(func()); ok {
				fn()
		}
}
