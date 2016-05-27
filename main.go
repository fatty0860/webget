package main

import "fmt"

func main() {
	var stkAry = []string{"2330", "2062", "2311", ""}
	var d StockRespose
	var o StockInfoResponse

	for _, text := range stkAry {
		QryStock(text, &d)

		fmt.Printf("msg : %s %s\n", d.RtnCode, d.RtnMessage)
		if d.RtnCode == "0000" && len(d.Info) > 0 {
			fmt.Printf("%s, %s, %s\n", d.Info[0].StkKey, d.Info[0].LowerPrice, d.Info[0].UpperPrice)

			QryStkInfo(d.Info[0].StkKey, &o)
			fmt.Printf("%s\n", o.RtnCode)
		}

	}

	/*
		QryStkInfo("tse_2330.tw", &info)
		if info.RtnCode == "0000" && len(info.Info) > 0 {
			fmt.Printf("%v\n", info.Info[0])
		}
	*/
}

/*
func main() {
	var url string
	var err error
	//var data []byte
	var resp *http.Response
	var token *html.Tokenizer
	var tp html.TokenType
	var tptk html.Token

	url = "http://mis.twse.com.tw/stock/fibest.jsp?stock=2330"
	resp, err = http.Get(url)
	if err != nil {
		fmt.Printf("http get error : %s\n", err.Error())
		return
	}
	defer resp.Body.Close()

	//data, _ = ioutil.ReadAll(resp.Body)
	//fmt.Printf("%v\n", string(data))

	token = html.NewTokenizer(resp.Body)

	for {
		tp = token.Next()
		switch {
		case tp == html.ErrorToken:
			return
		case tp == html.StartTagToken:
			tptk = token.Token()
			if tptk.Data == "a" {
				fmt.Printf("------------\n")
				for _, attr := range tptk.Attr {
					fmt.Printf("%s, %s, %s\n", attr.Namespace, attr.Key, attr.Val)
				}
			}
		}
	}

}
*/
