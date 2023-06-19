package main

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func FloatToBigInt(float float64, scale int) *big.Int {
	// Create a big.Float value from the input float64 value
	bigFloat := new(big.Float).SetFloat64(float)

	// Create the scale as a big.Float
	scaleFactor := new(big.Float).SetFloat64(math.Pow10(scale))

	// Multiply the input big.Float value with the scale
	bigFloat.Mul(bigFloat, scaleFactor)

	// Convert big.Float to big.Int
	bigInt := new(big.Int)
	bigFloat.Int(bigInt)

	return bigInt
}

func StringToFloat(str string) float64 {
	float, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(fmt.Sprintf("can't convert string to float - string = %s", str))
	}
	return float
}

func getAddressFromPrivateKey(privateKey *ecdsa.PrivateKey) (common.Address, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, errors.New("unable to get address from private key")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA)
	return address, nil
}
