package main

import "fmt"

func main() {
	//客户端完全未感受到火车站的存在
	var saleClient SaleTicket = NewTicketProxy()
	fmt.Println(saleClient.SaleTicket())
}

//卖票接口，火车站和代理处 都实现了
type SaleTicket interface {
	SaleTicket() string
}

//火车站
type Station struct {
}

func (receiver *Station) SaleTicket() string {
	return "火车站卖的票"
}

func NewTicketProxy() *TicketProxy {
	proxy := new(TicketProxy)
	proxy.Station = new(Station)
	return proxy
}

//火车站代理
type TicketProxy struct {
	Station SaleTicket
}

//火车站代理
func (receiver *TicketProxy) SaleTicket() string {
	fmt.Println("中介开始做事")
	//火车站代理最终还是通过火车站来购买的票
	ticket := receiver.Station.SaleTicket()
	fmt.Println("中介做事完成")
	return ticket
}
