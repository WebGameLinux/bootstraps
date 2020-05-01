package beego

import (
		"errors"
		"sync"
)

var registerBootStraps sync.Map

type BootStrap interface {
		Name() string                   // 引导器名
		Boot()                          // 执行引导加载注册相关逻辑(包括初始化自身)
		Booted() bool                   // 是否以引导加载
		App() interface{}               // 上线文应用体获取
		Container(string, interface{})  // 设置相关上下文信息
		Get(string) (interface{}, bool) // 	获取相关信息
		Block() bool                    // 引导逻辑是否有阻塞,是否需要异步等待
}

// 数据容器
type BaseBootStrapDto struct {
		Name      string   // 名字
		Booted    bool     // 是否已经加载过
		Container sync.Map // 容器
		Async     bool     // 是否需要等待，阻塞
}

type BaseBootStrapWrapper struct {
		BaseBootStrapDto
		BootHandler      func(dto *BaseBootStrapDto)
		BootedHandler    func(dto *BaseBootStrapDto) bool
		AppHandler       func(dto *BaseBootStrapDto) interface{}
		ContainerHandler func(dto *BaseBootStrapDto, key string, value interface{})
		GetHandler       func(dto *BaseBootStrapDto, key string) (interface{}, bool)
		BlockHandler     func(dto *BaseBootStrapDto) bool
}

func (this *BaseBootStrapWrapper) Name() string {
		return this.BaseBootStrapDto.Name
}

func (this *BaseBootStrapWrapper) Boot() {
		if this.BootHandler != nil {
				this.BootHandler(&this.BaseBootStrapDto)
				return
		}
		panic("implement me")
}

func (this *BaseBootStrapWrapper) Booted() bool {
		if this.BootHandler != nil {
				this.BootHandler(&this.BaseBootStrapDto)
		}
		return this.BaseBootStrapDto.Booted
}

func (this *BaseBootStrapWrapper) App() interface{} {
		v, _ := this.Get("app")
		return v
}

func (this *BaseBootStrapWrapper) Container(key string, v interface{}) {
		if this.ContainerHandler != nil {
				this.ContainerHandler(&this.BaseBootStrapDto, key, v)
		}
		this.BaseBootStrapDto.Container.Store(key, v)
}

func (this *BaseBootStrapWrapper) Get(key string) (interface{}, bool) {
		if this.GetHandler != nil {
				return this.GetHandler(&this.BaseBootStrapDto, key)
		}
		return this.BaseBootStrapDto.Container.Load(key)
}

func (this *BaseBootStrapWrapper) Block() bool {
		if this.BlockHandler != nil {
				return this.BlockHandler(&this.BaseBootStrapDto)
		}
		return this.Async
}

func (this *BaseBootStrapWrapper) InitByDto(dto *BaseBootStrapDto) {
		if dto == nil {
				return
		}
		this.BaseBootStrapDto.Name = dto.Name
		this.BaseBootStrapDto.Booted = dto.Booted
		this.BaseBootStrapDto.Container = dto.Container
		this.BaseBootStrapDto.Async = dto.Async
}

func BootNamed(name string) bool {
		if v, ok := registerBootStraps.Load(name); ok && v != nil {
				return true
		}
		return false
}

func BootNameRegister(name string, v ...interface{}) bool {
		if len(v) == 0 {
				v = append(v, true)
		}
		if _, ok := registerBootStraps.Load(name); !ok {
				registerBootStraps.Store(name, v[0])
		}
		return true
}

func BootNameReset(name string, v ...interface{}) bool {
		if len(v) == 0 {
				v = append(v, nil)
		}
		registerBootStraps.Store(name, v[0])
		return true
}

func GetBootstrapManager() *sync.Map {
		return &registerBootStraps
}

func NewBaseBootstrap(name string) *BaseBootStrapDto {
		if name == "" || BootNamed(name) {
				return nil
		}
		var dto = new(BaseBootStrapDto)
		dto.Name = name
		BootNameRegister(name, dto)
		return dto
}

func NewBootstrap(name string) BootStrap {
		var dto = NewBaseBootstrap(name)
		if dto == nil {
				dtoAny, _ := registerBootStraps.Load(name)
				if dtoAny == nil {
						panic(errors.New("bootstrap " + name + "new struct failed"))
				}
				if v, ok := dtoAny.(*BaseBootStrapDto); ok {
						dto = v
				}
				if v, ok := dtoAny.(*BaseBootStrapWrapper); ok {
						return v
				}
		}
		var wrapper = new(BaseBootStrapWrapper)
		wrapper.InitByDto(dto)
		BootNameReset(name, wrapper)
		return wrapper
}
