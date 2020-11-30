## 文件列表
```
env.sh  配置节点启动时的部分自定义参数
init.sh 在本地重新创建一个节点实例并运行，日志输出到 d.log
cstart.sh 启动rest-server服务, 日志输出到 c.log
test-rest.sh 展示了如何使用rest-server和oincli进行转账交易
map-register.sh 展示了如何实现oin持币用户向链上注册，实现代币映射
map-verify.sh   展示了用户如何在register之后，执行双向验证，验证成功后会实现代币映射
coinswap-mint.sh 展示了在监听到用户抵押了代币后，如何为用户创建铸币交易
coinswap-burn.sh 展示了监听到用户销毁代币时，如何为其兑换原始代币
```
## 概念
为了简化文档描述，下面定义几个项目中需要用到的几个概念.

##### 母链
母链是指 OIN ERC20 代币发行的链，以太坊主网或者测试网。

母链上包含了 OIN ERC20 合约，用户注册合约.

##### 原链
原链是指跨链资产的原始链，负责原始资产的抵押锁定和解锁释放.

原链上包含了锁仓合约, 用于控制资产的锁定和释放，合约可以更换管理者，拥有释放资产的权利.

##### 目标链
目标链是指跨链资产的目标链，负责铸造和销毁对应的映射资产。

目标链上包含映射的资产合约，支持铸币和销毁，可以更换管理者，拥有铸币的权利.

## 代币
Cross-Node 中有两种代币，一为 feecoin，创世时产生，用于用户在链上发送交易，普通用户可以申请获取feecoin；
二为 oincoin，通过从ethereum的ERC20 OIN 映射产生， 以及收益(挖矿、staking)产生，用户可以提取收益到ERC20账户.

## 节点的创建和启动
#### 创建节点
脚本默认使用存储位置为 ~/.oind 和 ~/.oincli 目录，执行后会删除所有旧的数据，并重新创建新的单节点链.
```
# ./init.sh 
```

#### 启动rest-server服务
脚本在 1317 端口开启`rest-server`服务，端口号可在 env.sh 中修改 `REST_LADDR`
```
# ./cstart.sh
```

## 用户注册和oin代币映射
用户注册指的是持有 OIN ERC20 代币的用户与链上的账户进行绑定，绑定后链上会获取用户的oin持有数量映射到链上。
用户注册需要几个步骤：
#### 1. 在链上创建新账户
使用命令 `oincli keys add [keyname]` 新建账户，keyname为自定义的名称
```
$ oincli keys add t1
Enter a passphrase to encrypt your key to disk:
Repeat the passphrase:
{
  "name": "t1",
  "type": "local",
  "address": "oaa1u9ynpdxtjuhmgh7lth6c93s66ml6hfx6dna9tj",
  "pubkey": "oap1addwnpepqtvgmtp97q5n3gejnnejgvtsf57enfth0kutpmq3m6md0ff7mfdg5zag7kz",
  "mnemonic": "another swear ritual reward age resource wreck force make panel federal cross engine solve wing awful submit real sketch squeeze bachelor rabbit worry rural"
}
$ 
```
#### 2. 账户激活
上一步骤中新建的账户是离线账户，并且没有任何资产，需要链上其他账户为其发送一部分feecoin用于以后操作.

示例使用创世账户`mykey`进行转账，默认密码为`12345678`, 是在env.sh中配置的.
```
$ oincli tx send mykey oaa1u9ynpdxtjuhmgh7lth6c93s66ml6hfx6dna9tj 10000feecoin
```
#### 3. 用户通过以太坊合约进行注册
此步骤属于合约设计的范畴，必要的参数为上一步骤中新建的账户地址`oaa1u9ynpdxtjuhmgh7lth6c93s66ml6hfx6dna9tj`

#### 4. borker 监控触发
borker 监控注册合约，监听到有新用户注册时，组装 map-register 交易，示例见 [map-register.sh](https://github.com/oinfinance/crossnode/blob/master/test/map-register.sh)

#### 5. 链上用户双向验证
borker 发送了`map-register` 交易之后，需要用户使用上面新建的账户 `oaa1u9ynpdxtjuhmgh7lth6c93s66ml6hfx6dna9tj`在链上发送`map-verify`交易

命令行方式：`oincli tx mapping verify [crossnode_address] [erc_address]`

```
oincli tx mapping verify oaa1u9ynpdxtjuhmgh7lth6c93s66ml6hfx6dna9tj a4d595F42f3b9CF98d1afe2EFa027c06280662c3
```

rest-server 方式：示例见 [map-verify.sh](https://github.com/oinfinance/crossnode/blob/master/test/map-verify.sh)

注: 用户在 map-register 交易发送之后，要在 150块内执行 map-verify, 否则将过期失效.
#### 6. 查询余额
上述步骤完成后即完成了账户映射，链上每隔 150块进行一次刷新，用户使用下面的命令进行查询.
```
$ oincli query account oaa1u9ynpdxtjuhmgh7lth6c93s66ml6hfx6dna9tj
{
  "type": "cosmos-sdk/Account",
  "value": {
    "address": "oaa1u9ynpdxtjuhmgh7lth6c93s66ml6hfx6dna9tj",
    "coins": [
      {
        "denom": "feecoin",
        "amount": "10000"
      }
    ],
    "public_key": {
      "type": "tendermint/PubKeySecp256k1",
      "value": "AtiNrCXwKTijMpzzJDFwTT2ZpXd9uLDsEd6216U+2lqK"
    },
    "account_number": "6",
    "sequence": "1"
  }
}
```

## 铸币
铸币的行为发生在用户在原链上发生了锁仓之后，具体有一下几个步骤.
#### 1.用户锁定资产
本步是用户在原链上完成，将原始资产锁定到合约中

#### 2. borker监控
borker程序监听到用户锁仓的操作后，组装`coinswap-mint`交易，并签名发送到 cross chain上,得到交易哈希`mint_txhash`.

示例见 [coinswap-mint.sh](https://github.com/oinfinance/crossnode/blob/master/test/coinswap-mint.sh)

等待交易被打包，然后查询交易执行结果, 保存其中的 data 字段的值，记做 `mint_data`
```
$ oincli query tx A70615C21F9A35FAE6B0225A964B80D103E0A130102B2047D02F6E1037655EB1
  {
    "height": "1713",
    "txhash": "A70615C21F9A35FAE6B0225A964B80D103E0A130102B2047D02F6E1037655EB1",
    "data": "9F315AECDAB8B65D2B2BA5EA74979B20DABB3110A1881DA1CF56E56A5DBCF5E6",
    "raw_log": "[{\"msg_index\":0,\"success\":true,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"coinswap\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "success": true,
        "log": "",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "coinswap"
              }
            ]
          }
        ]
      }
    ],
    "gas_wanted": "200000",
    "gas_used": "34029",
    "tx": {
      "type": "cosmos-sdk/StdTx",
      "value": {
        "msg": [
          {
            "type": "coinswap/MsgCoinSwap",
            "value": {
              "sender": "oaa1mhgegl29q5e0yutzflsz62yzn6y6qakr55xgfe",
              "txHash": "532D3B5B3681F660AF6B80F31A6DAA2DC5DD05102F955AC7E8585346816CE332",
              "fromChain": "1",
              "fromAddr": "a4d595F42f3b9CF98d1afe2EFa027c06280662c3",
              "token": "65537",
              "value": "55555",
              "toAddr": "a4d595F42f3b9CF98d1afe2EFa027c06280662c3",
              "toChain": "3",
              "eventType": "1"
            }
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "feecoin",
              "amount": "100"
            }
          ],
          "gas": "200000"
        },
        "signatures": [
          {
            "pub_key": {
              "type": "tendermint/PubKeySecp256k1",
              "value": "AwQSWQJ2k8IZPpn85sOt28qmCW9GFVFIrGrBPwOxULFl"
            },
            "signature": "PE0SrOFZTgbdoq4ROp48s18VOCsBmdaQn/FwI7tmtS8AcT6LkE+DwaHku4eUh16GAT36+vDKXASBD94Vcd0huw=="
          }
        ],
        "memo": ""
      }
    },
    "timestamp": "2020-11-30T05:42:24Z",
    "events": [
      {
        "type": "message",
        "attributes": [
          {
            "key": "action",
            "value": "coinswap"
          }
        ]
      }
    ]
  }

```

#### 3. 跨链执行 mint 交易
cross-chain 执行 mint 交易时，仅仅是保存了该条记录，并把记录的查询索引写到了data中，即上一步查询交易获取的 `mint_data`.

等到过了 150 个区块以后，cross-chain 会对该条mint签发一个铸币证明，用户可以通过下面的命令查询.
```
# 等待签发状态的receipt
$ oincli query coinswap receipt 9F315AECDAB8B65D2B2BA5EA74979B20DABB3110A1881DA1CF56E56A5DBCF5E6
{
  "type": "coinswap/CoinSwapReceipt",
  "value": {
    "status": "1",
    "receipt": ""
  }
}

# 签发完成状态的receipt
$ oincli query coinswap receipt 9F315AECDAB8B65D2B2BA5EA74979B20DABB3110A1881DA1CF56E56A5DBCF5E6
{
  "type": "coinswap/CoinSwapReceipt",
  "value": {
    "status": "2",
    "receipt": "73daa5284916bbfd5b5b71842afdb93da96865791d74af034c8da9d850dacdeb"
  }
}
```
`status` 为 1 表示mint还未签发； 为 2 表示签发成功，receipt 为签发的证明.

`receipt` 为用户提取铸币的证明.

#### 4. 用户提取铸币
用户使用第三步中的 `receipt` 作为参数，调用目标链的资产合约，即可获得相应数量的映射资产.

## 销毁
销毁的行为发生在用户在目标链上发生了赎回动作，具体有以下几个步骤.
#### 1.用户赎回/销毁资产
本步是用户在目标链上完成

#### 2. borker监控
borker程序监听到用户销毁的操作后，组装`coinswap-burn`交易，并签名发送到 cross chain上,得到交易哈希`burn_txhash`.

示例见 [coinswap-burn.sh](https://github.com/oinfinance/crossnode/blob/master/test/coinswap-burn.sh)

等待交易被打包，然后查询交易执行结果, 保存其中的 data 字段的值，记做 `burn_data`
```
$ oincli query tx 1830A1256F532955D1849B6709B865840F29A898D17F732D1C3176A2D86EAB26
  {
    "height": "1868",
    "txhash": "1830A1256F532955D1849B6709B865840F29A898D17F732D1C3176A2D86EAB26",
    "data": "5A940F2F5CF1F6A71E0EAC50CC024A601DEC82B7AEA9A36FB90D32C5B11D36B5",
    "raw_log": "[{\"msg_index\":0,\"success\":true,\"log\":\"\",\"events\":[{\"type\":\"message\",\"attributes\":[{\"key\":\"action\",\"value\":\"coinswap\"}]}]}]",
    "logs": [
      {
        "msg_index": 0,
        "success": true,
        "log": "",
        "events": [
          {
            "type": "message",
            "attributes": [
              {
                "key": "action",
                "value": "coinswap"
              }
            ]
          }
        ]
      }
    ],
    "gas_wanted": "200000",
    "gas_used": "35773",
    "tx": {
      "type": "cosmos-sdk/StdTx",
      "value": {
        "msg": [
          {
            "type": "coinswap/MsgCoinSwap",
            "value": {
              "sender": "oaa1mhgegl29q5e0yutzflsz62yzn6y6qakr55xgfe",
              "txHash": "532D3B5B3681F660AF6B80F31A6DAA2DC5DD05102F955AC7E8585346816CE332",
              "fromChain": "3",
              "fromAddr": "a4d595F42f3b9CF98d1afe2EFa027c06280662c3",
              "token": "65537",
              "value": "55555",
              "toAddr": "a4d595F42f3b9CF98d1afe2EFa027c06280662c3",
              "toChain": "1",
              "eventType": "0"
            }
          }
        ],
        "fee": {
          "amount": [
            {
              "denom": "feecoin",
              "amount": "100"
            }
          ],
          "gas": "200000"
        },
        "signatures": [
          {
            "pub_key": {
              "type": "tendermint/PubKeySecp256k1",
              "value": "AwQSWQJ2k8IZPpn85sOt28qmCW9GFVFIrGrBPwOxULFl"
            },
            "signature": "8C7YVGU4TdCaispc4c41zxCtoof43OAZuohmv6NkHuVaAck+prZ+HezJgKlCiMfWb1LFYuIIDBj7Wwy/wROf3A=="
          }
        ],
        "memo": ""
      }
    },
    "timestamp": "2020-11-30T05:55:26Z",
    "events": [
      {
        "type": "message",
        "attributes": [
          {
            "key": "action",
            "value": "coinswap"
          }
        ]
      }
    ]
  }
```

#### 3. 跨链执行 burn 交易
cross-chain 执行 burn 交易时，仅仅是保存了该条记录，并把记录的查询索引写到了data中，即上一步查询交易获取的 `burn_data`.

等到过了 150 个区块以后，cross-chain 会对该条burn签发一个释放资产证明，用户可以通过下面的命令查询.
```
# 等待签发状态的receipt
$ oincli query coinswap receipt 9F315AECDAB8B65D2B2BA5EA74979B20DABB3110A1881DA1CF56E56A5DBCF5E6
{
  "type": "coinswap/CoinSwapReceipt",
  "value": {
    "status": "1",
    "receipt": ""
  }
}

# 签发完成状态的receipt
$ oincli query coinswap receipt 9F315AECDAB8B65D2B2BA5EA74979B20DABB3110A1881DA1CF56E56A5DBCF5E6
{
  "type": "coinswap/CoinSwapReceipt",
  "value": {
    "status": "2",
    "receipt": "d5b5b7bbf52873491612afdb93da9688654a31ddaa47949d78acdef0bc8d50da"
  }
}
```
`status` 为 1 表示burn还未签发； 为 2 表示签发成功，receipt 为签发的证明.

`receipt` 为用户赎回资产的证明.

#### 4. 用户赎回资产
用户使用第三步中的 `receipt` 作为参数，调用原始链的锁仓/释放合约，即可赎回相应数量的资产.

