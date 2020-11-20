package service

import (
	"context"
	"fmt"
	"gitee.com/cristiane/micro-mall-order-cron/model/args"
	"gitee.com/cristiane/micro-mall-order-cron/model/mysql"
	"gitee.com/cristiane/micro-mall-order-cron/pkg/util"
	"gitee.com/cristiane/micro-mall-order-cron/proto/micro_mall_sku_proto/sku_business"
	"gitee.com/cristiane/micro-mall-order-cron/repository"
	"gitee.com/kelvins-io/kelvins"
	"strconv"
	"strings"
)

const (
	selectInvalidOrder    = "order_code,shop_id"
	selectInvalidOrderSku = "amount,sku_code,order_code"
	whereInvalidOrder     = " (pay_state in (2,4) or state = 2) and inventory_state = 0"
)

func RestoreOrderInventory() {
	ctx := context.Background()
	kelvins.BusinessLogger.Infof(ctx, "RestoreOrderInventory start")
	invalidOrderList, err := repository.FindInvalidOrderList(selectInvalidOrder, whereInvalidOrder)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "FindInvalidOrderList err: %v", err)
		return
	}
	if len(invalidOrderList) == 0 {
		return
	}
	shopIds := make([]int64, 0)
	orderCodes := make([]string, 0)
	for i := 0; i < len(invalidOrderList); i++ {
		shopIds = append(shopIds, invalidOrderList[i].ShopId)
		orderCodes = append(orderCodes, invalidOrderList[i].OrderCode)
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
	serverName := args.RpcServiceMicroMallSku
	conn, err := util.GetGrpcClient(serverName)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "GetGrpcClient %v,err: %v", serverName, err)
		return
	}
	defer conn.Close()
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
		kelvins.ErrLogger.Infof(ctx, "RestoreInventory not ok ,rsp: %+v, req: %+v", restoreRsp, restoreReq)
		return
	}
	tx := kelvins.XORM_DBEngine.NewSession()
	err = tx.Begin()
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "UpdateOrder Begin err: %v", err)
		return
	}
	// 记录扣减，防止重复扣减
	where := map[string]interface{}{
		"order_code":      orderCodes, // 订单code
		"inventory_state": 0,          // 未核实
	}
	maps := map[string]interface{}{
		"inventory_state": 1, // 已核实
	}
	rowAffected, err := repository.UpdateOrderByTx(tx, where, maps)
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "UpdateOrder err: %v", err)
		return
	}
	_ = rowAffected
	err = tx.Commit()
	if err != nil {
		kelvins.ErrLogger.Errorf(ctx, "UpdateOrder Commit err: %v", err)
		return
	}
}
