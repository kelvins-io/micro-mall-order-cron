package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-order-cron/repository"
	"gitee.com/kelvins-io/kelvins"
	"time"
)

func HandleOrderPayExpire() {
	ctx := context.Background()
	maps := map[string]interface{}{
		"pay_state":   4, // 支付过期取消
		"state":       2, // 无效
		"update_time": time.Now(),
	}
	payExpireWhere := time.Now()
	where := map[string]interface{}{
		"state":            0,           // 有效
		"pay_state":        []int{0, 1}, // 未支付,支付中
		"inventory_verify": 0,           // 库存未核实
	}
	rowAffected, err := repository.UpdateOrderPayExpire(where, payExpireWhere, maps)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "UpdateOrderPayExpire err: %v", err)
		return
	}
	_ = rowAffected
}

func HandleOrderPayFailed() {
	ctx := context.Background()
	where := map[string]interface{}{
		"pay_state": 2,
	}
	maps := map[string]interface{}{
		"state":       2,
		"update_time": time.Now(),
	}
	rowAffected, err := repository.UpdateOrder(where, maps)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "UpdateOrder err: %v", err)
		return
	}
	_ = rowAffected

}
