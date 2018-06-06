package main

import (
	"fmt"

	"github.com/blockcdn-go/exchange-sdk-go/global"

	"github.com/blockcdn-go/exchange-sdk-go/weex"
)

func main() {
	c := weex.NewClient(nil)
	s, err := c.GetAllSymbol()
	fmt.Printf("%+v, %+v\n", err, s)

	k, err := c.GetKline(global.KlineReq{
		Base:   s[0].Base,
		Quote:  s[0].Quote,
		Period: "1m",
	})
	fmt.Printf("%+v, %+v\n", err, k)

	tch, err := c.SubTicker(s[0])
	lch, err := c.SubLateTrade(s[0])

	for {
		select {
		case tk := <-tch:
			fmt.Printf("notify %+v\n", tk)
			break
		case lt := <-lch:
			fmt.Printf("notify %+v\n", lt)
			break
		}
	}
}