# micro-mall-order-cron

#### 介绍
订单库定时任务

#### 软件架构
cron

#### 框架，库依赖
kelvins框架支持（gRPC，cron，queue，web支持）：https://gitee.com/kelvins-io/kelvins   
g2cache缓存库支持（两级缓存）：https://gitee.com/kelvins-io/g2cache   

#### 安装教程

1.仅构建  sh build.sh   
2 运行  sh build-run.sh   

#### 使用说明
配置参考
```toml
[kelvins-server]
EndPoint = 8080
IsRecordCallResponse = true

[kelvins-logger]
RootPath = "./logs"
Level = "debug"

[kelvins-mysql]
Host = "127.0.0.1:3306"
UserName = "root"
Password = "xxxx"
DBName = "micro_mall_order"
Charset = "utf8mb4"
PoolNum =  10
MaxIdleConns = 5
ConnMaxLifeSecond = 3600
MultiStatements = true
ParseTime = true

[email-config]
User = "xxxx@qq.com"
Password = "xxx"
Host = "smtp.qq.com"
Port = "465"

[order-pay-expire-task]
Cron = "0 */3 * * * *"

[order-apy-failed-task]
Cron = "20 */8 * * * *"

[order-inventory-restore-task]
Cron = "30 */20 * * * *"
```
#### 参与贡献

1.  Fork 本仓库
2.  新建 Feat_xxx 分支
3.  提交代码
4.  新建 Pull Request

