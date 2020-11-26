#!/bin/bash 
source ./env.sh
rest_laddr="http://127.0.0.1:1317"

test_account=t1
$CLI keys add ${test_account} <<EOM
$PASSWD
EOM

from=`$CLI keys show $KEY -a `
to=`$CLI keys show ${test_account} -a`

#curl -s -k -H "Content-Type: application/json" -d "{\"base_req\":{\"from\":\"$from\",\"memo\":\"\",\"chain_id\":\"$CHAINID\",\"account_number\":\"0\",\"sequence\":\"1\",\"gas\":\"200000\",\"gas_adjustment\":\"1.2\",\"fees\":[{\"denom\":\"coin0\",\"amount\":\"1\"}],\"simulate\":false, \"generate_only\":true},\"amount\":[{\"denom\":\"coin0\",\"amount\":\"1\"}]}" ${rest_laddr}/bank/accounts/{$to}/transfers > unsignedTx.json
#curl -s -k -H "Content-Type: application/json" -d "{\"base_req\":{\"from\":\"$from\",\"memo\":\"\",\"chain_id\":\"$CHAINID\",\"gas\":\"200000\",\"gas_adjustment\":\"1.2\",\"fees\":[{\"denom\":\"coin0\",\"amount\":\"1\"}],\"simulate\":false, \"generate_only\":true},\"amount\":[{\"denom\":\"coin0\",\"amount\":\"1\"}]}" ${rest_laddr}/bank/accounts/{$to}/transfers > unsignedTx.json
curl -s -k -H "Content-Type: application/json" -d "{\"base_req\":{\"from\":\"$from\",\"chain_id\":\"$CHAINID\",\"fees\":[{\"denom\":\"coin1\",\"amount\":\"100\"}],\"generate_only\":true},\"amount\":[{\"denom\":\"coin0\",\"amount\":\"5000\"}]}" ${rest_laddr}/bank/accounts/{$to}/transfers > unsignedTx.json

