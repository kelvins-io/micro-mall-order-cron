package startup

import (
	"gitee.com/cristiane/micro-mall-order-cron/vars"
	"gitee.com/kelvins-io/kelvins/config"
	"log"
)

const (
	SectionEmailConfig              = "email-config"
	OrderPayExpireTaskConfig        = "order-pay-expire-task"
	OrderPayFailedTaskConfig        = "order-apy-failed-task"
	OrderInventoryRestoreTaskConfig = "order-inventory-restore-task"
)

// LoadConfig 加载配置对象映射
func LoadConfig() error {
	// 加载email数据源
	log.Printf("[info] Load default config %s", SectionEmailConfig)
	vars.EmailConfigSetting = new(vars.EmailConfigSettingS)
	config.MapConfig(SectionEmailConfig, vars.EmailConfigSetting)
	// 订单支付过期
	log.Printf("[info] Load default config %s", OrderPayExpireTaskConfig)
	vars.OrderPayExpireTaskSetting = new(vars.OrderPayExpireTaskSettingS)
	config.MapConfig(OrderPayExpireTaskConfig, vars.OrderPayExpireTaskSetting)
	// 订单支付失败
	log.Printf("[info] Load default config %s", OrderPayFailedTaskConfig)
	vars.OrderPayFailedTaskSetting = new(vars.OrderPayFailedTaskSettingS)
	config.MapConfig(OrderPayFailedTaskConfig, vars.OrderPayFailedTaskSetting)
	// 订单恢复
	log.Printf("[info] Load default config %s", OrderInventoryRestoreTaskConfig)
	vars.OrderInventoryRestoreTaskSetting = new(vars.OrderInventoryRestoreTaskSettingS)
	config.MapConfig(OrderInventoryRestoreTaskConfig, vars.OrderInventoryRestoreTaskSetting)
	return nil
}
