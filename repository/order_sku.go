package repository

import (
	"gitee.com/cristiane/micro-mall-order-cron/model/mysql"
	"gitee.com/kelvins-io/kelvins"
)

func GetOrderSkuList(sqlSelect string, shopId []int64, orderCode []string) ([]mysql.OrderSku, error) {
	var result = make([]mysql.OrderSku, 0)
	var err error
	err = kelvins.XORM_DBEngine.Table(mysql.TableOrderSku).Select(sqlSelect).In("order_code", orderCode).In("shop_id", shopId).Find(&result)
	return result, err
}

func FindOrderSku(sqlSelect string, orderCode []string) ([]mysql.OrderSku, error) {
	var result = make([]mysql.OrderSku, 0)
	var err error
	err = kelvins.XORM_DBEngine.Table(mysql.TableOrderSku).Select(sqlSelect).In("order_code", orderCode).Find(&result)
	return result, err
}
