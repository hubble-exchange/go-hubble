package gohubble

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math"
	"math/big"
	"net/url"
	"os"
	"strconv"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	rpcEndpoint       string
	webSocketEndpoint string
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

func getRPCEndpoint() string {
	if rpcEndpoint == "" {
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
		rpcEndpoint = rpcURL.String()
	}
	return rpcEndpoint
}

func getWebSocketEndpoint() string {
	if webSocketEndpoint == "" {
		rpcHost := os.Getenv("HUBBLE_RPC_HOST")
		if rpcHost == "" {
			panic("HUBBLE_RPC_HOST environment variable not set")
		}
		blockchainID := os.Getenv("HUBBLE_BLOCKCHAIN_ID")
		if rpcHost == "" {
			panic("HUBBLE_BLOCKCHAIN_ID environment variable not set")
		}
		path := fmt.Sprintf("/ext/bc/%s/ws", blockchainID)
		websocketURL := url.URL{Scheme: "wss", Host: rpcHost, Path: path}
		webSocketEndpoint = websocketURL.String()
	}

	return webSocketEndpoint
}
