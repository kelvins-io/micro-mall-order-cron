package repository

import (
	"gitee.com/cristiane/micro-mall-order-cron/model/mysql"
	"gitee.com/kelvins-io/kelvins"
	"time"
	"xorm.io/xorm"
)

func UpdateOrderPayExpire(where interface{}, payExpire time.Time, maps interface{}) (int64, error) {
	return kelvins.XORM_DBEngine.Table(mysql.TableOrder).Where(where).Where("pay_expire <= ?", payExpire).Update(maps)
}

func UpdateOrder(where, maps interface{}) (int64, error) {
	return kelvins.XORM_DBEngine.Table(mysql.TableOrder).Where(where).Update(maps)
}

func UpdateOrderByTx(tx *xorm.Session, where, maps interface{}) (int64, error) {
	return tx.Table(mysql.TableOrder).Where(where).Update(maps)
}

func FindInvalidOrderList(sqlSelect string, where interface{}, pageSize, pageNum int) ([]mysql.Order, error) {
	var result = make([]mysql.Order, 0)
	session := kelvins.XORM_DBEngine.Table(mysql.TableOrder).Select(sqlSelect).Where(where)
	if pageSize > 0 && pageNum > 0 {
		session = session.Limit(pageSize, (pageNum-1)*pageSize)
	}
	err := session.Find(&result)
	return result, err
}
