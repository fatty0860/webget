package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	//TwseAPIBase : Twse api url
	TwseAPIBase = "http://mis.twse.com.tw/stock/api"
)

//Stock : TWSE 商品基本資料結構
type Stock struct {
	Exchange   string `json:"ex"`
	TradeDate  string `json:"d"`
	It         string `json:"it"`
	Name       string `json:"\n"`
	I          string `json:"i"`
	IP         string `json:"ip"`
	LowerPrice string `json:"w"`
	UpperPrice string `json:"u"`
	T          string `json:"t"`
	P          string `json:"p"`
	StkCode    string `json:"ch"`
	StkKey     string `json:"key"`
	Y          string `json:"y"`
	StopEnd    string `json:"rt"`
	StopBegin  string `json:"st"`
}

//StockRespose : TWSE 商品基本資料擋查詢結果
type StockRespose struct {
	Info       []Stock `json:"msgArray"`
	RtnMessage string  `json:"rtmessage"`
	RtnCode    string  `json:"rtcode"`
}

//StockInfo : TWSE 商品報價資料結構
type StockInfo struct {
	Ts           string `json:"ts"`
	Fv           string `json:"fv"`
	Tk0          string `json:"tk0"`
	Tk1          string `json:"tk1"`
	Oa           string `json:"oa"`
	Ob           string `json:"ob"`
	Tlong        int64  `json:"tlong"`
	Ot           string `json:"ot"`
	Best5AskQty  string `json:"f"`
	Exchange     string `json:"ex"`
	Best5BidQty  string `json:"g"`
	Ov           string `json:"ov"`
	TradeDate    string `json:"d"`
	It           string `json:"it"`
	Best5BidPx   string `json:"b"`
	Symbol       string `json:"c"`
	Mt           string `json:"mt"`
	Best5AskPx   string `json:"a"`
	Name         string `json:"n"`
	OpenPx       string `json:"o"`
	DayLowerPx   string `json:"l"`
	Oz           string `json:"oz"`
	DayUpperPx   string `json:"h"`
	IP           string `json:"ip"`
	I            string `json:"i"`
	LowerLimitPx string `json:"w"`
	TotalVolumn  string `json:"v"`
	UpperLimitPx string `json:"u"`
	DateTime     string `json:"t"`
	S            string `json:"s"`
	Pz           string `json:"pz"`
	Volumn       string `json:"tv"`
	P            string `json:"p"`
	FullName     string `json:"nf"`
	Channel      string `json:"ch"`
	MatchPx      string `json:"z"`
	LastDayPx    string `json:"y"`
	Ps           string `json:"ps"`
}

// StockInfoResponse :
type StockInfoResponse struct {
	Info       []StockInfo `json:"msgArray"`
	RtnMessage string      `json:"rtmessage"`
	RtnCode    string      `json:"rtcode"`
	UserDelay  string      `json:"userDelay"`
}

// QryStock : 查詢 商品資料
// http://mis.twse.com.tw/stock/api/getStock.jsp?ch=2330.tw&json=1
func QryStock(sym string, rep *StockRespose) bool {
	var url string
	url = fmt.Sprintf("%s/getStock.jsp?ch=%s.tw&json=1", TwseAPIBase, sym)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("http get error : %s\n", err.Error())
		return false
	}
	defer resp.Body.Close()

	data, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal([]byte(data), rep)

	return true
}

//QryStkInfo : 查詢商品報價資料
// http://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=tse_2330.tw&json=1&delay=0
func QryStkInfo(sym string, rep *StockInfoResponse) bool {
	var url string
	var data []byte
	//var err error
	//var resp *http.Response

	url = fmt.Sprintf("%s/getStockInfo.jsp?ex_ch=%s&json=1", TwseAPIBase, sym)
	fmt.Printf("%s\n", url)

	resp, err := http.Get(url)

	fmt.Printf("t2\n")
	if err != nil {
		fmt.Printf("http get error : %s\n", err.Error())
		return false
	}
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	json.Unmarshal(data, rep)

	return true
}
