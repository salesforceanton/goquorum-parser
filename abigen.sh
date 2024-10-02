#!/bin/bash

ABI_DIRS=(
services/backend/cryptogate/account_manager_abi 
services/backend/cryptogate/node_manager_abi 
services/backend/cryptogate/org_manager_abi 
services/backend/cryptogate/permissions_impl_abi 
services/backend/cryptogate/role_manager_abi 
services/backend/cryptogate/voter_manager_abi 
)
PACKAGES=(account_manager_abi node_manager_abi org_manager_abi permissions_impl_abi role_manager_abi voter_manager_abi)
ABI_FILES=(account_manager.abi node_manager.abi org_manager.abi permissions_impl.abi role_manager.abi voter_manager.abi)

for ((i=0; i<${#ABI_DIRS[@]}; i++)); do
    ABI_FILE="${ABI_DIRS[$i]}/${ABI_FILES[$i]}"
    PACKAGES_FILE="${ABI_DIRS[$i]}/${PACKAGES[$i]}.go"

    if [ ! -f "$PACKAGES_FILE" ] || [ "$(md5sum $ABI_FILE | awk '{print $1}')" != "$(md5sum $PACKAGES_FILE | awk '{print $1}')" ]; then
        abigen --abi=$ABI_FILE --pkg=${PACKAGES[$i]} --out=$PACKAGES_FILE --alias=_totalSupply=totalSupplyPrivate
    fi
done