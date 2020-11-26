#!/bin/bash 
source ./env.sh
oind_node="tcp://localhost:26657"

$CLI rest-server --node=${oind_node} --chain-id=${CHAINID} --laddr=${REST_LADDR}  > c.log 2>&1 &

