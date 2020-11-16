package near

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

type nearclient struct {
	valid  bool
	url    string
	client *geth.EthereumClient
}

type NearCluster struct {
	Clients []*nearclient
}

func NewClientCluster(urls []string) *NearCluster {
	var cluster = &NearCluster{Clients: make([]*nearclient, len(urls))}

	for i, url := range urls {
		if client, err := geth.NewEthereumClient(url); err != nil {
			cluster.Clients[i] = &nearclient{client: client, valid: false, url: url}
		} else {
			cluster.Clients[i] = &nearclient{client: client, valid: true, url: url}
		}
	}
	return cluster
}

func (c *NearCluster) Add(url string) error {
	for _, client := range c.Clients {
		if client.url == url {
			return errors.New("client already exist")
		}
	}
	if client, err := geth.NewEthereumClient(url); err != nil {
		nc := &nearclient{client: client, valid: false, url: url}
		c.Clients = append(c.Clients, nc)
	} else {
		nc := &nearclient{client: client, valid: true, url: url}
		c.Clients = append(c.Clients, nc)
	}
	return nil
}

func (c *NearCluster) GetBalance(contract string, address []byte) (*big.Int, error) {
	msg, err := BuildBalanceOfMsg(contract, address)
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

func (c *NearCluster) SendTransaction(contract string, from string, to string, value *big.Int) ([]byte, error) {
	return nil, nil
}

func BuildBalanceOfMsg(contract string, holder []byte) (*geth.CallMsg, error) {
	strholder := hex.EncodeToString(holder)
	data := balanceOfPrefix + strholder

	ethContractAddr, err := geth.NewAddressFromHex(contract)
	if err != nil {
		return nil, err
	}
	msg := geth.NewCallMsg()
	msg.SetTo(ethContractAddr)
	msg.SetData([]byte(data))
	return msg, nil
}
