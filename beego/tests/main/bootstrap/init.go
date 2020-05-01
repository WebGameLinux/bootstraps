package bootstrap

import "github.com/WebGameLinux/bootstraps/beego/exports"

func init() {
		exports.Exports().Store(BeeGoWorkBootStrapName(), NewBeeGoWebWorkerBootstrap())
}
