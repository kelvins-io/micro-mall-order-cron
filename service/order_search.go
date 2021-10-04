package service

import (
	"context"
	"gitee.com/cristiane/micro-mall-order-cron/model/args"
	"gitee.com/cristiane/micro-mall-order-cron/model/mysql"
	"gitee.com/cristiane/micro-mall-order-cron/repository"
	"gitee.com/cristiane/micro-mall-order-cron/vars"
	"gitee.com/kelvins-io/common/json"
	"gitee.com/kelvins-io/kelvins"
	"github.com/google/uuid"
	"strings"
)

var (
	orderSearchSyncPageSize = 50
	orderSearchSyncPageNum  = 1
)

func OrderSearchSync() {
	count := 0
	for {
		if count > 2 {
			break
		}
		count++
		orderSearchSyncOne(orderSearchSyncPageSize, orderSearchSyncPageNum)
		orderSearchSyncPageNum++
	}
}

const sqlSelectOrderSearch = "order_code,description,device_code"

func orderSearchSyncOne(pageSize, pageNum int) {
	ctx := context.TODO()
	where := map[string]interface{}{}
	orderList, err := repository.FindOrder(sqlSelectOrderSearch, where, pageSize, pageNum)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "FindOrder err: %v", err)
		return
	}
	if len(orderList) == 0 {
		return
	}
	orderCodeList := make([]string, len(orderList))
	for i := 0; i < len(orderList); i++ {
		orderCodeList[i] = orderList[i].OrderCode
	}
	orderSkuList, err := repository.FindOrderSku("order_code,name", orderCodeList)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "FindOrderSku err: %v orderCodeList: %v", err, json.MarshalToStringNoError(orderCodeList))
		return
	}
	if len(orderSkuList) == 0 {
		return
	}
	orderCodeToSku := map[string][]mysql.OrderSku{}
	for i := 0; i < len(orderSkuList); i++ {
		orderCodeToSku[orderSkuList[i].OrderCode] = append(orderCodeToSku[orderSkuList[i].OrderCode], orderSkuList[i])
	}
	orderSceneList, err := repository.FindOrderScene("order_code,shop_name,shop_address", orderCodeList)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "FindOrderScene err: %v orderCodeList: %v", err, json.MarshalToStringNoError(orderCodeList))
		return
	}
	orderCodeToScene := map[string]mysql.OrderSceneShop{}
	for i := 0; i < len(orderSceneList); i++ {
		orderCodeToScene[orderSceneList[i].OrderCode] = orderSceneList[i]
	}
	orderInfoList := make([]args.SearchTradeOrderEntry,0)
	for i := 0; i < len(orderList); i++ {
		shopOrder := args.SearchTradeOrderEntry{
			Description: orderList[i].Description,
			DeviceId:    orderList[i].DeviceCode,
			ShopName:    orderCodeToScene[orderList[i].OrderCode].ShopName,
			ShopAddress: orderCodeToScene[orderList[i].OrderCode].ShopAddress,
			OrderCode:   orderList[i].OrderCode,
		}
		goodsName := strings.Builder{}
		for _, v := range orderCodeToSku[orderList[i].OrderCode] {
			goodsName.WriteString(v.Name)
			goodsName.WriteString(";")
		}
		shopOrder.GoodsName = goodsName.String()
		orderInfoList = append(orderInfoList, shopOrder)
	}

	tradeOrderSearchSyncNotice(orderInfoList)
}

// 订单搜索同步
func tradeOrderSearchSyncNotice(orderInfoList []args.SearchTradeOrderEntry) {
	if len(orderInfoList) == 0 {
		return
	}
	for i := 0; i < len(orderInfoList); i++ {
		var ctx = context.TODO()
		var msg = &args.CommonBusinessMsg{
			Type:    args.TradeOrderInfoSearchNoticeType,
			Tag:     "交易订单搜索通知",
			UUID:    uuid.New().String(),
			Content: json.MarshalToStringNoError(orderInfoList[i]),
		}
		vars.TradeOrderInfoSearchNoticePusher.PushMessage(ctx, msg)
	}
}
