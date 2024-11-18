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

#eth.getPrivateTransactionReceipt("0x4a56e16174d64aa2df8841cce3e9790f3eb986d39ac91148397c8e1229c5b676")
#eth.getTransactionReceipt("0x63d7967c0c243f09f7d6861806cdf1de9b008f63bd4b6182d9239377f56236d4")

# private from 2 to 1
# curl -X POST http://localhost:20000 --data '{ "jsonrpc":"2.0", "id":2, "method":"eth_getPrivateTransactionByHash", "params": ["0x1c472f8dfd6d0367f531e306f69960173100ef60b61bc6a845f393790e34bd29"]}' --header "Content-Type: application/json"
# curl -X POST http://localhost:20002 --data '{ "jsonrpc":"2.0", "id":2, "method":"eth_getPrivateTransactionByHash", "params": ["0x1c472f8dfd6d0367f531e306f69960173100ef60b61bc6a845f393790e34bd29"]}' --header "Content-Type: application/json"
# curl -X POST http://localhost:20004 --data '{ "jsonrpc":"2.0", "id":2, "method":"eth_getPrivateTransactionByHash", "params": ["0x1c472f8dfd6d0367f531e306f69960173100ef60b61bc6a845f393790e34bd29"]}' --header "Content-Type: application/json"

# curl -X POST http://localhost:20000 --data '{ "jsonrpc":"2.0", "id":2, "method":"eth_getPrivateTransactionReceipt", "params": ["0x1c472f8dfd6d0367f531e306f69960173100ef60b61bc6a845f393790e34bd29"]}' --header "Content-Type: application/json"
# curl -X POST http://localhost:20002 --data '{ "jsonrpc":"2.0", "id":2, "method":"eth_getPrivateTransactionReceipt", "params": ["0x1c472f8dfd6d0367f531e306f69960173100ef60b61bc6a845f393790e34bd29"]}' --header "Content-Type: application/json"
# curl -X POST http://localhost:20004 --data '{ "jsonrpc":"2.0", "id":2, "method":"eth_getPrivateTransactionReceipt", "params": ["0x1c472f8dfd6d0367f531e306f69960173100ef60b61bc6a845f393790e34bd29"]}' --header "Content-Type: application/json"

# curl -X POST http://localhost:8545 --data '{ "jsonrpc":"2.0", "id":2, "method":"eth_getPrivateTransactionReceipt", "params": ["0x1c472f8dfd6d0367f531e306f69960173100ef60b61bc6a845f393790e34bd29"]}' --header "Content-Type: application/json"