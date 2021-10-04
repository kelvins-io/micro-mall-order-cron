package vars

type EmailConfigSettingS struct {
	Enable   bool   `json:"enable"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

type OrderPayExpireTaskSettingS struct {
	Cron string `json:"cron"`
}

type OrderPayFailedTaskSettingS struct {
	Cron string `json:"cron"`
}

type OrderInventoryRestoreTaskSettingS struct {
	Cron string `json:"cron"`
}

type OrderSearchSyncTaskSettingS struct {
	Cron          string `json:"cron"`
	SingleSyncNum int    `json:"single_sync_num"`
}
