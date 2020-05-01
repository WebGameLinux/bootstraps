package bootstrap

import "github.com/WebGameLinux/bootstraps/beego/worker/exports"

func init() {
		exports.Exports().Store(BeeGoWorkBootStrapName(), NewBeeGoWebWorkerBootstrap())
}
