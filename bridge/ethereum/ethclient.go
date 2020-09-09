package ethereum

import (
	"encoding/hex"
	"errors"
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

func (c *EthCluster) GetBalance(contract string, address []byte) (*big.Int, error) {
	msg,err := BuildBalanceOfMsg(contract, address)
	if err != nil {
		return big.NewInt(0), err
	}

	for _, client := range c.Clients {
		if client.valid {
			ctx := geth.NewContext()
			if balance, err := client.client.CallContract(ctx, msg, -1); err != nil {
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

func BuildBalanceOfMsg(contract string, holder []byte) (*geth.CallMsg, error) {
	strholder := hex.EncodeToString(holder)
	data := balanceOfPrefix + strholder

	ethContractAddr,err := geth.NewAddressFromHex(contract)
	if err != nil {
		return nil, err
	}
	msg := geth.NewCallMsg()
	msg.SetTo(ethContractAddr)
	msg.SetData([]byte(data))
	return msg, nil
}