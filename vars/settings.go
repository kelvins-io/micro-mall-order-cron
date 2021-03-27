package vars

type EmailConfigSettingS struct {
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