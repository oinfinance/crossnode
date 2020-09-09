package ethereum

import (
	"errors"
	"github.com/ethereum/go-ethereum/mobile"
	"math/big"
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

func (c *EthCluster) GetBalance(contract string, address string) (*big.Int, error) {
	for _, client := range c.Clients {
		if client.valid {
			//client.client.SendTransaction()
		}
	}
	return big.NewInt(0), nil
}

func (c *EthCluster) SendTransaction(contract string, from string, to string, value *big.Int) ([]byte, error) {
	return nil, nil
}
