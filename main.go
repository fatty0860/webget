package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	var netClient = &http.Client{
		Timeout: time.Second * 5,
	}

	//rep, err := netClient.Get("http://mis.twse.com.tw/stock/fibest.jsp")
	rep, err := netClient.Get("http://mis.twse.com.tw/stock/")
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	req, _ := http.NewRequest("GET", "http://mis.twse.com.tw/stock/api/getStockInfo.jsp?ex_ch=tse_1101.tw|tse_2330.tw&json=1&delay=0&d=20160526", nil)

	cookies := rep.Cookies()
	for _, c := range cookies {
		//h, m, s := c.Expires.Clock()
		//fmt.Printf("%s = %s  [%d:%d:%d]\n", c.Name, c.Value, h, m, s)
		fmt.Printf("%v\n", c)

		req.AddCookie(c)
	}

	r, err2 := netClient.Do(req)
	if err2 != nil {
		fmt.Printf("%v\n", err2)
		return
	}
	data, _ := ioutil.ReadAll(r.Body)
	fmt.Printf("data : %v\n", string(data))

}
