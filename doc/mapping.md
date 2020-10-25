# mapping 模块
mapping 模块用于处理用户在 ETH 链上的oin资产向 cross-chain 链上的映射。
映射过来的资产表示用户持有的资产，将在轮转挖矿时获得不同比例的出块概率。

# 映射步骤
例如，用户在以太坊上的账户为 A(0x..85), 启动了一个crossnode后创建了新的地址 C(oa....99).

step1:用户首先在 crossnode上用账户 C发送一笔mapping交易，附带 A账户信息，C 与 A 之间处于待验证关系.

step2:用户再使用账户A从以太坊上向合约中发送以太验证注册的交易，附带上参数C账户信息，borker节点监听到此交易之后
转换为verify交易转发到cross-node中，cross-node 通过互相比较之后，C 和 A 的绑定完成，

step3:cross链每隔150个区块会从以太坊上查询A账户的余额然后同步到账户C上。

