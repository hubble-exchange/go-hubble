package hubble

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/hubble-exchange/go-hubble/contracts"
)

const CHAIN_ID = 321123

var hubbleClient *HubbleClient

type HubbleClient struct {
	rpcEndpoint   string
	privateKey    *ecdsa.PrivateKey
	publicAddress common.Address
}

func GetHubbleClient() *HubbleClient {
	if hubbleClient == nil {
		rpcEndpoint := os.Getenv("RPC_ENDPOINT")
		if rpcEndpoint == "" {
			panic("RPC_ENDPOINT environment variable not set")
		}

		hubbleClient = &HubbleClient{
			rpcEndpoint: rpcEndpoint,
		}

	}
	return hubbleClient
}

func (c *HubbleClient) WithTrader(privateKey string) *HubbleClient {
	privateKeyECDSA, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		panic("Invalid private key")
	}

	publicAddress, err := getAddressFromPrivateKey(privateKeyECDSA)
	if err != nil {
		panic("can't derive public address from private key")
	}

	return &HubbleClient{
		rpcEndpoint:   c.rpcEndpoint,
		privateKey:    privateKeyECDSA,
		publicAddress: publicAddress,
	}
}

func (c *HubbleClient) GetOrderBook(market Market) (*TradingOrderBookDepthResponse, error) {
	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "trading_getTradingOrderBookDepth",
		"params":  []interface{}{int64(market)},
	}

	apiResponse, err := c.makeRPCRequest(requestBody)
	if err != nil {
		return nil, err
	}

	orderBookResponse := TradingOrderBookDepthResponse{}
	err = decodeResponse(apiResponse, &orderBookResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response - err: %s", err)
	}
	return &orderBookResponse, nil
}

func (c *HubbleClient) GetMarginAndPositions() (*GetPositionsResponse, error) {
	if c.privateKey == nil {
		return nil, fmt.Errorf("private key is not set")
	}

	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "trading_getMarginAndPositions",
		"params":  []interface{}{c.publicAddress.String()},
	}

	apiResponse, err := c.makeRPCRequest(requestBody)
	if err != nil {
		return nil, err
	}

	getPositionsResponse := GetPositionsResponse{}
	decodeResponse(apiResponse, &getPositionsResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response - err: %s", err)
	}

	return &getPositionsResponse, nil
}

func (c *HubbleClient) GetOrderStatus(orderID common.Hash) (*OrderStatusResponse, error) {
	requestBody := map[string]interface{}{
		"jsonrpc": "2.0",
		"id":      1,
		"method":  "trading_getOrderStatus",
		"params":  []interface{}{orderID},
	}

	apiResponse, err := c.makeRPCRequest(requestBody)
	if err != nil {
		return nil, err
	}

	orderStatusResponse := OrderStatusResponse{}
	decodeResponse(apiResponse, &orderStatusResponse)
	if err != nil {
		return nil, fmt.Errorf("failed to decode response - err: %s", err)
	}

	return &orderStatusResponse, nil
}

func (c *HubbleClient) PlaceOrder(market int64, baseAssetQuantity float64, price float64, reduceOnly bool) (Order, error) {
	if c.privateKey == nil {
		return Order{}, fmt.Errorf("private key is not set")
	}

	saltStr := strconv.Itoa(int(time.Now().Unix())) + strconv.Itoa(int(rand.Int63n(1000)))
	salt, _ := big.NewInt(0).SetString(saltStr, 10)

	order := contracts.IOrderBookOrder{
		AmmIndex:          big.NewInt(market),
		Trader:            c.publicAddress,
		BaseAssetQuantity: FloatToBigInt(baseAssetQuantity, 18),
		Price:             FloatToBigInt(price, 6),
		Salt:              salt,
		ReduceOnly:        reduceOnly,
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(c.privateKey, big.NewInt(CHAIN_ID))
	if err != nil {
		return Order{}, err
	}

	orderId, err := placeOrder(txOpts, order)

	order_ := Order{
		Id:                orderId,
		AmmIndex:          order.AmmIndex,
		Trader:            order.Trader,
		BaseAssetQuantity: order.BaseAssetQuantity,
		Price:             order.Price,
		Salt:              order.Salt,
		ReduceOnly:        order.ReduceOnly,
	}
	return order_, err
}

func (c *HubbleClient) CancelOrders(orders []Order) error {

	if c.privateKey == nil {
		return fmt.Errorf("private key is not set")
	}

	contractOrders := []contracts.IOrderBookOrder{}
	for _, order := range orders {
		order_ := contracts.IOrderBookOrder{
			AmmIndex:          order.AmmIndex,
			Trader:            order.Trader,
			BaseAssetQuantity: order.BaseAssetQuantity,
			Price:             order.Price,
			Salt:              order.Salt,
			ReduceOnly:        order.ReduceOnly,
		}
		contractOrders = append(contractOrders, order_)

	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(c.privateKey, big.NewInt(CHAIN_ID))
	if err != nil {
		return err
	}

	return cancelOrder(txOpts, contractOrders)
}

func (c *HubbleClient) CancelOrdersById(orderId common.Hash) error {

	if c.privateKey == nil {
		return fmt.Errorf("private key is not set")
	}

	orderDetails, err := c.GetOrderStatus(orderId)
	if err != nil {
		return fmt.Errorf("failed to get order status - err: %s", err)
	}

	order := Order{
		AmmIndex:          big.NewInt(orderDetails.Symbol),
		Trader:            c.publicAddress,
		BaseAssetQuantity: FloatToBigInt(StringToFloat(orderDetails.OrigQty), 18),
		Price:             FloatToBigInt(StringToFloat(orderDetails.Price), 6),
		Salt:              orderDetails.Salt,
		ReduceOnly:        orderDetails.ReduceOnly,
	}
	fmt.Println("order=", order)

	return c.CancelOrders([]Order{order})
}

func (c *HubbleClient) makeRPCRequest(requestBody map[string]interface{}) (*APIResponse, error) {
	jsonBody, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.Post(c.rpcEndpoint, "application/json", bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("RPC request failed with status code: %d", resp.StatusCode)
	}

	var apiResponse APIResponse
	err = json.Unmarshal(responseBody, &apiResponse)

	return &apiResponse, err
}

func decodeResponse(apiResponse *APIResponse, responseStruct interface{}) error {
	bytes, err := json.Marshal(apiResponse.Result)
	if err != nil {
		return fmt.Errorf("failed to marshal response: %s", err)
	}
	err = json.Unmarshal(bytes, &responseStruct)
	if err != nil {
		return fmt.Errorf("failed to unmarshal response: %s", err)
	}
	return nil
}
