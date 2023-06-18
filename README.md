# go-hubble
Go SDK for Hubble Exchange

Example usage:

```go
import hubble "github.com/hubble-exchange/go-hubble"

// Initialise client
client := hubble.GetHubbleClient()
client = client.WithTrader("59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d")

// get order book depth
resp, err := client.GetOrderBook(0) // for market = 0

// place order
order, err := client.PlaceOrder(0 /* market */, 0.1 /* quantity */, 1800 /* price */, false /* reduceOnly */)

// get order status
resp, err := client.GetOrderStatus(order.Id)

// get positions and margin
resp, err := client.GetMarginAndPositions()

// cancel order
err = client.CancelOrders([]Order{order})

// cancel order by id
err = client.CancelOrderById(order.Id)
```
