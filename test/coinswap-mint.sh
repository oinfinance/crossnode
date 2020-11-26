#!/bin/bash 
source ./env.sh
rest_laddr="http://127.0.0.1:1317"

test_txhash="532D3B5B3681F660AF6B80F31A6DAA2DC5DD05102F955AC7E8585346816CE332"
test_fromchain="ethereum"
test_fromaddr="a4d595F42f3b9CF98d1afe2EFa027c06280662c3"
test_token="oin"
test_value=55555
test_tochain="near"
test_toaddr="a4d595F42f3b9CF98d1afe2EFa027c06280662c3"

test_borker_key="t2"

query_exist=$($CLI keys show ${test_borker_key} 2>/dev/null | grep "name")
new_test_borker_key=0

if [ "${query_exist}" == "" ];
then
    echo "add key $test_borker_key"
    new_test_borker_key=1
    $CLI keys add ${test_borker_key} <<EOM
$PASSWD
EOM
fi

test_borker_address=`$CLI keys show ${test_borker_key} -a`

# send feecoin to new account.
$CLI tx send $KEY $test_borker_address 10000000000feecoin 
if [ "$new_test_borker_key" == "1" ];
then
    echo "wait tx send processed"
    sleep 10
fi

from=$test_borker_address
unsignedTx="unsigned_coinswap_mint.json"
signedTx="signed_coinswap_mint.json"


curl -s  -X POST --data-binary "{\"base_req\":{\"from\":\"$from\",\"chain_id\":\"$CHAINID\",\"fees\":[{\"denom\":\"feecoin\",\"amount\":\"100\"}],\"generate_only\":true},\"txHash\":\"$test_txhash\",\"fromChain\":\"$test_fromchain\", \"fromAddr\":\"$test_fromaddr\",\"token\":\"$test_token\",\"value\":\"$test_value\",\"toAddr\":\"$test_toaddr\",\"toChain\":\"$test_tochain\"}" ${rest_laddr}/coinswap/mint/ > $unsignedTx
accountInfo=`$CLI query account $($CLI keys show $test_borker_key -a)`
account_number=`echo $accountInfo | jq -r .value.account_number`
sequence=`echo $accountInfo | jq -r .value.sequence`
$CLI tx sign $unsignedTx --chain-id=$CHAINID --from=$from --account-number=${account_number} --sequence=${sequence} --offline > $signedTx 
$CLI tx broadcast $signedTx
echo "finished"
