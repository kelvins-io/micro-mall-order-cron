package startup

import (
	"gitee.com/cristiane/micro-mall-order-cron/service"
	"gitee.com/cristiane/micro-mall-order-cron/vars"
	"gitee.com/kelvins-io/kelvins"
)

func GenCronJobs() []*kelvins.CronJob {
	tasks := make([]*kelvins.CronJob, 0)
	if vars.OrderPayExpireTaskSetting != nil {
		if vars.OrderPayExpireTaskSetting.Cron != "" {
			tasks = append(tasks, &kelvins.CronJob{
				Name: "订单支付超时处理",
				Spec: vars.OrderPayExpireTaskSetting.Cron,
				Job:  service.HandleOrderPayExpire,
			})
		}
	}

	if vars.OrderPayFailedTaskSetting != nil {
		if vars.OrderPayFailedTaskSetting.Cron != "" {
			tasks = append(tasks, &kelvins.CronJob{
				Name: "订单支付失败处理",
				Spec: vars.OrderPayFailedTaskSetting.Cron,
				Job:  service.HandleOrderPayFailed,
			})
		}
	}

	if vars.OrderInventoryRestoreTaskSetting != nil {
		if vars.OrderInventoryRestoreTaskSetting.Cron != "" {
			tasks = append(tasks, &kelvins.CronJob{
				Name: "订单库存恢复",
				Spec: vars.OrderInventoryRestoreTaskSetting.Cron,
				Job:  service.RestoreOrderInventory,
			})
		}
	}

	return tasks
}
