package gohubble

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/hubble-exchange/go-hubble/contracts"
)

var OrderBookContractAddress = common.HexToAddress("0x0300000000000000000000000000000000000000")

func placeOrder(txOptions *bind.TransactOpts, order contracts.IOrderBookOrder) (common.Hash, error) {
	client, err := ethclient.Dial(getRPCEndpoint())
	if err != nil {
		return common.Hash{}, err
	}

	instance, err := contracts.NewOrderBook(OrderBookContractAddress, client)
	if err != nil {
		return common.Hash{}, err
	}

	orderHash := GetOrderHash(order)
	_, err = instance.PlaceOrder(txOptions, order)
	if err != nil {
		return orderHash, err
	}

	return orderHash, nil
}

func cancelOrder(txOptions *bind.TransactOpts, orders []contracts.IOrderBookOrder) error {
	client, err := ethclient.Dial(getRPCEndpoint())
	if err != nil {
		return err
	}

	instance, err := contracts.NewOrderBook(OrderBookContractAddress, client)
	if err != nil {
		return err
	}

	_, err = instance.CancelOrders(txOptions, orders)
	if err != nil {
		return err
	}

	return nil
}
