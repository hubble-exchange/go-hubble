package gohubble

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
	c "github.com/hubble-exchange/go-hubble/contracts"
)

// sample
func main() {
	salt, _ := new(big.Int).SetString("299651279690175792460159626333696937226", 10)
	order := c.IOrderBookOrder{
		AmmIndex:          new(big.Int).SetInt64(9),
		Trader:            common.HexToAddress("0x93dAc05dE54C9d5ee5C59F77518F931168FDEC9b"),
		BaseAssetQuantity: new(big.Int).SetInt64(200000000000000000),
		Price:             new(big.Int).SetInt64(363000),
		Salt:              salt,
		ReduceOnly:        false,
	}

	fmt.Printf("order: %+v\n", order)
	hash, err := GetOrderHash(order)
	if err != nil {
		panic(err)
	}
	fmt.Printf("order: %+v\n", hash.String()) // 0x0012b3d78225cc2ffb93b8090a585fcca867976becff8854b797402d7e13bf97
}

func GetOrderHash(o c.IOrderBookOrder) (hash common.Hash, err error) {
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
func EncodeForSigning(typedData apitypes.TypedData) (hash common.Hash, err error) {
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	fmt.Printf("domainSeparator: %+v\n", domainSeparator)
	if err != nil {
		return
	}
	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	fmt.Printf("typedDataHash: %+v\n", typedDataHash)
	if err != nil {
		return
	}
	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	hash = common.BytesToHash(crypto.Keccak256(rawData))
	return
}

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
