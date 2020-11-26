#!/bin/bash 
source ./env.sh
unsignedTx="unsignedTx.json"
signedTx="signedTx.json"

from=$KEY
accountInfo=`$CLI query account $($CLI keys show $from -a)`
account_number=`echo $accountInfo | jq -r .value.account_number`
sequence=`echo $accountInfo | jq -r .value.sequence`
$CLI tx sign $unsignedTx --chain-id=$CHAINID --from=$from --account-number=${account_number} --sequence=${sequence} --offline > $signedTx 
$CLI tx broadcast $signedTx
