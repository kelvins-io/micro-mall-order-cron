package repository

import (
	"gitee.com/cristiane/micro-mall-order-cron/model/mysql"
	"gitee.com/kelvins-io/kelvins"
	"time"
	"xorm.io/xorm"
)

func UpdateOrderPayExpire(payExpire time.Time, maps interface{}) (int64, error) {
	return kelvins.XORM_DBEngine.Table(mysql.TableOrder).Where("pay_expire <= ?", payExpire).Update(maps)
}

func UpdateOrder(where, maps interface{}) (int64, error) {
	return kelvins.XORM_DBEngine.Table(mysql.TableOrder).Where(where).Update(maps)
}

func UpdateOrderByTx(tx *xorm.Session, where, maps interface{}) (int64, error) {
	return tx.Table(mysql.TableOrder).Where(where).Update(maps)
}

func FindInvalidOrderList(sqlSelect string, where interface{}) ([]mysql.Order, error) {
	var result = make([]mysql.Order, 0)
	err := kelvins.XORM_DBEngine.Table(mysql.TableOrder).Select(sqlSelect).Where(where).Find(&result)
	return result, err
}
