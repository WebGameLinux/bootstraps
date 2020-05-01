package api

import "sync"

var lock sync.Once

// 初始化一次
func init() {
		lock.Do(workflow)
}

// 初始化流程
func workflow() {
		// 初始化应用配置
		config()
		// 初始化服务提供
		providers()
		// 初始化命令行
		command()
		// 初始化日志
		logger()
		// 初始化中间
		middlewares()
		// 初始化服务路由
		apis()
		// 初始化基础服务
		services()
		// 初始化数据库
		database()
		// worker 最终阻塞
		workers()
}
