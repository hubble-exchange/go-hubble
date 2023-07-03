package gohubble

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"

	"github.com/hubble-exchange/go-hubble/contracts"
)

var Eip712OrderTypes = apitypes.Types{
	"EIP712Domain": {
		{
			Name: "name",
			Type: "string",
		},
		{
			Name: "version",
			Type: "string",
		},
		{
			Name: "chainId",
			Type: "uint256",
		},
		{
			Name: "verifyingContract",
			Type: "address",
		},
	},
	"Order": {
		{
			Name: "ammIndex",
			Type: "uint256",
		},
		{
			Name: "trader",
			Type: "address",
		},
		{
			Name: "baseAssetQuantity",
			Type: "int256",
		},
		{
			Name: "price",
			Type: "uint256",
		},
		{
			Name: "salt",
			Type: "uint256",
		},
		{
			Name: "reduceOnly",
			Type: "bool",
		},
	},
}

func GetOrderHash(o contracts.IOrderBookOrder) (hash common.Hash) {
	message := map[string]interface{}{
		"ammIndex":          o.AmmIndex.String(),
		"trader":            o.Trader.String(),
		"baseAssetQuantity": o.BaseAssetQuantity.String(),
		"price":             o.Price.String(),
		"salt":              o.Salt.String(),
		"reduceOnly":        o.ReduceOnly,
	}
	domain := apitypes.TypedDataDomain{
		Name:              "Hubble",
		Version:           "2.0",
		ChainId:           math.NewHexOrDecimal256(321123),
		VerifyingContract: common.HexToAddress("0x0300000000000000000000000000000000000000").String(),
	}
	typedData := apitypes.TypedData{
		Types:       Eip712OrderTypes,
		PrimaryType: "Order",
		Domain:      domain,
		Message:     message,
	}
	return EncodeForSigning(typedData)
}

// EncodeForSigning - Encoding the typed data
func EncodeForSigning(typedData apitypes.TypedData) (hash common.Hash) {
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		panic(fmt.Errorf("failed to hash domain separator: %w", err))
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		panic(fmt.Errorf("failed to hash typed data: %w", err))
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	hash = common.BytesToHash(crypto.Keccak256(rawData))
	return
}
