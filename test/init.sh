#!/bin/bash

source ./env.sh

# remove existing daemon and client
rm -rf ~/.oin*

#make install

#$CLI config keyring-backend test

# Set up config for CLI
$CLI config chain-id $CHAINID
$CLI config output json
$CLI config indent true
$CLI config trust-node true

# if $KEY exists it should be deleted
$CLI keys add $KEY <<EOM
$PASSWD
EOM

# Set moniker and chain-id for oind (Moniker can be anything, chain-id must be an integer)
$OIND init $MONIKER --chain-id $CHAINID

# Allocate genesis accounts (cosmos formatted addresses)
$OIND add-genesis-account $($CLI keys show $KEY -a) 10000000000000000000000feecoin,10000000000000000oincoin

# Sign genesis transaction
$OIND gentx --name $KEY --amount 1000000oincoin<<EOM
$PASSWD
EOM

# Collect genesis tx
$OIND collect-gentxs

# Run this to ensure everything worked and that the genesis file is setup correctly
$OIND validate-genesis
sed  's/stake/oincoin/g' -i ~/.oind/config/genesis.json

# Command to run the rest server in a different terminal/window
echo -e '\nrun the following command in a different terminal/window to run the REST server :'
echo -e "$CLI rest-server --laddr \"tcp://localhost:8545\" --unlock-key $KEY --chain-id $CHAINID --trace\n"

# Start the node (remove the --pruning=nothing flag if historical queries are not needed)
$OIND start --pruning=nothing --rpc.unsafe --log_level "main:info,state:info,mempool:info" --trace > d.log 2>&1 &
