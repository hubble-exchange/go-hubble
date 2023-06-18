package hubble

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
	AvgPrice      string   `json:"avgPrice"`
	ClientOrderID string   `json:"clientOrderId"`
	CumQuote      string   `json:"cumQuote"`
	ExecutedQty   string   `json:"executedQty"`
	OrderID       string   `json:"orderId"`
	OrigQty       string   `json:"origQty"`
	OrigType      string   `json:"origType"`
	Price         string   `json:"price"`
	ReduceOnly    bool     `json:"reduceOnly"`
	Side          string   `json:"side"`
	PositionSide  string   `json:"positionSide"`
	Status        string   `json:"status"`
	StopPrice     string   `json:"stopPrice"`
	ClosePosition bool     `json:"closePosition"`
	Symbol        int64    `json:"symbol"`
	Time          int64    `json:"time"`
	TimeInForce   string   `json:"timeInForce"`
	Type          string   `json:"type"`
	ActivatePrice string   `json:"activatePrice"`
	PriceRate     string   `json:"priceRate"`
	UpdateTime    int64    `json:"updateTime"`
	WorkingType   string   `json:"workingType"`
	PriceProtect  bool     `json:"priceProtect"`
	Salt          *big.Int `json:"salt"`
}

type TradingOrderBookDepthResponse struct {
	LastUpdateID int        `json:"lastUpdateId"`
	E            int64      `json:"E"`
	T            int64      `json:"T"`
	Bids         [][]string `json:"bids"`
	Asks         [][]string `json:"asks"`
}

type Market int64

type TraderPosition struct {
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
	Margin         string           `json:"margin"`
	ReservedMargin string           `json:"reservedMargin"`
	Positions      []TraderPosition `json:"positions"`
}
