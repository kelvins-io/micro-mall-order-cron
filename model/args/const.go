package args

const (
	RpcServiceMicroMallSku = "micro-mall-sku"
)

const (
	TradeOrderInfoSearchNoticeType = 10001
)

const (
	TradeOrderInfoSearchNoticeTag    = "trade_order_info_search_notice"
	TradeOrderInfoSearchNoticeTagErr = "trade_order_info_search_notice_err"
)

type SearchTradeOrderEntry struct {
	Description string `json:"description"`
	DeviceId    string `json:"device_id"`
	ShopName    string `json:"shop_name"`
	ShopAddress string `json:"shop_address"`
	GoodsName   string `json:"goods_name"`
	OrderCode   string `json:"order_code"`
}

type CommonBusinessMsg struct {
	Type    int    `json:"type"`
	Tag     string `json:"tag"`
	UUID    string `json:"uuid"`
	Content string `json:"content"`
}

const (
	Unknown = 0
)

var MsgFlags = map[int]string{
	Unknown: "未知",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[Unknown]
}
