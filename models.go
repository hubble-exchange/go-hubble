package gohubble

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type APIResponse struct {
	JsonRpc string      `json:"jsonrpc"`
	ID      int         `json:"id"`
	Result  interface{} `json:"result"`
}
type Order struct {
	Id                common.Hash    `json:"id"`
	AmmIndex          *big.Int       `json:"ammIndex"`
	Trader            common.Address `json:"trader"`
	BaseAssetQuantity *big.Int       `json:"baseAssetQuantity"`
	Price             *big.Int       `json:"price"`
	Salt              *big.Int       `json:"salt"`
	ReduceOnly        bool           `json:"reduceOnly"`
}

type OrderStatusResponse struct {
	ExecutedQty  string   `json:"executedQty"`
	OrderID      string   `json:"orderId"`
	OrigQty      string   `json:"origQty"`
	Price        string   `json:"price"`
	ReduceOnly   bool     `json:"reduceOnly"`
	PositionSide string   `json:"positionSide"`
	Status       string   `json:"status"`
	Symbol       int64    `json:"symbol"`
	Time         int64    `json:"time"`
	Type         string   `json:"type"`
	UpdateTime   int64    `json:"updateTime"`
	Salt         *big.Int `json:"salt"`
}

type OrderBookDepthResponse struct {
	LastUpdateID int        `json:"lastUpdateId"`
	E            int64      `json:"E"`
	T            int64      `json:"T"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

type Market int64

type Position struct {
	Market               Market `json:"market"`
	OpenNotional         string `json:"openNotional"`
	Size                 string `json:"size"`
	UnrealisedFunding    string `json:"unrealisedFunding"`
	LiquidationThreshold string `json:"liquidationThreshold"`
	NotionalPosition     string `json:"notionalPosition"`
	UnrealisedProfit     string `json:"unrealisedProfit"`
	MarginFraction       string `json:"marginFraction"`
	LiquidationPrice     string `json:"liquidationPrice"`
	MarkPrice            string `json:"markPrice"`
}

type GetPositionsResponse struct {
	Margin         string     `json:"margin"`
	ReservedMargin string     `json:"reservedMargin"`
	Positions      []Position `json:"positions"`
}

type Message struct {
	Jsonrpc string        `json:"jsonrpc"`
	ID      int           `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

type WebsocketResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  struct {
		Subscription string                       `json:"subscription"`
		Result       OrderBookDepthUpdateResponse `json:"result"`
	}
}

type OrderBookDepthUpdateResponse struct {
	T      int64      `json:"T"`
	Symbol int64      `json:"s"`
	Bids   [][]string `json:"b"`
	Asks   [][]string `json:"a"`
}
