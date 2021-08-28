package service

import (
	"context"
	"fmt"
	"gitee.com/cristiane/micro-mall-order-cron/model/args"
	"gitee.com/cristiane/micro-mall-order-cron/model/mysql"
	"gitee.com/cristiane/micro-mall-order-cron/pkg/util"
	"gitee.com/cristiane/micro-mall-order-cron/proto/micro_mall_sku_proto/sku_business"
	"gitee.com/cristiane/micro-mall-order-cron/repository"
	"gitee.com/kelvins-io/common/json"
	"gitee.com/kelvins-io/kelvins"
	"strconv"
	"strings"
	"time"
)

const (
	selectInvalidOrder    = "order_code,shop_id"
	selectInvalidOrderSku = "shop_id,amount,sku_code,order_code"
	whereInvalidOrder     = " (pay_state in (2,4) or state = 2) and inventory_verify = 0"
)

func RestoreOrderInventory() {
	ctx := context.Background()
	invalidOrderList, err := repository.FindInvalidOrderList(selectInvalidOrder, whereInvalidOrder, 300, 1)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "FindInvalidOrderList err: %v", err)
		return
	}
	if len(invalidOrderList) == 0 {
		return
	}
	shopIds := make([]int64, 0)
	shopIdsSet := map[int64]struct{}{}
	orderCodes := make([]string, 0)
	orderCodesSet := map[string]struct{}{}
	for i := 0; i < len(invalidOrderList); i++ {
		if _, ok := shopIdsSet[invalidOrderList[i].ShopId]; !ok {
			shopIds = append(shopIds, invalidOrderList[i].ShopId)
			shopIdsSet[invalidOrderList[i].ShopId] = struct{}{}
		}
		if _, ok := orderCodesSet[invalidOrderList[i].OrderCode]; !ok {
			orderCodesSet[invalidOrderList[i].OrderCode] = struct{}{}
			orderCodes = append(orderCodes, invalidOrderList[i].OrderCode)
		}
	}
	invalidOrderSkuList, err := repository.GetOrderSkuList(selectInvalidOrderSku, shopIds, orderCodes)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetOrderSkuList err: %v", err)
		return
	}
	if len(invalidOrderSkuList) == 0 {
		return
	}
	shopIdToOrderSku := map[string][]mysql.OrderSku{}
	for i := 0; i < len(invalidOrderSkuList); i++ {
		row := invalidOrderSkuList[i]
		key := fmt.Sprintf("%d|%s", row.ShopId, row.OrderCode)
		shopIdToOrderSku[key] = append(shopIdToOrderSku[key], invalidOrderSkuList[i])
	}
	inventoryEntryShopList := make([]*sku_business.InventoryEntryShop, 0)
	for k, _ := range shopIdToOrderSku {
		infos := strings.Split(k, "|")
		shopId, err := strconv.ParseInt(infos[0], 10, 64)
		if err != nil {
			kelvins.ErrLogger.Errorf(ctx, "shopId parse err: %v, infos: %v", err, infos)
			return
		}
		inventoryEntryList := make([]*sku_business.InventoryEntryDetail, 0)
		for i := 0; i < len(shopIdToOrderSku[k]); i++ {
			row := shopIdToOrderSku[k][i]
			entry := &sku_business.InventoryEntryDetail{
				SkuCode: row.SkuCode,
				Amount:  int64(row.Amount),
			}
			inventoryEntryList = append(inventoryEntryList, entry)
		}
		inventoryEntryShop := &sku_business.InventoryEntryShop{
			ShopId:     shopId,
			OutTradeNo: infos[1],
			Detail:     inventoryEntryList,
		}
		inventoryEntryShopList = append(inventoryEntryShopList, inventoryEntryShop)
	}
	if len(inventoryEntryShopList) == 0 {
		return
	}
	serverName := args.RpcServiceMicroMallSku
	conn, err := util.GetGrpcClient(ctx, serverName)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return
	}
	//defer conn.Close()
	skuClient := sku_business.NewSkuBusinessServiceClient(conn)
	restoreReq := sku_business.RestoreInventoryRequest{
		List: inventoryEntryShopList,
		OperationMeta: &sku_business.OperationMeta{
			OpUid: 0,
			OpIp:  "micro_mall_order_cron",
		},
	}
	restoreRsp, err := skuClient.RestoreInventory(ctx, &restoreReq)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "DeductInventory %v,err: %v", serverName, err)
		return
	}
	if restoreRsp.Common.Code != sku_business.RetCode_SUCCESS || !restoreRsp.IsSuccess {
		kelvins.ErrLogger.Infof(ctx, "RestoreInventory req: %v, req: %v", json.MarshalToStringNoError(restoreReq), json.MarshalToStringNoError(restoreRsp))
		return
	}
	if len(orderCodes) == 0 {
		return
	}
	// 记录扣减，防止重复扣减
	where := map[string]interface{}{
		"order_code":       orderCodes, // 订单code
		"inventory_verify": 0,          // 未核实
	}
	maps := map[string]interface{}{
		"inventory_verify": 1, // 已核实
		"update_time":      time.Now(),
	}
	rowAffected, err := repository.UpdateOrder(where, maps)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "UpdateOrder err: %v", err)
		return
	}
	_ = rowAffected
}
