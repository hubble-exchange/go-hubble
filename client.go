package main

import (
	"bytes"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/websocket"
	"github.com/hubble-exchange/go-hubble/contracts"
)

const CHAIN_ID = 321123

var hubbleClient *HubbleClient

type HubbleClient struct {
	rpcEndpoint       string
	websocketEndpoint string
	privateKey        *ecdsa.PrivateKey
	publicAddress     common.Address
}

func GetHubbleClient() *HubbleClient {
	if hubbleClient == nil {
		rpcHost := os.Getenv("HUBBLE_RPC_HOST")
		if rpcHost == "" {
			panic("HUBBLE_RPC_HOST environment variable not set")
		}
		blockchainID := os.Getenv("HUBBLE_BLOCKCHAIN_ID")
		if rpcHost == "" {
			panic("HUBBLE_BLOCKCHAIN_ID environment variable not set")
		}
		path := fmt.Sprintf("/ext/bc/%s/rpc", blockchainID)
		rpcURL := url.URL{Scheme: "https", Host: rpcHost, Path: path}

		path = fmt.Sprintf("/ext/bc/%s/ws", blockchainID)
		websocketURL := url.URL{Scheme: "wss", Host: rpcHost, Path: path}
		hubbleClient = &HubbleClient{
			rpcEndpoint:       rpcURL.String(),
			websocketEndpoint: websocketURL.String(),
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

func (c *HubbleClient) GetOrderBook(market Market) (*OrderBookDepthResponse, error) {
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

	orderBookResponse := OrderBookDepthResponse{}
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

func (c *HubbleClient) CancelOrderById(orderId common.Hash) error {

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

	return c.CancelOrders([]Order{order})
}

func (client *HubbleClient) SubscribeToOrderBookDepth(market int) (chan OrderBookDepthUpdateResponse, *websocket.Conn, error) {
	c, _, err := websocket.DefaultDialer.Dial(client.websocketEndpoint, nil)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to connect to websocket endpoint - err: %s", err)
	}

	msg := Message{
		Jsonrpc: "2.0",
		ID:      1,
		Method:  "trading_subscribe",
		Params:  []interface{}{"streamDepthUpdateForMarket", market},
	}

	err = c.WriteJSON(msg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to write message to websocket - err: %s", err)
	}

	respChan := make(chan OrderBookDepthUpdateResponse)

	go func() {
		defer close(respChan)
		defer c.Close()

		for {
			var resp WebsocketResponse

			err := c.ReadJSON(&resp)
			if err != nil {
				log.Fatalf("error reading websocket message: %v", err)
				return
			}

			respChan <- resp.Params.Result
		}
	}()

	return respChan, c, nil
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
