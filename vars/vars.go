package vars

import (
	"gitee.com/kelvins-io/common/queue"
	"gitee.com/kelvins-io/kelvins/config/setting"
	"gitee.com/kelvins-io/kelvins/util/queue_helper"
)

var (
	EmailConfigSetting                *EmailConfigSettingS
	OrderPayExpireTaskSetting         *OrderPayExpireTaskSettingS
	OrderPayFailedTaskSetting         *OrderPayFailedTaskSettingS
	OrderInventoryRestoreTaskSetting  *OrderInventoryRestoreTaskSettingS
	OrderSearchSyncTaskSetting        *OrderSearchSyncTaskSettingS
	TradeOrderInfoSearchNoticeSetting *setting.QueueAMQPSettingS
	TradeOrderInfoSearchNoticeServer  *queue.MachineryQueue
	TradeOrderInfoSearchNoticePusher  *queue_helper.PublishService
)
