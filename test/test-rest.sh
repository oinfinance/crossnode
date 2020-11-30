#!/bin/bash 
source ./env.sh
rest_laddr="http://127.0.0.1:1317"

test_account=t1
$CLI keys add ${test_account} <<EOM
$PASSWD
EOM

from=`$CLI keys show $KEY -a `
to=`$CLI keys show ${test_account} -a`

unsignedTx="unsignedTx.json"
signedTx="signedTx.json"

# 生成未签名的交易数据
curl -s -k -H "Content-Type: application/json" -d "{\"base_req\":{\"from\":\"$from\",\"chain_id\":\"$CHAINID\",\"fees\":[{\"denom\":\"coin1\",\"amount\":\"100\"}],\"generate_only\":true},\"amount\":[{\"denom\":\"coin0\",\"amount\":\"5000\"}]}" ${rest_laddr}/bank/accounts/{$to}/transfers > unsignedTx.json

accountInfo=`$CLI query account $($CLI keys show $from -a)`
account_number=`echo $accountInfo | jq -r .value.account_number`
sequence=`echo $accountInfo | jq -r .value.sequence`
# 签名交易
$CLI tx sign $unsignedTx --chain-id=$CHAINID --from=$from --account-number=${account_number} --sequence=${sequence} --offline > $signedTx
# 广播交易
$CLI tx broadcast $signedTx
