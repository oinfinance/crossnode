## coinswap
假设用户需要将 A 映射到链 B 上，将在链B上创建相应的代币 oA，此过程叫Mint，
oA 可以在链B上参与各种DAPP生态应用，
当用户想要兑换回A时，向oA的合约地址（或者CrossNode链）发送销毁的消息，
CrossNode将收回oA，并在A上释放相应数量的代币到所提供的账户上，此过程叫Burn.

#### Mint 铸币
A-Account   : Borker节点监控的A上的地址，监听用户锁仓信息；
UA-Account  : 用户在A上的地址，用户通过此账户向 A-Account 中锁仓；
TX-A        : 用户在A上锁仓时的交易信息；
Borker监控到之后，提供 Chain_A, Chain_B, TX-A，  
B-Account   : 
`CrossNode`主要完成的功能是由*Borker节点*监听`链A`上的*监管账户*，
当监控到有用户在`链A`上向*监管账户*中锁定了一定数量的资产时，
*Borker节点*向`CrossNode`发送`MsgMint`交易，`CrossNode`根据`MsgMint`的内容，
签署一个交易签名，用于在`链B`上铸币，然后用户使用该签名在链B上领取铸币。

#### 销毁
当用户想要将``

coinswap 主要实现由crossnode管控账户的资产铸造，销毁，以及释放等操作。

包含的消息类型包括：

- Mint: 铸造代币
例如：用户在链A上锁定了一定量代币，想要转到链B上，那么在链B上需要按照 1:1 比例相应数量的代币 Cro_A, 
Cro_A 代币的管控由 CrossNode 管理。

消息内容：
```
type msgMint struct {
    ChainA  int `json:"chain_A"`
    ChainB  int `json:"chain_B"`
    Amount  int `json:"amount"`
    Reciept []byte  `json:"reciept"`    // 锁仓凭证
    Account []byte  `json:"account"`    // 铸币账户 
}
```

当发生了Mint交易时，需要CrossNode对Mint的合法性进行验证和签名，当签名通过后，将签署一个chain_B上的交易，用于在B上获取铸币。



- Burn：销毁代币
例如：用户在链B上，想要将Cro_A 兑换回链A上的代币，那么需要销毁相应数量的 Cro_A，此权限由CrossNode管理。

