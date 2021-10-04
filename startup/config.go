package startup

import (
	"gitee.com/cristiane/micro-mall-order-cron/vars"
	"gitee.com/kelvins-io/kelvins/config"
	"gitee.com/kelvins-io/kelvins/config/setting"
)

const (
	SectionEmailConfig                = "email-config"
	OrderPayExpireTaskConfig          = "order-pay-expire-task"
	OrderPayFailedTaskConfig          = "order-apy-failed-task"
	OrderInventoryRestoreTaskConfig   = "order-inventory-restore-task"
	OrderSearchSyncTaskConfig         = "order-search-sync-task"
	SectionTradeOrderInfoSearchNotice = "trade-order-info-search-notice"
)

// LoadConfig 加载配置对象映射
func LoadConfig() error {
	// 加载email数据源
	vars.EmailConfigSetting = new(vars.EmailConfigSettingS)
	config.MapConfig(SectionEmailConfig, vars.EmailConfigSetting)
	// 订单支付过期
	vars.OrderPayExpireTaskSetting = new(vars.OrderPayExpireTaskSettingS)
	config.MapConfig(OrderPayExpireTaskConfig, vars.OrderPayExpireTaskSetting)
	// 订单支付失败
	vars.OrderPayFailedTaskSetting = new(vars.OrderPayFailedTaskSettingS)
	config.MapConfig(OrderPayFailedTaskConfig, vars.OrderPayFailedTaskSetting)
	// 订单恢复
	vars.OrderInventoryRestoreTaskSetting = new(vars.OrderInventoryRestoreTaskSettingS)
	config.MapConfig(OrderInventoryRestoreTaskConfig, vars.OrderInventoryRestoreTaskSetting)
	// 订单搜索同步
	vars.OrderSearchSyncTaskSetting = new(vars.OrderSearchSyncTaskSettingS)
	config.MapConfig(OrderSearchSyncTaskConfig, vars.OrderSearchSyncTaskSetting)

	// 订单搜索通知
	vars.TradeOrderInfoSearchNoticeSetting = new(setting.QueueAMQPSettingS)
	config.MapConfig(SectionTradeOrderInfoSearchNotice, vars.TradeOrderInfoSearchNoticeSetting)

	return nil
}
