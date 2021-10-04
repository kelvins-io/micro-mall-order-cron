package repository

import (
	"gitee.com/cristiane/micro-mall-order-cron/model/mysql"
	"gitee.com/kelvins-io/kelvins"
)

func FindOrderScene(sqlSelect string, orderCode []string) ([]mysql.OrderSceneShop, error) {
	var result = make([]mysql.OrderSceneShop, 0)
	var err error
	err = kelvins.XORM_DBEngine.Table(mysql.TableOrderSceneShop).Select(sqlSelect).In("order_code", orderCode).Find(&result)
	return result, err
}
