package exports

import "sync"

var (
		lock      sync.Once
		container *sync.Map
)

type HookName string

type BootStrapHook func(ty string)

const (
		WorkerPrefix            = "worker-"
		BeforeHookName HookName = "workerBefore"
		AfterHookName  HookName = "workerAfter"
)

func init() {
		lock.Do(func() {
				if container == nil {
						container = new(sync.Map)
				}
		})
}

func Exports() *sync.Map {
		return container
}

func GetHook(name HookName) BootStrapHook {
		if v, ok := container.Load(name); ok {
				if fn, ok := v.(BootStrapHook); ok {
						return fn
				}
		}
		return nil
}
