package startup

import (
	"gitee.com/cristiane/micro-mall-order-cron/model/args"
	"gitee.com/cristiane/micro-mall-order-cron/vars"
	"gitee.com/kelvins-io/kelvins"
	"gitee.com/kelvins-io/kelvins/setup"
	"gitee.com/kelvins-io/kelvins/util/queue_helper"
)

// SetupVars 加载变量
func SetupVars() error {
	var err error
	err = setupTradeOrderInfoSearchNotice()
	if err != nil {
		return err
	}
	return nil
}

func setupTradeOrderInfoSearchNotice() error {
	var err error
	if vars.TradeOrderInfoSearchNoticeSetting != nil {
		vars.TradeOrderInfoSearchNoticeServer, err = setup.NewAMQPQueue(vars.TradeOrderInfoSearchNoticeSetting, nil)
		if err != nil {
			return err
		}
		vars.TradeOrderInfoSearchNoticePusher, err = queue_helper.NewPublishService(
			vars.TradeOrderInfoSearchNoticeServer, &queue_helper.PushMsgTag{
				DeliveryTag:    args.TradeOrderInfoSearchNoticeTag,
				DeliveryErrTag: args.TradeOrderInfoSearchNoticeTagErr,
				RetryCount:     vars.TradeOrderInfoSearchNoticeSetting.TaskRetryCount,
				RetryTimeout:   vars.TradeOrderInfoSearchNoticeSetting.TaskRetryTimeout,
			}, kelvins.BusinessLogger)
		if err != nil {
			return err
		}
	}
	return err
}
