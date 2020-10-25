## mining

mining 模块用来记录和分配挖矿收益,
在模块内记录各个节点的未提现收益数量，用户通过发送交易进行提现。

tx: MsgWithDraw 提现

query: QueryRewards 查询可提现余额数量

endBlocker: 为出块者增加挖矿收益，增加Rewards 余额数量。