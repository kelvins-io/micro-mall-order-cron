package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-order-cron/repository"
	"gitee.com/kelvins-io/kelvins"
	"time"
)

func HandleOrderPayExpire() {
	ctx := context.Background()
	kelvins.BusinessLogger.Infof(ctx, "HandleOrderPayExpire start")
	maps := map[string]interface{}{
		"pay_state": 4, // 支付过期取消
		"state":     2,
	}
	payExpire := time.Now()
	rowAffected, err := repository.UpdateOrderPayExpire(payExpire, maps)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "UpdateOrderPayExpire err: %v", err)
		return
	}
	_ = rowAffected
}

func HandleOrderPayFailed() {
	ctx := context.Background()
	kelvins.BusinessLogger.Infof(ctx, "HandleOrderPayFailed start")
	where := map[string]interface{}{
		"pay_state": 2,
	}
	maps := map[string]interface{}{
		"state": 2,
	}
	rowAffected, err := repository.UpdateOrder(where, maps)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "UpdateOrder err: %v", err)
		return
	}
	_ = rowAffected

}
