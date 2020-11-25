package main

import (
	"gitee.com/cristiane/micro-mall-order-cron/startup"
	"gitee.com/cristiane/micro-mall-order-cron/vars"
	"gitee.com/kelvins-io/kelvins"
	"gitee.com/kelvins-io/kelvins/app"
)

const APP_NAME = "micro-mall-order-cron"

func main() {
	vars.AppName = APP_NAME
	application := &kelvins.CronApplication{
		Application: &kelvins.Application{
			LoadConfig: startup.LoadConfig,
			SetupVars:  startup.SetupVars,
			Name:       APP_NAME,
		},
		GenCronJobs: startup.GenCronJobs,
	}
	app.RunCronApplication(application)
}
