package bridge

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/mobile"
	"math/big"
)

const (
	balanceOf = "balanceOf(address)"
	/*
	 * crypto.Keccak256([]byte(balanceOf)) + [24]byte
	 */
	balanceOfPrefix = "0x70a08231000000000000000000000000"
)

type ethclient struct {
	valid  bool
	url    string
	client *geth.EthereumClient
}

type EthCluster struct {
	Clients []*ethclient
}

func NewClientCluster(urls []string) *EthCluster {
	var cluster = &EthCluster{Clients: make([]*ethclient, len(urls))}

	for i, url := range urls {
		if client, err := geth.NewEthereumClient(url); err != nil {
			cluster.Clients[i] = &ethclient{client: client, valid: false, url: url}
		} else {
			cluster.Clients[i] = &ethclient{client: client, valid: true, url: url}
		}
	}
	return cluster
}

func (c *EthCluster) Add(url string) error {
	for _, client := range c.Clients {
		if client.url == url {
			return errors.New("client already exist")
		}
	}
	if client, err := geth.NewEthereumClient(url); err != nil {
		nc := &ethclient{client: client, valid: false, url: url}
		c.Clients = append(c.Clients, nc)
	} else {
		nc := &ethclient{client: client, valid: true, url: url}
		c.Clients = append(c.Clients, nc)
	}
	return nil
}

func (c *EthCluster) GetBalance(contract string, address string, number int64) (*big.Int, error) {
	msg, err := buildBalanceOfMsg(contract, address)
	if err != nil {
		return big.NewInt(0), err
	}

	for _, client := range c.Clients {
		if client.valid {
			ctx := geth.NewContext()
			if balance, err := client.client.CallContract(ctx, msg, number); err != nil {
				return big.NewInt(0), err
			} else {
				return big.NewInt(0).SetBytes(balance), nil
			}
		}
	}
	return big.NewInt(0), nil
}

func (c *EthCluster) SendTransaction(contract string, from string, to string, value *big.Int) ([]byte, error) {
	return nil, nil
}

func buildBalanceOfMsg(contract string, holder string) (*geth.CallMsg, error) {
	if len(holder) != 40 {
		return nil, errors.New(fmt.Sprintf("invalid holder address:%s", holder))
	}
	data := balanceOfPrefix + holder

	ethContractAddr, err := geth.NewAddressFromHex(contract)
	if err != nil {
		return nil, err
	}
	msg := geth.NewCallMsg()
	msg.SetTo(ethContractAddr)
	msg.SetData([]byte(data))
	return msg, nil
}
