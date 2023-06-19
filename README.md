# Go SDK for Hubble Exchange

Requires HUBBLE_RPC_HOST and HUBBLE_BLOCKCHAIN_ID environment variable to be set

```shell
export HUBBLE_RPC_HOST=candy-hubblenet-rpc.hubble.exchange
export HUBBLE_BLOCKCHAIN_ID=iKMFgo49o4X3Pd3UWUkmPwjKom3xZz3Vo6Y1kkwL2Ce6DZaPm
```

### Example usage:

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

// subscribe to order book updates
respChan, _, err := client.SubscribeToOrderBookDepth(0)
for resp := range respChan {
    // Implement your logic to handle the response
    fmt.Println(resp)
}
```
