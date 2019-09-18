package wxpay

import "encoding/xml"

type RespBase struct {
	XMLName    xml.Name `xml:"xml"`
	Text       string   `xml:",chardata"`
	ReturnCode string   `xml:"return_code"`
	ReturnMsg  string   `xml:"return_msg"`
}

type RespUnifiedorder struct {
	XMLName    xml.Name `xml:"xml"`
	Text       string   `xml:",chardata"`
	ReturnCode string   `xml:"return_code"`
	ReturnMsg  string   `xml:"return_msg"`
	Appid      string   `xml:"appid"`
	SubAppid   string   `xml:"sub_appid"`
	MchID      string   `xml:"mch_id"`
	SubMchID   string   `xml:"sub_mch_id"`
	NonceStr   string   `xml:"nonce_str"`
	Sign       string   `xml:"sign"`
	ResultCode string   `xml:"result_code"`
	PrepayID   string   `xml:"prepay_id"`
	TradeType  string   `xml:"trade_type"`
}

type UnifiedorderResultNotify struct {
	XMLName       xml.Name `xml:"xml"`
	Text          string   `xml:",chardata"`
	ReturnCode    string   `xml:"return_code"`
	ReturnMsg     string   `xml:"return_msg"`
	Appid         string   `xml:"appid"`
	MchId         string   `xml:"mch_id"`
	TransactionId string   `xml:"transaction_id"`
	OutTradeNo    string   `xml:"out_trade_no"`
	ResultCode    string   `xml:"result_code"`
	ErrCode       string   `xml:"err_code"`
	ErrCodeDes    string   `xml:"err_code_des"`
	TimeEnd       string   `xml:"time_end"`
	TotalFee      int64    `xml:"total_fee"`
	Openid        string   `xml:"openid"`
	Attach        string   `xml:"attach"`
	BankType      string   `xml:"bank_type"`
	FeeType       string   `xml:"fee_type"`
	IsSubscribe   string   `xml:"is_subscribe"`
	NonceStr      string   `xml:"nonce_str"`
	Sign          string   `xml:"sign"`
	CouponFee     []string `xml:"coupon_fee"`
	CouponCount   string   `xml:"coupon_count"`
	CouponType    string   `xml:"coupon_type"`
	CouponId      string   `xml:"coupon_id"`
	TradeType     string   `xml:"trade_type"`
}
