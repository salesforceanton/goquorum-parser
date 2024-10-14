#!/bin/bash

geth attach http://localhost:8545

# MEMBER_NUMBER=$1
# MEMBER_NODE="member${MEMBER_NUMBER}quorum"
# echo "Attempting to connect to $MEMBER_NODE"
# if [ -z `docker compose -f docker-compose.yml ps -q ${MEMBER_NODE} 2>/dev/null` ] ; then
#   echo "$MEMBER_NUMBER is not a valid member node. Must be between 1 and 3" >&2
#   exit 1
# fi

# docker compose exec ${MEMBER_NODE} /bin/sh -c "geth attach /data/geth.ipc"

#eth.getPrivateTransactionReceipt("0x802b9d78080bf7045b99b6878068b34134de8d6a16a7868b8508988b6a29625e")
#eth.getTransactionReceipt("0x802b9d78080bf7045b99b6878068b34134de8d6a16a7868b8508988b6a29625e")

# curl -X POST http://localhost:20000 --data '{ "jsonrpc":"2.0", "id":2, "method":"eth_getPrivateTransactionByHash", "params": ["0x802b9d78080bf7045b99b6878068b34134de8d6a16a7868b8508988b6a29625e"]}' --header "Content-Type: application/json"
# curl -X POST http://localhost:20002 --data '{ "jsonrpc":"2.0", "id":2, "method":"eth_getPrivateTransactionByHash", "params": ["0x802b9d78080bf7045b99b6878068b34134de8d6a16a7868b8508988b6a29625e"]}' --header "Content-Type: application/json"