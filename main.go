package hubble

import (
	"fmt"
	"time"
)

func main() {
	client := GetHubbleClient()
	client = client.WithTrader("59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d")

	order, err := client.WithTrader("59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d").PlaceOrder(0, 0.1, 1800, false)
	fmt.Println("err=", err)
	fmt.Println("order=", order)

	// resp, err := client.GetOrderBook(0)
	// fmt.Println("err=", err)
	// fmt.Println("resp=", resp)

	// resp, err := client.WithTrader("59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d").GetMarginAndPositions()
	// fmt.Println("err=", err)
	// fmt.Println("resp=", resp)

	// time.Sleep(time.Second * 5)
	// resp, err := client.GetOrderStatus(order.Id)
	// fmt.Println("err=", err)
	// fmt.Println("resp=", resp)

	// time.Sleep(time.Second * 5)
	// err = client.CancelOrders([]Order{order})
	// fmt.Println("err=", err)

	time.Sleep(time.Second * 5)
	err = client.CancelOrdersById(order.Id)
	fmt.Println("err=", err)

	// fmt.Println(common.BytesToHash(hash).String())
}
