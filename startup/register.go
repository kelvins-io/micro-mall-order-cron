package startup

import (
	"gitee.com/cristiane/micro-mall-order-cron/service"
	"gitee.com/kelvins-io/kelvins"
)

const (
	CronHandleOrderPayExpireTask    = "30 */5 * * * *"
	CronHandleOrderPayFailedTask    = "30 */5 * * * *"
	CronHandleOrderInventoryRestore = "30 */5 * * * *"
)

func GenCronJobs() []*kelvins.CronJob {
	tasks := make([]*kelvins.CronJob, 0)
	tasks = append(tasks, &kelvins.CronJob{
		Name: "订单支付超时处理",
		Spec: CronHandleOrderPayExpireTask,
		Job:  service.HandleOrderPayExpire,
	})
	tasks = append(tasks, &kelvins.CronJob{
		Name: "订单支付失败处理",
		Spec: CronHandleOrderPayFailedTask,
		Job:  service.HandleOrderPayFailed,
	})
	tasks = append(tasks, &kelvins.CronJob{
		Name: "订单库存恢复",
		Spec: CronHandleOrderInventoryRestore,
		Job:  service.RestoreOrderInventory,
	})
	return tasks
}
