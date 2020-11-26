#!/bin/bash 
source ./env.sh
rest_laddr="http://127.0.0.1:1317"

test_oin_erc20_addr="a4d595F42f3b9CF98d1afe2EFa027c06280662c3"
test_cross_chain_key="t2"

#query_exist=$($CLI keys show ${test_cross_chain_key} 2>/dev/null | grep "name")
#new_test_cross_chain_key=0

#if [ "${query_exist}" == "" ];
#then
#    echo "add key $test_cross_chain_key"
#    new_test_cross_chain_key=1
#    $CLI keys add ${test_cross_chain_key} <<EOM
#$PASSWD
#EOM
#fi

test_cross_chain_address=`$CLI keys show ${test_cross_chain_key} -a`

# send feecoin to new account.
#$CLI tx send $KEY $test_cross_chain_address 10000000000feecoin 
#if [ "$new_test_cross_chain_key" == "1" ];
#then
#    echo "wait tx send processed"
#    sleep 10
#fi

from=$test_cross_chain_address
unsignedTx="unsigned_map_verify.json"
signedTx="signed_map_verify.json"


curl -s  -X POST --data-binary "{\"base_req\":{\"from\":\"$from\",\"chain_id\":\"$CHAINID\",\"fees\":[{\"denom\":\"feecoin\",\"amount\":\"100\"}],\"generate_only\":true},\"erc_addr\":\"$test_oin_erc20_addr\",\"cc_addr\":\"${test_cross_chain_address}\"}" ${rest_laddr}/mapping/verify/ > $unsignedTx

accountInfo=`$CLI query account $($CLI keys show $test_cross_chain_key -a)`
account_number=`echo $accountInfo | jq -r .value.account_number`
sequence=`echo $accountInfo | jq -r .value.sequence`
$CLI tx sign $unsignedTx --chain-id=$CHAINID --from=$from --account-number=${account_number} --sequence=${sequence} --offline > $signedTx 
$CLI tx broadcast $signedTx
echo "finished"
