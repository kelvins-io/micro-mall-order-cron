package repository

import (
	"gitee.com/cristiane/micro-mall-order-cron/model/mysql"
	"gitee.com/kelvins-io/kelvins"
	"time"
)

func UpdateOrderPayExpire(where interface{}, payExpire time.Time, maps interface{}) (int64, error) {
	return kelvins.XORM_DBEngine.Table(mysql.TableOrder).Where(where).Where("pay_expire <= ?", payExpire).Update(maps)
}

func UpdateOrder(where, maps interface{}) (int64, error) {
	return kelvins.XORM_DBEngine.Table(mysql.TableOrder).Where(where).Update(maps)
}

func FindOrder(sqlSelect string, where interface{}, pageSize, pageNum int) ([]mysql.Order, error) {
	var result = make([]mysql.Order, 0)
	session := kelvins.XORM_DBEngine.Table(mysql.TableOrder).Select(sqlSelect).Where(where).Desc("update_time")
	if pageSize > 0 && pageNum >= 1 {
		session = session.Limit(pageSize, (pageNum-1)*pageSize)
	}
	err := session.Find(&result)
	return result, err
}
